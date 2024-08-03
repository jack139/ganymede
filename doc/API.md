##  应用层API



###  一、 说明

​		应用层API作为可信业务区与区块链节点一起部署，提供给客户端调用，进行基础的区块链功能操作。



### 二、 概念和定义

#### 1. 节点

​		节点是区块链上的一个业务处理和存储的单元，是一个具有独立处理区块链业务的服务程序。节点可以是一台物理服务器，也可以是多个节点共用一个物理服务器，通过不同端口提供各自节点的功能。

#### 2. 链用户

​		链用户是具有提交区块链交易权限的用户，线下可定义为机构或个人。每个链用户通过一对密钥识别，同时使用此密钥进行数据的加密解密操作。链用户密钥由系统保存，新建时会返回用户密码串，链用户的密码串是恢复用户的唯一凭证，需要妥善保管。

#### 3. 交易和查询

​		交易是指会产生链上数据变动的操作，交易会在所有节点中传播，并记录在链上。查询是指对链上数据的检索，不会对链上数据进行修改和记录。

#### 4. 数据交换

​		不同链用户之间可以进行数据交换（“请求/授权”过程），整个过程会上链记录，所有交互过程链上可见，但均为密文不可读。“请求/授权”过程使用密钥协商进行加密密钥的传输，数据载荷使用对称加密，确保链上数据传输安全。数据交换支持跨链，即不同链上的链用户之间也可以进行“请求/授权”过程。

#### 5. 跨链传输

​		在不同chain id的链之间进行数据传输，称为跨链。目前API支持基于IBC协议同构跨链操作，对API用户透明。用户只需知晓目标链chain id和目标链用户的地址即可实现跨链数据传输和接收。

#### 6. 国密算法 

​		客户端API在应用层支持国密标准算法。API接口验签可选择使用SM2算法。数据交换密钥协商使用SM2和SM3算法、数据加密传输使用SM4算法。上链数据可选择使用SM4算法进行加密。

#### 7. IPFS

​		IPFS作为大块数据存储，不对客户端API用户直接开放，对客户端透明。当API中数据超过一定尺寸后，自动使用IPFS进行存储，API返回数据时进行自动转换。



### 三、 API提供的功能

#### 1. 交易接口

| 序号               | 接口功能                 | URI                        |
| :------------------: | -------------------------- | -------------------------- |
| 1 | 注册用户         | /tx/user/new     |
| 2 | 修改用户信息      | /tx/user/update |
| 3 | 审核用户 | /tx/user/audit |
| 4 | 添加KV数据 | /tx/kv/new |
| 5 | 修改KV数据 | /tx/kv/update |
| 6 | 删除KV数据 | /tx/kv/delete |
| 7 | 请求数据交换 | /tx/exchange/ask |
| 8 | 响应数据交换请求 | /tx/exchange/reply |
| 9 | 发送跨链传输数据 | /tx/post/send |
| 10 | 请求跨链数据交换 | /tx/post/ask |
| 11 | 响应跨链数据交换请求 | /tx/post/reply |



#### 2. 查询接口

| 序号               | 接口功能                 | URI                        |
| :------------------: | -------------------------- | -------------------------- |
| 1 | 查询用户清单 | /q/user/list |
| 2 | 查询用户信息     | /q/user/info       |
| 3 | 验证用户身份      | /q/user/verify       |
| 4 | 查询用户余额 | /q/bank/balance |
| 5 | 检索KV数据列表 | /q/kv/list |
| 6 | 检索KV数据 | /q/kv/show |
| 7 | 查询数据交换请求 | /q/exchange/ask/list |
| 8 | 查询数据交换响应结果 | /q/exhcange/reply/list |
| 9 | 查询数据交换响应结果内容 | /q/exhcange/reply/show |
| 10 | 查询已发送的跨链信息 | /q/post/sent/list |
| 11 | 查询发送超时的跨链信息 | /q/post/timeout/list |
| 12 | 查询收到的跨链信息 | /q/post/recv/list |
| 13 | 查询收到的跨链信息内容 | /q/post/recv/show |
| 14 | 查询指定交易数据 | /q/block/tx |
| 15 | 查询指定条件的交易数据 | /q/block/txs |
| 16 | 查询指定高度区块原始数据 | /q/block/height |



### 四、接口定义

#### 1. 全局接口定义

##### 输入参数

| 参数      | 类型   | 说明                          | 示例        |
| --------- | ------ | ----------------------------- | ----------- |
| appid | string | 应用渠道编号                  |             |
| version   | string | 版本号                        | 1 |
| sign_type | string | 签名算法，目前支持SHA256和SM2算法 | SHA256或SM2 |
| sign_data | string | 签名数据，具体算法见下文      |             |
| timestamp | int    | unix时间戳（秒）              |             |
| data      | json   | 接口数据，详见各接口定义      |             |

> 签名/验签算法：
>
> 1. appid和app_secret均从线下获得。
> 2. 筛选，获取参数键值对，剔除sign_data参数。data参数按key升序排列进行json序列化。
> 3. 排序，按key升序排序；data中json也按key升序排序。
> 4. 拼接，按排序好的顺序拼接请求参数。
>
> ```key1=value1&key2=value2&...&key=appSecret```，key=app_secret固定拼接在参数串末尾。
>
> 4. 签名，使用```sign_type```指定的算法进行加签获取二进制字节，使用 16进制进行编码Hex.encode得到签名串，然后base64编码。
> 5. 验签，对收到的参数按1-4步骤签名，比对得到的签名串与提交的签名串是否一致。



签名示例：

```json
请求参数：
{
    "appid":"4fcf3871f4a023712bec9ed44ee4b709",
    "version": "1",
    "sign_type": "SM2",
    "sign_data": "...",
    "timestamp":1681894715,
    "data": {
        "test1": "test1", 
        "atest2": "test2", 
        "Atest2": "test2"
    }
}

密钥：
appSecret="MjdjNGQxNGU3NjA1OWI0MGVmODIyN2FkOTEwYTViNDQzYTNjNTIyNSAgLQo="
SM2_privateKey="JShsBOJL0RgPAoPttEB1hgtPAvCikOl0V1oTOYL7k5U="

待加签串：
appid=4fcf3871f4a023712bec9ed44ee4b709&data={"Atest2":"test2","atest2":"test2","test1":"test1"}&sign_type=SM2&timestamp=1681894715&version=1&key=MjdjNGQxNGU3NjA1OWI0MGVmODIyN2FkOTEwYTViNDQzYTNjNTIyNSAgLQo=

SHA256加签结果：
"2c16865510262b1a88bac31f63b6110af862832da39b3d927d352d129e9843a1"

base64后结果：
"MmMxNjg2NTUxMDI2MmIxYTg4YmFjMzFmNjNiNjExMGFmODYyODMyZGEzOWIzZDkyN2QzNTJkMTI5ZTk4NDNhMQ=="

SM2加签结果（每次不同）：
"k/rHs2TbtIcIEJ6x0xQWr0ej1d+FPCAUqyiBZQ4MY0HFVXVw4KTl+C67CXxxf/WvJw8MRxfkINENmkzBCf+JJw=="

```



##### 返回结果

| 参数      | 类型    | 说明                                                         |
| --------- | ------- | ------------------------------------------------------------ |
| code      | int   | 状态代码，0 表示成功，非0 表示出错                                 |
| msg   | string | 成功时返回success；出错时，返回出错信息                                                     |
| data      | json    | 成功时返回结果数据，详见具体接口                |

返回示例

```json
{
    "code": 0, 
    "msg": "success", 
    "data": {
    }
}
```

全局出错代码

| 编码 | 说明             |
| ---- | ---------------- |
| 9000 | 签名错误         |
| 9001 | 缺少参数         |
| 9002 | 参数格式错误     |
| 9003 | 参数长度超过限制 |
| 9004 | 非法用户         |
| 9005 | key已存在        |
| 9090 | 同一用户同时提交  |
| 9099 | 交易提交失败     |



#### 2. 交易接口

##### 2.1 注册用户

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/user/new

请求方式

> POST

输入参数（data字段下）

| 参数      | 类型   | 必填 | 说明                                    |
| --------- | ------ | ---- | --------------------------------------- |
| key_name  | string | Y    | 链上key名称（限英文字母数字），长度<512 |
| user_type | string | Y    | 用户类型                                |
| name      | string |      | 用户名称，长度<512                      |
| acc_no    | string |      | 账户号码，长度<512                      |
| address   | string |      | 联系地址，长度<512                      |
| phone     | string |      | 联系电话，长度<512                      |
| ref       | string |      | 其他信息，长度<512                      |

> user_type 取值：
>
> USR 普通用户
>
> ORG 机构（需审核）
>
> 3PT 第三方（需审核）
>
> GOV 管理机构（需审核）
>
> 注意：用户状态如果不是```ACTIVE```将不能提交上链交易

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 用户链地址、密码字符串                  |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "key_name": "test2", 
        "user_type": "USR", 
        "phone": "1234567", 
        "name": "test22222"
    }, 
    "timestamp": 1682062252, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "09plyXAHSuTNFj8xdxlwUfIGpkuzHPyVHdbWLTlkG4CeR7iXPPXvQqmdo9udG3MhEVZdr6Y7id6REROH4T8dgw=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "chain_addr":"saturn1lpef9a8nqptjkgktxu2wqfpkve00htmde0z7vm",
        "mystery":"vote walk holiday keep impulse give jungle either modify cluster idea black jump screen rib crazy ocean dial gallery balance gaze labor later barrel",
        "txhash":"D4D8BF4269C147D73D3B0F1F05BA9D6725EDC9FC0FBC0CB5EBDB855D8650BC75"
    },
    "msg":"success"
}
```



##### 2.2 修改用户信息

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/user/update

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明               |
| ---------- | ------ | ---- | ------------------ |
| chain_addr | string | Y    | 用户的链地址       |
| name       | string |      | 用户名称，长度<512 |
| acc_no     | string |      | 账户号码，长度<512 |
| address    | string |      | 联系地址，长度<512 |
| phone      | string |      | 联系电话，长度<512 |
| ref        | string |      | 其他信息，长度<512 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块hash                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "chain_addr": "saturn1x2us24u67d2fyjz7mkwq50p0re2vs9n577xyrt", 
        "phone": "1234567", 
        "address": "xxxxxxxxxx", 
        "ref": "111111111"
    }, 
    "timestamp": 1682062357, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "ZBAemkIQaB4qbYy5eiGY3OCUDqHSzna6vrG6niQGRxqcVPLb/AqOOvqePr1pIogJ1SlLjMHzWuGy2Jn3n2mkXQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"34C76C1DAA392597F459571CC3663CC9BB785A4B10B6337B509381971B7BA4C5"
    },
    "msg":"success"
}
```



##### 2.3 审核用户

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/user/audit

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明       |
| ---------- | ------ | ---- | ---------- |
| chain_addr | string | Y    | 用户链地址 |
| status     | string | Y    | 用户状态   |

> 说明：用户状态如果不是```ACTIVE```将不能提交上链交易

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块hash                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "chain_addr": "saturn1x2us24u67d2fyjz7mkwq50p0re2vs9n577xyrt", 
        "status": "ACTIVE"
    }, 
    "timestamp": 1682062459, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "/EeW/UCWQcXju4H5pUfvYYwAtL473iZMh0cNR4GEDjv2O4TErIDvhw9Z3Dyg6ONtZI7qloTPVyQYbEiq0Xr0Ag=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"7DC725FF1EF37DB9EB8270E7673A135F90E58358A40ED52BD95286114AFCF5BA"
    },
    "msg":"success"
}
```



##### 2.4 添加KV数据

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/kv/new

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明                         |
| ---------- | ------ | ---- | ---------------------------- |
| owner_addr | string | Y    | 所有者链地址                 |
| key        | string | Y    | 键，限英文字母数字，长度<512 |
| value      | string | Y    | 值，无IPFS时长度<10240       |
| crypto     | bool   |      | 是否加密上链，默认为 false   |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块hash                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "owner_addr": "saturn1kdf98zx4ykv398w0ualaq6r9smcrcmh0d72uf9", 
        "key": "k2", 
        "value": "v2"
    }, 
    "timestamp": 1682316019, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "ZCdJI4QlP9aD15OS9sbWpiqt7qj4146SRAtyKUQf41Ruqsdsv4oI/UGcgGeM47buTOyjfrOFVJpzh5mzLVX7rQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"0376D597DB7F6136054D991B9A9B328414E72DCBB44FF741D6EF71E8F0F23759"
    },
    "msg":"success"
}
```



##### 2.5 修改KV数据

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/kv/update

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明                       |
| ---------- | ------ | ---- | -------------------------- |
| owner_addr | string | Y    | 所有者链地址               |
| key        | string | Y    | 键，长度<512               |
| value      | string | Y    | 值，无IPFS时长度<10240     |
| crypto     | bool   |      | 是否加密上链，默认为 false |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区hash                              |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "owner_addr": "saturn1kdf98zx4ykv398w0ualaq6r9smcrcmh0d72uf9", 
        "key": "k3", 
        "value": "value content3"
    }, 
    "timestamp": 1682316098, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "Jt1xkydEcKVgVa1MobvUt25bVJOrAlmuHmNUkPFrzNYP2AIvVOH1NVC/7L7qICHEcQgoXqSXWky99WGt9tu4Ag=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"C965B39F4B38C3CEC09BB29D3E4AA0AB886612E8AD5F9A2897471C4332B19153"
    },
    "msg":"success"
}
```



##### 2.6 删除KV数据

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/kv/delete

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明         |
| ---------- | ------ | ---- | ------------ |
| owner_addr | string | Y    | 所有者链地址 |
| key        | string | Y    | 键           |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块hash                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "owner_addr": "saturn1kdf98zx4ykv398w0ualaq6r9smcrcmh0d72uf9", 
        "key": "k3"
    }, 
    "timestamp": 1682315677, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "R7YO4GkeJ5O+Nh4gMJd5OHuwjfsRMGI853veEkHc+eR2ji49hkMqr+41kYadQuuPb6P/kgf2TWlvNBSLpyjpog=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"FCEFF9FF7C5AB1349C9EDA83ACCF76A671761F35BF314C323EA4D02FBF5CFD07"
    },
    "msg":"success"
}
```



##### 2.7 请求数据交换

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/exchange/ask

请求方式

> POST

输入参数（data字段下）

| 参数         | 类型   | 必填 | 说明                |
| ------------ | ------ | ---- | ------------------- |
| asker_addr   | string | Y    | 请求者链地址        |
| replier_addr | string | Y    | 响应者链地址        |
| payload      | string | Y    | 请求信息，长度<1024 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块hash                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "asker_addr": "saturn16mahue67jqhwn63n30dt29ngkcekwgn268naqk", 
        "replier_addr": "saturn12a48qgv4cx5htlakx3d0m97a7r89dtzqeqd6t6", 
        "payload": "text payload ask 3"
    }, 
    "timestamp": 1682647969, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NiW2xuGc11rPIS7vrj73PnpjV3AYYtX686d41FFR1NMO/jbj0DEN0n3KFQJpzTiGs6T5eZ7iGWGKNXVQiP2TCw=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"2112F1A7E5086CDB4F33038A74EED428759BE41AF220AEC75E81F8750245A2E1",
        "uuid":"c066c05b-3b66-4463-a0e0-74795ea2c1e9"
    },
    "msg":"success"
}
```



##### 2.8 响应数据交换请求

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/exchange/reply

请求方式

> POST

输入参数（data字段下）

| 参数         | 类型   | 必填 | 说明                                                |
| ------------ | ------ | ---- | --------------------------------------------------- |
| asker_addr   | string | Y    | 请求者链地址                                        |
| replier_addr | string | Y    | 响应者链地址                                        |
| ask_id       | uint   | Y    | 请求id                                              |
| reply        | bool   | Y    | 是否授权（响应）                                    |
| payload      | string |      | 响应数据，reply==false 时可不填，无IPFS时长度<10240 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块hash                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "asker_addr": "saturn16mahue67jqhwn63n30dt29ngkcekwgn268naqk", 
        "replier_addr": "saturn12a48qgv4cx5htlakx3d0m97a7r89dtzqeqd6t6", 
        "payload": "text payload reply 3", 
        "ask_id": 8, 
        "reply": true
    }, 
    "timestamp": 1682648045, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "A+n0fhlJTu4eaYsZ1GcEgTYPA7fxMUW7pVWrzJJCIKO/q3O+YX73Js38Llj1g2fXjaQfjXDhac2aZKWBaqLzPg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"960871088BD7BFC34E8AFC65416176E510672B4CD1A65CE9764A76D23F6F2371",
        "uuid":"c066c05b-3b66-4463-a0e0-74795ea2c1e9"
    },
    "msg":"success"
}
```



##### 2.9 发送跨链传输数据

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/post/send

请求方式

> POST

输入参数（data字段下）

| 参数         | 类型   | 必填 | 说明                         |
| ------------ | ------ | ---- | ---------------------------- |
| sender_addr  | string | Y    | 发送者链地址                 |
| target_addr  | string | Y    | 接收者链地址                 |
| post_channel | string | Y    | 跨链通道                     |
| payload      | string | Y    | 数据内容，无IPFS时长度<10240 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块hash                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "sender_addr": "saturn18e6mjnm2892r4x924skc8qqcak7xa8dswmukcg", 
        "target_addr": "saturn1xchhq2f0glhtn3slzvqec9fcxef389r9uwak24", 
        "post_channel": "channel-0", 
        "payload": "post test content"
    }, 
    "timestamp": 1683527693, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "HVDZ/52lTlmVYbytU4oT6/pVO2zyW2ZA/5+pAgElMbJIeS5sk8zEVasJ6MfrcgVLCdtP3DFkGPvmEN/FIHAu1g=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"4D932AE6F8517F3D501E7A2A72F88F69D766705648FD7A64C64F89C8DDC10DBA",
        "uuid":"3c2f8ad9-b5c5-4cc8-b9c0-b22a5066c741"
    },
    "msg":"success"
}
```



##### 2.10 请求跨链数据交换

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/post/ask

请求方式

> POST

输入参数（data字段下）

| 参数         | 类型   | 必填 | 说明                |
| ------------ | ------ | ---- | ------------------- |
| asker_addr   | string | Y    | 请求者链地址        |
| replier_addr | string | Y    | 响应者链地址        |
| post_channel | string | Y    | 跨链通道            |
| payload      | string | Y    | 请求信息，长度<1024 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块hash                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "asker_addr": "saturn16mahue67jqhwn63n30dt29ngkcekwgn268naqk", 
        "replier_addr": "saturn12a48qgv4cx5htlakx3d0m97a7r89dtzqeqd6t6", 
        "payload": "text payload ask 3",
        "post_channel": "channel-0"
    }, 
    "timestamp": 1682647969, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NiW2xuGc11rPIS7vrj73PnpjV3AYYtX686d41FFR1NMO/jbj0DEN0n3KFQJpzTiGs6T5eZ7iGWGKNXVQiP2TCw=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"2112F1A7E5086CDB4F33038A74EED428759BE41AF220AEC75E81F8750245A2E1",
        "uuid":"3c2f8ad9-b5c5-4cc8-b9c0-b22a5066c741"
    },
    "msg":"success"
}
```



##### 2.11 响应跨链数据交换请求

请求URL

> http://\<host\>:\<port\>/api/\<version\>/tx/post/reply

请求方式

> POST

输入参数（data字段下）

| 参数         | 类型   | 必填 | 说明                                                |
| ------------ | ------ | ---- | --------------------------------------------------- |
| asker_addr   | string | Y    | 请求者链地址                                        |
| replier_addr | string | Y    | 响应者链地址                                        |
| post_channel | string | Y    | 跨链通道                                            |
| ask_post_id  | uint   | Y    | 请求id，从```q/post/recv/list```里获得              |
| reply        | bool   | Y    | 是否授权（响应）                                    |
| payload      | string |      | 响应数据，reply==false 时可不填，无IPFS时长度<10240 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块hash                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "asker_addr": "saturn16mahue67jqhwn63n30dt29ngkcekwgn268naqk", 
        "replier_addr": "saturn12a48qgv4cx5htlakx3d0m97a7r89dtzqeqd6t6", 
        "payload": "text payload reply 3", 
        "post_channel": "channel-0",
        "ask_post_id": 8, 
        "reply": true
    }, 
    "timestamp": 1682648045, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "A+n0fhlJTu4eaYsZ1GcEgTYPA7fxMUW7pVWrzJJCIKO/q3O+YX73Js38Llj1g2fXjaQfjXDhac2aZKWBaqLzPg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txhash":"960871088BD7BFC34E8AFC65416176E510672B4CD1A65CE9764A76D23F6F2371",
        "uuid":"3c2f8ad9-b5c5-4cc8-b9c0-b22a5066c741"
    },
    "msg":"success"
}
```





#### 3. 查询接口



##### 3.1 查询用户清单

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/user/list

请求方式

> POST

输入参数（data字段下）

| 参数  | 类型 | 必填 | 说明               |
| ----- | ---- | ---- | ------------------ |
| status | string |    | 按状态值过滤     |
| page  | uint |      | 第几页，缺省为1    |
| limit | uint |      | 每页数量，缺省为50 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 用户清单数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {}, 
    "timestamp": 1682062692, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "Tc1a9pspT7gqmql4VoEC0V7iIXDySmyWVIYpina/sOIPmax48o2E0sifmlAAY37uvSsADLkTh7ORTRFYzWJojQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "user_list":[
            {
                "chain_addr":"saturn19zj7t6f90gpg0dzza2ac9s0ywref4r74nntyth",
                "key_name":"test1234",
                "last_date":"2023-04-20 14:08:13",
                "reg_date":"2023-04-20 14:08:12",
                "status":"ACTIVE",
                "user_type":"USR"
            },
            {
                "chain_addr":"saturn1lpef9a8nqptjkgktxu2wqfpkve00htmde0z7vm",
                "key_name":"test2",
                "last_date":"2023-04-21 15:30:55",
                "reg_date":"2023-04-21 15:30:52",
                "status":"ACTIVE",
                "user_type":"USR"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.2 查询用户信息

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/user/info

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明       |
| ---------- | ------ | ---- | ---------- |
| chain_addr | string | Y    | 用户链地址 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 用户信息数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "chain_addr": "saturn1x2us24u67d2fyjz7mkwq50p0re2vs9n577xyrt"
    }, 
    "timestamp": 1682062726, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "8QTunjVquVPt0rZOYWR5v4XIuuDV/UFToLWQzhXNiyj+1eU7wT+m9X85PJYhtw2cZBDPFF6+bdl32cEfDVISbw=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "user":{
            "acc_no":"",
            "address":"xxxxxxxxxx",
            "chain_addr":"saturn1x2us24u67d2fyjz7mkwq50p0re2vs9n577xyrt",
            "key_name":"test12345",
            "last_date":"2023-04-21 15:34:20",
            "name":"",
            "phone":"1234567",
            "ref":"111111111",
            "reg_date":"2023-04-20 14:09:32",
            "status":"ACTIVE",
            "user_type":"USR"
        }
    },
    "msg":"success"
}
```



##### 3.3 验证用户身份

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/user/verify

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明                         |
| ---------- | ------ | ---- | ---------------------------- |
| chain_addr | string | Y    | 用户链地址                   |
| mystery    | string | Y    | 密码单词串，以空格分隔       |
| positions  | string | Y    | 密码单词顺序位置，以空格分隔 |

> 1. 密码单词为用户注册时返回的密码串
> 2. 顺序位置为单词在原始密码串中的顺序
> 3. 顺序位置值为整数，最小为1，最大为原始密码串单词个数
> 4. 调用时，至少提供3个密码单词和位置

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 是否验证通过                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "chain_addr": "saturn1ht290msugezfhmwayyq3qhdd0sgrh6w76u4w66", 
        "mystery": "art where frozen news", 
        "positions": "1 3 2 24"
    }, 
    "timestamp": 1682392948, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "4xcttbnRiY5bdxrEbCwC5/Xb/Bf5SySSLVCO2v4mQk26vP61OTVMZUiK8TG4dj3xOYf+4F2GkQQjnmxlocAb1Q=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "verified":true
    },
    "msg":"success"
}
```



##### 3.4 查询用户余额

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/bank/balance

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明       |
| ---------- | ------ | ---- | ---------- |
| chain_addr | string | Y    | 用户链地址 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 余额信息                                |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "chain_addr": "saturn1x2us24u67d2fyjz7mkwq50p0re2vs9n577xyrt"
    }, 
    "timestamp": 1682062891, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "26I1BU0KQ/ddWghCs9npBqvtayOA1xAS3XSnkhsZd4McT40X5F1EZNEB1xW5Zhj9FRqQE/Ib7GQaGRle25szDA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "balance":{
            "amount":"1",
            "denom":"credit"
        }
    },
    "msg":"success"
}
```



##### 3.5 检索KV数据列表

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/kv/list

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明                       |
| ---------- | ------ | ---- | -------------------------- |
| owner_addr | string |      | 所有者链地址，为空则不过滤 |
| page       | uint   |      | 第几页，缺省为1            |
| limit      | uint   |      | 每页数量，缺省为50         |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | KV数据列表                              |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "owner_addr": "saturn1kdf98zx4ykv398w0ualaq6r9smcrcmh0d72uf9"
    }, 
    "timestamp": 1682306966, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "C6nkkgBF7+FADsaturnx0xY0q4E1+LiLMyxbUZedPVF0LjS0slEuqlLF5hlXStdWx8T4/bxvMAUZVnGmgdCPSHeDfQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "kv_list":[
            {
                "key":"key_test2",
                "last_date":"2023-04-24 10:42:24",
                "owner_addr":"saturn1kdf98zx4ykv398w0ualaq6r9smcrcmh0d72uf9",
                "value":"value content"
            },
            {
                "key":"key_test",
                "last_date":"2023-04-24 10:17:40",
                "owner_addr":"saturn1kdf98zx4ykv398w0ualaq6r9smcrcmh0d72uf9",
                "value":"value content"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.6 检索KV数据

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/kv/show

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明                       |
| ---------- | ------ | ---- | -------------------------- |
| owner_addr | string | Y    | 所有者链地址，为空则不过滤 |
| key        | string | Y    | 键，为空则不过滤           |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | KV数据                                  |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "owner_addr": "saturn1zc2uv5hm0tj9xyfg07e7tws5uwm04xfgq3xu5c", 
        "key": "key_test2"
    }, 
    "timestamp": 1682306920, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "8sFeaWOca0q5gpHN1X3Of7kE+VyQ0ukymoLPEYTeJMz9gE79HhVHalEXLvB/nHMhh23nHGA3+xjKHtau93os1g=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "kv":{
            "key":"key_test2",
            "last_date":"2023-04-24 11:16:18",
            "owner_addr":"saturn1zc2uv5hm0tj9xyfg07e7tws5uwm04xfgq3xu5c",
            "value":"value content3"
        }
    },
    "msg":"success"
}
```



##### 3.7 查询数据交换请求

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/exchange/ask/list

请求方式

> POST

输入参数（data字段下）

| 参数         | 类型   | 必填 | 说明                       |
| ------------ | ------ | ---- | -------------------------- |
| asker_addr   | string |      | 请求者链地址，为空则不过滤 |
| replier_addr | string |      | 响应者链地址，为空则不过滤 |
| uuid         | string |      | 请求uuid，为空则不过滤     |
| page         | uint   |      | 第几页，缺省为1            |
| limit        | uint   |      | 每页数量，缺省为50         |

> 1. sender_addr 和 replier_addr 不能同时指定
> 1. uuid不能单独使用，需同时设置sender或replier

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 数据交换请求列表                        |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {}, 
    "timestamp": 1682660213, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "zUW2nebyLksaqazbCa5+sHE5NhjviMJHK/RZbujlZpglSQg0OXl+pL+YeV2SLoJD/ghdgQxWWoYqJHLMfJIYXg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "ask_list":[
            {
                "ask_id":"0",
                "asker_addr":"saturn1xk8vapqwpfsmwkg2snd8mzdh9rdpmulern7cea",
                "payload":"ask payload 测试",
                "replier_addr":"saturn1wejlrxdx6048hkwkqwm20ej9wrtr4xl4kwychd",
                "sent_date":"2023-05-11 11:04:55",
                "uuid":"c066c05b-3b66-4463-a0e0-74795ea2c1e9"
            },
            {
                "ask_id":"1",
                "asker_addr":"saturn1xk8vapqwpfsmwkg2snd8mzdh9rdpmulern7cea",
                "payload":"ask payload 测试 2222",
                "replier_addr":"saturn1wejlrxdx6048hkwkqwm20ej9wrtr4xl4kwychd",
                "sent_date":"2023-05-11 11:06:36",
                "uuid":"24cac5bc-24bb-4228-a4a9-ca3c44028220"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.8 查询数据交换响应结果

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/exchange/reply/list

请求方式

> POST

输入参数（data字段下）

| 参数         | 类型   | 必填 | 说明                       |
| ------------ | ------ | ---- | -------------------------- |
| asker_addr   | string |      | 请求者链地址，为空则不过滤 |
| replier_addr | string |      | 响应者链地址，为空则不过滤 |
| uuid         | string |      | 请求uuid，为空则不过滤     |
| page         | uint   |      | 第几页，缺省为1            |
| limit        | uint   |      | 每页数量，缺省为50         |

> 1. sender_addr 和 replier_addr 不能同时指定
> 2. uuid不能单独使用，需同时设置sender或replier

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 数据交换响应列表                        |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {}, 
    "timestamp": 1682662304, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "tUsGnwIRcmWY1mmzlMMqGKX7J6hAU/eEOZsJrplYWS5ZjYSk2bDfxLjJja8skyfGWXjGasm8NTNxDW50g3YQ5g=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "reply_list":[
            {
                "ask_id":"1",
                "asker_addr":"saturn1xk8vapqwpfsmwkg2snd8mzdh9rdpmulern7cea",
                "replier_addr":"saturn1wejlrxdx6048hkwkqwm20ej9wrtr4xl4kwychd",
                "reply":false,
                "reply_id":"0",
                "sent_date":"2023-05-11 11:08:11",
                "uuid":"24cac5bc-24bb-4228-a4a9-ca3c44028220"
            },
            {
                "ask_id":"0",
                "asker_addr":"saturn1xk8vapqwpfsmwkg2snd8mzdh9rdpmulern7cea",
                "replier_addr":"saturn1wejlrxdx6048hkwkqwm20ej9wrtr4xl4kwychd",
                "reply":true,
                "reply_id":"1",
                "sent_date":"2023-05-11 11:08:31",
                "uuid":"c066c05b-3b66-4463-a0e0-74795ea2c1e9"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.9 查询数据交换响应结果内容

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/exchange/reply/show

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明                     |
| ---------- | ------ | ---- | ------------------------ |
| asker_addr | string | Y    | 请求者链地址             |
| reply_id   | uint   | Y    | 响应id                   |
| decrypt    | bool   |      | 是否解密文本，默认不解密 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 数据交换响应数据                        |

> 返回data中，bool 字段 reply 标记请求是否被拒绝：true 发送了响应内容；false 拒绝了请求

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "asker_addr": "saturn16mahue67jqhwn63n30dt29ngkcekwgn268naqk", 
        "reply_id": 4, 
    }, 
    "timestamp": 1682649329, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "5/cHb4TFpegZsRsaturnJDkGZfuz4A9BDWtTZdIw1O6iOrOjEClqRMZtyFvUUpI2frhVfUpljtlMWj7uUgc4eXGebA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "reply":{
            "ask_id":"8",
            "payload":"{\"crypt\":\"IQHT5X8jW+yjRqeNY4eIlx/PllrdN1cPi6mSlIbWJoArwh4Kq/MoVrLfz5NQCfje1vFuBhaH9RtYQiypnYGU5/on\",\"pubkey\":\"ADb3tsdtthkiV7fnK0tKfNquyKm3m3fzaqs0w4bX0PqV\"}",
            "replier_addr":"saturn12a48qgv4cx5htlakx3d0m97a7r89dtzqeqd6t6",
            "reply":true,
            "reply_id":"4",
            "asker_addr":"saturn16mahue67jqhwn63n30dt29ngkcekwgn268naqk",
            "uuid":"c066c05b-3b66-4463-a0e0-74795ea2c1e9",
            "sent_date":"2023-04-28 10:14:09"
        }
    },
    "msg":"success"
}
```



##### 3.10 查询已发送的跨链信息

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/post/sent/list

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明                       |
| ----------- | ------ | ---- | -------------------------- |
| sender_addr | string |      | 发送者链地址，为空则不过滤 |
| target_addr | string |      | 接收者链地址，为空则不过滤 |
| uuid        | string |      | 请求uuid，为空则不过滤     |
| page        | uint   |      | 第几页，缺省为1            |
| limit       | uint   |      | 每页数量，缺省为50         |

> 1. sender_addr 和 target_addr 不能同时指定
> 2. uuid不能单独使用，需同时设置sender或replier

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 用户信息数据                            |

> ```post_type``` 取值3种：```POST```普通跨链数据；```EXCH:ASK```跨链交换请求；```EXCH:RPLY```跨链交换响应

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {}, 
    "timestamp": 1683697449, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "C0gqB7POrtZDD2EiO30pPYI88uTjVj/4m7zOVjhmj+NeepiPJwc8yBw6wJw9Ny9t6iYAU1PGf7bpQ6kmBfYqKA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "sent_list":[
            {
                "post_channel":"postoffice-channel-0",
                "post_id":"0",
                "post_type":"POST",
                "sender_addr":"saturn1ldw4j3mmktz86nc6l79zz7yrjwukj4f2e623vt",
                "sent_date":"2023-05-11 11:18:40",
                "sent_id":"0",
                "target_addr":"saturn1wejlrxdx6048hkwkqwm20ej9wrtr4xl4kwychd",
                "uuid":"3c2f8ad9-b5c5-4cc8-b9c0-b22a5066c741"
            },
            {
                "ask_post_id":"1",
                "asker_addr":"saturn1xk8vapqwpfsmwkg2snd8mzdh9rdpmulern7cea",
                "post_channel":"postoffice-channel-0",
                "post_type":"EXCH:RPLY",
                "replier_addr":"saturn1ldw4j3mmktz86nc6l79zz7yrjwukj4f2e623vt",
                "reply":true,
                "reply_post_id":"4",
                "sent_date":"2023-05-11 13:31:52",
                "sent_id":"4",
                "uuid":"b687bb0b-8889-4c6d-a22e-a5d8068e125b"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.11 查询发送超时的跨链信息

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/post/timeout/list

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明                       |
| ----------- | ------ | ---- | -------------------------- |
| sender_addr | string |      | 发送者链地址，为空则不过滤 |
| target_addr | string |      | 接收者链地址，为空则不过滤 |
| uuid        | string |      | 请求uuid，为空则不过滤     |
| page        | uint   |      | 第几页，缺省为1            |
| limit       | uint   |      | 每页数量，缺省为50         |

> 1. sender_addr 和 target_addr 不能同时指定
> 2. uuid不能单独使用，需同时设置sender或replier

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 超时的跨链信息列表                      |

>  ```post_type``` 取值3种：```POST```普通跨链数据；```EXCH:ASK```跨链交换请求；```EXCH:RPLY```跨链交换响应

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {}, 
    "timestamp": 1683697566, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "II1KHPvjbaFxIbXwC1G2tRUb+aDORluJaIK+/zAVGas1oFw6Qa6C7RnaLGAG1KE2YxGUUqWy0Q97hv8IMPV60A=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "timeout_list":[
            {
                "post_channel":"postoffice-channel-0",
                "post_type":"POST",
                "sender_addr":"saturn1ldw4j3mmktz86nc6l79zz7yrjwukj4f2e623vt",
                "target_addr":"saturn1wejlrxdx6048hkwkqwm20ej9wrtr4xl4kwychd",
                "timeout_date":"2023-05-11 11:18:15",
                "timeout_id":"0",
                "uuid":"513016b1-5df2-4171-8169-64b59d2fdf65"
            },
            {
                "asker_addr":"saturn1xk8vapqwpfsmwkg2snd8mzdh9rdpmulern7cea",
                "post_channel":"postoffice-channel-0",
                "post_type":"EXCH:RPLY",
                "replier_addr":"saturn1ldw4j3mmktz86nc6l79zz7yrjwukj4f2e623vt",
                "timeout_date":"2023-05-11 13:31:22",
                "timeout_id":"1",
                "uuid":"b687bb0b-8889-4c6d-a22e-a5d8068e125b"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.12 查询收到的跨链信息

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/post/recv/list

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明                       |
| ----------- | ------ | ---- | -------------------------- |
| sender_addr | string |      | 发送者链地址，为空则不过滤 |
| target_addr | string |      | 接收者链地址，为空则不过滤 |
| uuid        | string |      | 请求uuid，为空则不过滤     |
| page        | uint   |      | 第几页，缺省为1            |
| limit       | uint   |      | 每页数量，缺省为50         |

> 1. sender_addr 和 target_addr 不能同时指定
> 2. uuid不能单独使用，需同时设置sender或replier

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 收到的跨链信息列表                      |

>  ```post_type``` 取值3种：```POST```普通跨链数据；```EXCH:ASK```跨链交换请求；```EXCH:RPLY```跨链交换响应

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {}, 
    "timestamp": 1683702304, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "lz8k8EzYvqpZOwgDQLqfZ33M4nbrB4Fa0rG8bK0hqKU7jEVtzwwRbiQPQEhnYTwujCwP/Gj0vYnyivOqsM/lYA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "recv_list":[
            {
                "post_channel":"postoffice-channel-0",
                "post_id":"0",
                "post_type":"POST",
                "recv_date":"2023-05-11 11:18:34",
                "sender_addr":"saturn1ldw4j3mmktz86nc6l79zz7yrjwukj4f2e623vt",
                "target_addr":"saturn1wejlrxdx6048hkwkqwm20ej9wrtr4xl4kwychd",
                "uuid":"3c2f8ad9-b5c5-4cc8-b9c0-b22a5066c741"
            },
            {
               "ask_post_id":"1",
                "asker_addr":"saturn1xk8vapqwpfsmwkg2snd8mzdh9rdpmulern7cea",
                "post_channel":"postoffice-channel-0",
                "post_id":"4",
                "post_type":"EXCH:RPLY",
                "recv_date":"2023-05-11 13:31:46",
                "replier_addr":"saturn1ldw4j3mmktz86nc6l79zz7yrjwukj4f2e623vt",
                "reply":true,
                "uuid":"b687bb0b-8889-4c6d-a22e-a5d8068e125b"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.13 查询收到的跨链信息内容

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/post/recv/show

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明                     |
| ----------- | ------ | ---- | ------------------------ |
| target_addr | string | Y    | 接收者链地址             |
| post_id     | uint   | Y    | 收到的消息id             |
| decrypt     | bool   |      | 是否解密文本，默认不解密 |

> 1. decrypt仅当跨链信息为数据交换响应结果时起作用（```post_type=="EXCH:RLY"```），因为其他条件时均为明文
> 2. ```post_id```来自```q/post/recv/list```的返回结果

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 收到的跨链信息内容                      |

>  ```post_type``` 取值3种：```POST```普通跨链数据；```EXCH:ASK```跨链交换请求；```EXCH:RPLY```跨链交换响应

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "target_addr": "saturn1xk8vapqwpfsmwkg2snd8mzdh9rdpmulern7cea", 
        "post_id": 4,
        "decrypt" : true
    }, 
    "timestamp": 1683702772, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "1S3pWbbueSpcggdMuxDufgM2oTwwqEO8LRy8J2GLMcfI6R1QopT8AVYtXlY8YpFGKFMN8GsTynFRkrAOi+adXA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "recv":{
            "ask_post_id":"1",
            "asker_addr":"saturn1xk8vapqwpfsmwkg2snd8mzdh9rdpmulern7cea",
            "payload":"post reply content 5555",
            "post_channel":"postoffice-channel-0",
            "post_id":"4",
            "post_type":"EXCH:RPLY",
            "recv_date":"2023-05-11 13:31:46",
            "replier_addr":"saturn1ldw4j3mmktz86nc6l79zz7yrjwukj4f2e623vt",
            "reply":true,
            "uuid":"b687bb0b-8889-4c6d-a22e-a5d8068e125b"
        }
    },
    "msg":"success"
}
```



##### 3.14 查询指定交易数据  

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/block/tx

请求方式

> POST

输入参数（data字段下）

| 参数   | 类型   | 必填 | 说明     |
| ------ | ------ | ---- | -------- |
| txhash | string | Y    | 交易hash |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 指定区块的原始区块数据                  |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "txhash": "A83D56175119567F48EB00C005E7C9504D72731A97132BA8C8B9277DFA10001E"
    }, 
    "timestamp": 1684136080, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "7GTKc0DZ+CrXb5GbNPP2+PQNSKIqj/sE2PVyF878Xam1p+kRDrXELri5PDgc6hmBDzfaF/PW6cm28h6T2kGGKA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "blcok":{
            "code":0,
            "codespace":"",
            "data":"12360A342F636F736D6F73746573742E67616E796D6564652E67616E796D6564652E4D73674372656174655573657273526573706F6E7365",
            "events":[
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"ZmVl",
                            "value":null
                        },
                        {
                            "index":true,
                            "key":"ZmVlX3BheWVy",
                            "value":"eWgxcWZjdGE5MDhmcmxwNWo5aGx3dWU1ZnA2cTV6eDRuZ2s4OWV6NGw="
                        }
                    ],
                    "type":"tx"
                },
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"YWNjX3NlcQ==",
                            "value":"eWgxcWZjdGE5MDhmcmxwNWo5aGx3dWU1ZnA2cTV6eDRuZ2s4OWV6NGwvMQ=="
                        }
                    ],
                    "type":"tx"
                },
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"c2lnbmF0dXJl",
                            "value":"NWx4Ym5ndGVTdFVQb0Q3a25BeGhWTlJuK2FPVDc5bW5lazFWd1p0Ni9QUjgxcVVrdlhMQnhpWFF5VGNrOHdrOWZiMHZhRkNNeTRLNXdFRWRabkluVHc9PQ=="
                        }
                    ],
                    "type":"tx"
                },
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"YWN0aW9u",
                            "value":"L2Nvc21vc3Rlc3QuZ2FueW1lZGUuZ2FueW1lZGUuTXNnQ3JlYXRlVXNlcnM="
                        }
                    ],
                    "type":"message"
                },
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"cmVjZWl2ZXI=",
                            "value":"eWgxbTNoMzB3bHZzZjhsbHJ1eHRwdWtkdnN5MGttMmt1bThjM3oybGo="
                        },
                        {
                            "index":true,
                            "key":"YW1vdW50",
                            "value":"MWNyZWRpdA=="
                        }
                    ],
                    "type":"coin_received"
                },
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"bWludGVy",
                            "value":"eWgxbTNoMzB3bHZzZjhsbHJ1eHRwdWtkdnN5MGttMmt1bThjM3oybGo="
                        },
                        {
                            "index":true,
                            "key":"YW1vdW50",
                            "value":"MWNyZWRpdA=="
                        }
                    ],
                    "type":"coinbase"
                },
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"c3BlbmRlcg==",
                            "value":"eWgxbTNoMzB3bHZzZjhsbHJ1eHRwdWtkdnN5MGttMmt1bThjM3oybGo="
                        },
                        {
                            "index":true,
                            "key":"YW1vdW50",
                            "value":"MWNyZWRpdA=="
                        }
                    ],
                    "type":"coin_spent"
                },
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"cmVjZWl2ZXI=",
                            "value":"eWgxbHFqOXd2NGpwM3Ywa2R2ZXVyd2s4cHluZ3NqZTN2OGFueDQ5ZGo="
                        },
                        {
                            "index":true,
                            "key":"YW1vdW50",
                            "value":"MWNyZWRpdA=="
                        }
                    ],
                    "type":"coin_received"
                },
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"cmVjaXBpZW50",
                            "value":"eWgxbHFqOXd2NGpwM3Ywa2R2ZXVyd2s4cHluZ3NqZTN2OGFueDQ5ZGo="
                        },
                        {
                            "index":true,
                            "key":"c2VuZGVy",
                            "value":"eWgxbTNoMzB3bHZzZjhsbHJ1eHRwdWtkdnN5MGttMmt1bThjM3oybGo="
                        },
                        {
                            "index":true,
                            "key":"YW1vdW50",
                            "value":"MWNyZWRpdA=="
                        }
                    ],
                    "type":"transfer"
                },
                {
                    "attributes":[
                        {
                            "index":true,
                            "key":"c2VuZGVy",
                            "value":"eWgxbTNoMzB3bHZzZjhsbHJ1eHRwdWtkdnN5MGttMmt1bThjM3oybGo="
                        }
                    ],
                    "type":"message"
                }
            ],
            "gas_used":"112625",
            "gas_wanted":"171850",
            "height":"183",
            "info":"",
            "logs":[
                {
                    "events":[
                        {
                            "attributes":[
                                {
                                    "key":"receiver",
                                    "value":"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj"
                                },
                                {
                                    "key":"amount",
                                    "value":"1credit"
                                },
                                {
                                    "key":"receiver",
                                    "value":"saturn1lqj9wv4jp3v0kdveurwk8pyngsje3v8anx49dj"
                                },
                                {
                                    "key":"amount",
                                    "value":"1credit"
                                }
                            ],
                            "type":"coin_received"
                        },
                        {
                            "attributes":[
                                {
                                    "key":"spender",
                                    "value":"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj"
                                },
                                {
                                    "key":"amount",
                                    "value":"1credit"
                                }
                            ],
                            "type":"coin_spent"
                        },
                        {
                            "attributes":[
                                {
                                    "key":"minter",
                                    "value":"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj"
                                },
                                {
                                    "key":"amount",
                                    "value":"1credit"
                                }
                            ],
                            "type":"coinbase"
                        },
                        {
                            "attributes":[
                                {
                                    "key":"action",
                                    "value":"/cosmostest.ganymede.ganymede.MsgCreateUsers"
                                },
                                {
                                    "key":"sender",
                                    "value":"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj"
                                }
                            ],
                            "type":"message"
                        },
                        {
                            "attributes":[
                                {
                                    "key":"recipient",
                                    "value":"saturn1lqj9wv4jp3v0kdveurwk8pyngsje3v8anx49dj"
                                },
                                {
                                    "key":"sender",
                                    "value":"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj"
                                },
                                {
                                    "key":"amount",
                                    "value":"1credit"
                                }
                            ],
                            "type":"transfer"
                        }
                    ],
                    "log":"",
                    "msg_index":0
                }
            ],
            "raw_log":"[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj\"},{\"key\":\"amount\",\"value\":\"1credit\"},{\"key\":\"receiver\",\"value\":\"saturn1lqj9wv4jp3v0kdveurwk8pyngsje3v8anx49dj\"},{\"key\":\"amount\",\"value\":\"1credit\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj\"},{\"key\":\"amount\",\"value\":\"1credit\"}]},{\"type\":\"coinbase\",\"attributes\":[{\"key\":\"minter\",\"value\":\"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj\"},{\"key\":\"amount\",\"value\":\"1credit\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmostest.ganymede.ganymede.MsgCreateUsers\"},{\"key\":\"sender\",\"value\":\"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"saturn1lqj9wv4jp3v0kdveurwk8pyngsje3v8anx49dj\"},{\"key\":\"sender\",\"value\":\"saturn1m3h30wlvsf8llruxtpukdvsy0km2kum8c3z2lj\"},{\"key\":\"amount\",\"value\":\"1credit\"}]}]}]",
            "timestamp":"2023-05-12T06:47:18Z",
            "tx":{
                "@type":"/cosmos.tx.v1beta1.Tx",
                "auth_info":{
                    "fee":{
                        "amount":[],
                        "gas_limit":"171850",
                        "granter":"",
                        "payer":""
                    },
                    "signer_infos":[
                        {
                            "mode_info":{
                                "single":{
                                    "mode":"SIGN_MODE_DIRECT"
                                }
                            },
                            "public_key":{
                                "@type":"/cosmos.crypto.secp256k1.PubKey",
                                "key":"Ald/FKzYuz7TktvAA4t4Ut7K+UdcvzORRLT7tZkVCK2K"
                            },
                            "sequence":"1"
                        }
                    ],
                    "tip":null
                },
                "body":{
                    "extension_options":[],
                    "memo":"",
                    "messages":[
                        {
                            "@type":"/cosmostest.ganymede.ganymede.MsgCreateUsers",
                            "accountNo":"",
                            "address":"\"A48kl9MLEouNNywRTqL4w7ipiXwsDZGjV4UvyxxjVle/\"",
                            "chainAddr":"saturn1lqj9wv4jp3v0kdveurwk8pyngsje3v8anx49dj",
                            "creator":"saturn1qfcta908frlp5j9hlwue5fp6q5zx4ngk89ez4l",
                            "keyName":"test5",
                            "lastDate":"2023-05-12 14:47:18",
                            "linkStatus":"",
                            "name":"@@SM4:63iLRpxoOJwOk8/ViM7fWBXSOrwzyinq2uFBTH/AuMiVMp5LcijtcwPwig7bKqb6c3mzJ4+R6jn2pvADs0sCFxwplPduao5NShmzOuTYrXTPy9LyvBvnGYnab9LzscTN1vzNC8nG+hlwGdfylBda/stmazQG1fkTbAHN9vuRechbEUE7Z3jBaimG++6/RS1n8+ACC9ZWhCcb+3wXzuLSy/uJtvG8ZxsqgpWJUHBYx210I9YP/B+lEn94ggrqCudM1oeHwfmdKayUAoK5vuG4yR/mlre4bAsJ6J4rnp0aguJ3/CRIX/0GmuUZs0zvL8Ff",
                            "phone":"",
                            "ref":"",
                            "regDate":"2023-05-12 14:47:18",
                            "status":"ACTIVE",
                            "userType":"USR"
                        }
                    ],
                    "non_critical_extension_options":[],
                    "timeout_height":"0"
                },
                "signatures":[
                    "5lxbngteStUPoD7knAxhVNRn+aOT79mnek1VwZt6/PR81qUkvXLBxiXQyTck8wk9fb0vaFCMy4K5wEEdZnInTw=="
                ]
            },
            "txhash":"A83D56175119567F48EB00C005E7C9504D72731A97132BA8C8B9277DFA10001E"
        }
    },
    "msg":"success"
}
```



##### 3.15 查询指定条件的交易数据 

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/block/txs

请求方式

> POST

输入参数（data字段下）

| 参数         | 类型   | 必填 | 说明               |
| ------------ | ------ | ---- | ------------------ |
| creator_addr | string |      | 交易提交者         |
| tx_action    | string |      | 交易类型           |
| page         | uint   |      | 第几页，缺省为1    |
| limit        | uint   |      | 每页数量，缺省为50 |

> ```creator_addr```和```tx_action```至少使用一个
>
> ```tx_action```可设置的值：```user/new```, ```user/update```, ```kv/new```, ```kv/update```, ```kv/delete```, ```exchange/ask```, ```exchange/reply```, ```post/send```

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 指定区块的原始区块数据                  |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "tx_action": "kv/new"
    }, 
    "timestamp": 1685585814, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "0D0r5WbFoDH7cNcoQu9PbXze0w0x46TkuA/w8B8YB0tE7I7gFGVFVCRaqXADRV5v5d8Ds+w8gU2GyquNfd4oxw=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "txs":{
            "count":"1",
            "limit":"50",
            "page_number":"1",
            "page_total":"1",
            "total_count":"1",
            "txs":[
                {
                    "code":0,
                    "codespace":"",
                    "data":"12310A2F2F636F736D6F73746573742E67616E796D6564652E7A6F6F2E4D73674372656174654B767A6F6F526573706F6E7365",
                    "events":[
                        {
                            "attributes":[
                                {
                                    "index":true,
                                    "key":"ZmVl",
                                    "value":null
                                },
                                {
                                    "index":true,
                                    "key":"ZmVlX3BheWVy",
                                    "value":"eWgxczVhNzQ2OHYyYWt2NTIzbTM5d2NhMnB0cDhmOXdoMnhxMmZ6aHM="
                                }
                            ],
                            "type":"tx"
                        },
                        {
                            "attributes":[
                                {
                                    "index":true,
                                    "key":"YWNjX3NlcQ==",
                                    "value":"eWgxczVhNzQ2OHYyYWt2NTIzbTM5d2NhMnB0cDhmOXdoMnhxMmZ6aHMvMA=="
                                }
                            ],
                            "type":"tx"
                        },
                        {
                            "attributes":[
                                {
                                    "index":true,
                                    "key":"c2lnbmF0dXJl",
                                    "value":"QUNkaWtsNWNDSW9wS2VhWGNNaFo1OTFlOXFpM1VmRlVvNGZjbzd4YkZ6azREMFBvclVXVVNscG9EdWlIRjVFWEFhdkxmRFZFQThEenY5RDYva1o3YkE9PQ=="
                                }
                            ],
                            "type":"tx"
                        },
                        {
                            "attributes":[
                                {
                                    "index":true,
                                    "key":"YWN0aW9u",
                                    "value":"L2Nvc21vc3Rlc3QuZ2FueW1lZGUuem9vLk1zZ0NyZWF0ZUt2em9v"
                                }
                            ],
                            "type":"message"
                        }
                    ],
                    "gas_used":"59274",
                    "gas_wanted":"91824",
                    "height":"414",
                    "info":"",
                    "logs":[
                        {
                            "events":[
                                {
                                    "attributes":[
                                        {
                                            "key":"action",
                                            "value":"/cosmostest.ganymede.zoo.MsgCreateKvzoo"
                                        }
                                    ],
                                    "type":"message"
                                }
                            ],
                            "log":"",
                            "msg_index":0
                        }
                    ],
                    "raw_log":"[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmostest.ganymede.zoo.MsgCreateKvzoo\"}]}]}]",
                    "timestamp":"2023-05-19T06:12:16Z",
                    "tx":{
                        "@type":"/cosmos.tx.v1beta1.Tx",
                        "auth_info":{
                            "fee":{
                                "amount":[],
                                "gas_limit":"91824",
                                "granter":"",
                                "payer":""
                            },
                            "signer_infos":[
                                {
                                    "mode_info":{
                                        "single":{
                                            "mode":"SIGN_MODE_DIRECT"
                                        }
                                    },
                                    "public_key":{
                                        "@type":"/cosmos.crypto.secp256k1.PubKey",
                                        "key":"A3RWq+0GDHLs0iSVlVXRTofEo4r0G2jgGlFfbNTmCjAt"
                                    },
                                    "sequence":"0"
                                }
                            ],
                            "tip":null
                        },
                        "body":{
                            "extension_options":[],
                            "memo":"",
                            "messages":[
                                {
                                    "@type":"/cosmostest.ganymede.zoo.MsgCreateKvzoo",
                                    "creator":"saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs",
                                    "lastDate":"2023-05-19 14:12:17",
                                    "linkOwner":"",
                                    "owner":"saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs",
                                    "zooKey":"k1",
                                    "zooValue":"PLAIN:v1"
                                }
                            ],
                            "non_critical_extension_options":[],
                            "timeout_height":"0"
                        },
                        "signatures":[
                            "ACdikl5cCIopKeaXcMhZ591e9qi3UfFUo4fco7xbFzk4D0PorUWUSlpoDuiHF5EXAavLfDVEA8Dzv9D6/kZ7bA=="
                        ]
                    },
                    "txhash":"3A5A4245B78DE0978EF5979C1D9BF2ED18F7EE1B6DF1BAC305687F4EEF7942A8"
                }
            ]
        }
    },
    "msg":"success"
}
```



##### 3.16 查询指定高度区块原始数据

请求URL

> http://\<host\>:\<port\>/api/\<version\>/q/block/height

请求方式

> POST

输入参数（data字段下）

| 参数   | 类型   | 必填 | 说明     |
| ------ | ------ | ---- | -------- |
| height | string | Y    | 区块高度 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 指定区块的原始区块数据                  |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SM2", 
    "data": {
        "height": "985"
    }, 
    "timestamp": 1682062837, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "IzTMUOd6NS3gBI/ijILHA5JG8KFn1v1i4UivZAOtxHi1pdjvIIItbPllN9jXlD8+rsRz4bFKPCEXL4LUzM3+rQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "blcok":{
            "block":{
                "data":{
                    "txs":null
                },
                "evidence":{
                    "evidence":null
                },
                "header":{
                    "app_hash":"79DCD8D5A48AC730C49090243F6D2DA891B184A646588C11D2D290A76D3B15FC",
                    "chain_id":"testchain",
                    "consensus_hash":"048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
                    "data_hash":"E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                    "evidence_hash":"E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                    "height":"985",
                    "last_block_id":{
                        "hash":"3BDE0403E93A2B729A9293105E95D7F298C38FBFF9B5B9A04A4AB57AB6D59B27",
                        "parts":{
                            "hash":"8AE5FF8F267ECFB2F3CAD15C9F42A2E6B20020BE3A1AF1D3AE84A59ABBFE018C",
                            "total":1
                        }
                    },
                    "last_commit_hash":"8B9A9F66A2F0F4AD5CA797FA9F47143D0ACFE01F5A89B0379143FA5E719364B7",
                    "last_results_hash":"E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                    "next_validators_hash":"514432D330E85117C6DE8B9C0BA35128DE3179031B72E97910573B7D90E4BB70",
                    "proposer_address":"73D2C6EAD104FDDDAF36A5856613D5B5E28EDCCF",
                    "time":"2023-04-20T06:43:15.088340963Z",
                    "validators_hash":"514432D330E85117C6DE8B9C0BA35128DE3179031B72E97910573B7D90E4BB70",
                    "version":{
                        "block":"11"
                    }
                },
                "last_commit":{
                    "block_id":{
                        "hash":"3BDE0403E93A2B729A9293105E95D7F298C38FBFF9B5B9A04A4AB57AB6D59B27",
                        "parts":{
                            "hash":"8AE5FF8F267ECFB2F3CAD15C9F42A2E6B20020BE3A1AF1D3AE84A59ABBFE018C",
                            "total":1
                        }
                    },
                    "height":"984",
                    "round":0,
                    "signatures":[
                        {
                            "block_id_flag":2,
                            "signature":"ROaZWaisbZbfDUYWd5dyGA0o2lSoHvqpslpWgvsqiBuJj+71FlU2R3+BMhUTJWYNoucLzDh2W5IJfI0ab0TiDQ==",
                            "timestamp":"2023-04-20T06:43:15.088340963Z",
                            "validator_address":"73D2C6EAD104FDDDAF36A5856613D5B5E28EDCCF"
                        }
                    ]
                }
            },
            "block_id":{
                "hash":"8BD6A02E96AD96031B729A84CA40C602D42E04677FF1E2F60B3C990729311D83",
                "parts":{
                    "hash":"C00C3C657D0FDA4ECF5D2B56D25DCD3B1724AE6CD999888404477C446BD366D6",
                    "total":1
                }
            }
        }
    },
    "msg":"success"
}
```

