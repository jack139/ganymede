# 开发环境准备



## 生成docker image

```bash
docker build -f dockerfile-ubuntu . -t ignite_0_27
```



## 运行交互docker

```bash
docker run --name ignitebox --expose 5000 -v /data_disk/VMs/shared_disk/docker:/opt/dev -it ignite_0_27 /bin/bash

docker container start ignitebox
docker container attach ignitebox

# docker中显示中文
export LANG=C.UTF-8
```



## 本地调试

```
// for debug only
replace github.com/cosmos/cosmos-sdk => ../../source/cosmos-sdk-0.47.3
```




## 代码框架



### 1. 区块链框架

```bash
ignite scaffold chain github.com/jack139/ganymede/ganymede
```


### 2. Ganymede模块

```bash
#ignite scaffold message add-account addr --module ganymede
ignite scaffold map users keyName userType name address phone accountNo ref regDate status lastDate link_status --index chainAddr --module ganymede

ignite scaffold query list-by-status status page:uint limit:uint -r users:Users --module ganymede
# 需手工加 repeated: proto/ganymede/ganymede/query.proto
ignite generate proto-go
```



### 3. Zoo模块

```bash
ignite scaffold module zoo --dep ganymede
ignite scaffold map kvzoo zooValue lastDate link_owner --index owner,zooKey --module zoo

ignite scaffold query list-by-owner owner page:uint limit:uint -r kvzoo:Kvzoo --module zoo
# 需手工加 repeated: proto/ganymede/zoo/query.proto
ignite generate proto-go
```



### 4. Exchange模块

```bash
ignite scaffold module exchange --dep ganymede
ignite scaffold list ask sender replier payload sentDate creator link_sender link_replier --no-message --module exchange
ignite scaffold list reply askId sender replier payload sentDate creator link_sender link_replier --no-message --module exchange

ignite scaffold message new-ask sender replier payload sentDate --module exchange
ignite scaffold message new-reply askId sender replier payload sentDate --module exchange

ignite scaffold query list-ask-by-sender sender page:uint limit:uint -r ask:Ask --module exchange
ignite scaffold query list-ask-by-replier replier page:uint limit:uint -r ask:Ask --module exchange
ignite scaffold query list-reply-by-sender sender page:uint limit:uint -r reply:Reply --module exchange
ignite scaffold query list-reply-by-replier replier page:uint limit:uint -r reply:Reply --module exchange
# 需手工加 repeated: proto/ganymede/exchange/query.proto
ignite generate proto-go
```



### 5. PostOffice模块

```bash
ignite scaffold module postoffice --ibc
ignite scaffold list post title payload fromChain sender receiver senderInfo sentDate link_sender link_receiver --no-message --module postoffice
ignite scaffold list sentPost postID title payload toChain sender receiver sentDate link_sender link_receiver --no-message --module postoffice
ignite scaffold list timedoutPost title toChain sender receiver sentDate link_sender link_receiver  --no-message --module postoffice

ignite scaffold packet ibcPost title content sentDate --ack postID --module postoffice
#proto/ganymede/postoffice/packet.proto 添加 creator
ignite generate proto-go

ignite scaffold query list-post-by-sender sender page:uint limit:uint -r post:Post --module postoffice
ignite scaffold query list-post-by-receiver receiver page:uint limit:uint -r post:Post --module postoffice
ignite scaffold query list-sent-by-sender sender page:uint limit:uint -r sentPost:SentPost --module postoffice
ignite scaffold query list-sent-by-receiver receiver page:uint limit:uint -r sentPost:SentPost --module postoffice
ignite scaffold query list-timeout-by-sender sender page:uint limit:uint -r timedoutPost:TimedoutPost --module postoffice
ignite scaffold query list-timeout-by-receiver receiver page:uint limit:uint -r timedoutPost:TimedoutPost --module postoffice
# 需手工加 repeated: proto/ganymede/postoffice/query.proto
ignite generate proto-go
```
