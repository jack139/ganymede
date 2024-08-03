## 模块数据结构设计



### 1. Ganymede

平台数据。list users存放从API注册的用户信息。用户类型：USR 普通用户、ORG 医疗机构、GOV 政府机构、3PT 其他机构



#### 用户信息 list users

```golang
chainAddr      // 区块链用户地址(AccAddress)  --index
keyName        // 用户key名称，限英文
userType       // 用户类型: USR, 3PT, ORG, GOV
name           // 账户名称
address        // 地址
phone          // 电话
accountNo      // 账户号码
ref            // 其他信息
regDate        // 注册日期
status         // 用户状态： ACTIVE, CLOSE, WAIT
lastDate       // 最后修改日期
link_status    // link list
```



### 2. Zoo

应用业务KV数据。数据内容均私钥加密。



#### 数据信息 map kvzoo

```golang
owner          // 所有者chainAddr   --index
zooKey         // 键               --index
zooValue       // 值，在api端私钥加密
lastDate       // 最后修改日期
link_owner     // link list
```



### 3. Exchange

用户授权的数据交换。同一链上的用户之间；不同链上的用户之间。



#### 请求记录 list ask

```golang
id            // id(自动生成)  --index
sender        // 请求者chainAddr
replier       // 响应者chainAddr
payload       // 请求的内容
sentDate      // 发送时间
creator       // 创建者（api调用者）
link_sender   // link list
link_replier  // link list
```



#### 响应记录 list reply

```golang
id            // id(自动生成)  --index
askId         // 请求id
sender        // 请求者chainAddr
replier       // 响应者chainAddr
payload       // 响应的内容
sentDate      // 发送时间
creator       // 创建者（api调用者）
link_sender   // link list
link_replier  // link list
```



### 4. PostOffice (IBC)

跨链传输的数据记录



#### 收到的信息 list post

```golang
id            // id(自动生成)  --index
title         // 标题（用于存uuid）
payload       // 内容
fromChain     // IBC channel
sender        // 发送者chainAddr
receiver      // 接收者chainAddr
senderInfo    // 发送者的描述信息 (users信息摘要)
sentDate      // 发送时间
link_sender     // link list
link_receiver   // link list
```



#### 发送的信息 list sentPost

```golang
id            // id(自动生成)  --index
postID        // 消息id
title         // 标题（用于存uuid）
payload       // 内容
toChain       // 目的链id
sender        // 发送者chainAddr
receiver      // 接收者chainAddr
sentDate      // 发送时间
link_sender    // link list
link_receiver  // link list
```



#### 超时的信息 list timedoutPost

```golang
id            // id(自动生成)  --index
title         // 标题（用于存uuid）
toChain       // 目的链id
sender        // 发送者chainAddr
receiver      // 接收者chainAddr
sentDate      // 发送日期

link_sender    // link list
link_receiver  // link list
```



## 数据标记

- ```PLAIN:``` 开头表现ipfs hash值

- ```@IPFS:``` 开头表现ipfs hash值

- ```@@SM2:``` 开头表示SM2密钥加密数据

- ```@@SM4:``` 开头表示SM4密钥加密数据




## 内置链表设计

模块数据中，为了进行索引，对需要索引的字段构建链表，目前只支持单字段索引。以下以kvzoo为例，对owner建立链表，以便按owner进行检索所有相同owner的数据。

- 数据示例

| owner | zooKey | zooValue | lastDate | link_owner |
| ------- | ------ | -------- | -------- | ------------ |
| C1      | K1     | xxx      | xxx      | @@LINK:$     |
| C2      | K2     | xxx      | xxx      | @@LINK:$     |
| C3      | K1     | xxx      | xxx      | @@LINK:$     |
| C1      | K2     | xxx      | xxx      | C1/K1        |
| C2      | K1     | xxx      | xxx      | C2/K2        |

- 链表头数据

| key               | value |
| ----------------- | ----- |
| @@LINK:owner:C1 | C1/K2 |
| @@LINK:owner:C2 | C2/K1 |
| @@LINK:owner:C3 | C3/K1 |

如上示例，link_owner字段存储上一个数据的index key（kvzoo的index是```owner/zooKey```），其中```@@LINK:$```表示链表结尾。链表头数据也存储在kvzoo的store中，key自定义为不同owner取值，value为链表头数据的index key（可以由```x/zoo/types/keys.go```中```KeyPrefix（）```定义）。



### 数据检索

数据检索时，例如检索```owner=="C1"```的数据：

1. 先从store中检索到```@@LINK:owner:C1```值，即为链表头，作为当前指针P
2. 使用当前指针P作为key在store中检索数据并保存value，将link_owner替换P，重复步骤2，直到```link_owner=="@@LINK:$"```



### 数据添加

数据添加时:

1. 生成新数据的index key

2. 检索store，如果已存在，则结束，否则继续步骤3

3. 生成```@@LINK:owner:???```链表头key

4. 使用链表头key在store里检索，

   如未检索到，说明是新的值，添加新表头数据：key为链表头key，value为新数据的index key，设置P为```@@LINK:$```；

   如检索到表头，保存表头value为P，将表头value改写为新数据的index key

5. 添加新数据，其中新数据的link_owner为P

