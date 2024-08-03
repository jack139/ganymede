
## IBC测试



### 1. 初始化链

```bash
# 删除已有kingring
rm ~/.local/share/keyrings/cosmos.keyring

# 初始化链参数
../shell/init_chain testchain1 n1 node1
../shell/init_chain testchain2 n2 node2
../shell/init_chain testchain3 n3 node3
```



### 2. 需设置的端口（多链测试）

```bash
#config/config.toml
proxy_app = "tcp://127.0.0.1:26658" # TCP or UNIX socket address of the ABCI application,
[rpc]
laddr = "tcp://127.0.0.1:26657"
pprof_laddr = "localhost:6060"
[p2p]
laddr = "tcp://0.0.0.0:26656"

#config/app.toml
pruning = "custom" # 会影响 application.db 存储
pruning-keep-recent = "500"
pruning-interval = "10"
[grpc]
address = "0.0.0.0:9090"
[grpc-web]
address = "0.0.0.0:9091"


#config/client.toml
node = "tcp://localhost:26657"
keyring-backend = "test"  # 客户端默认 keyring-backend

#tcp    127.0.0.1:6060   # pprof listen address 
#tcp    127.0.0.1:26657  # TCP or UNIX socket address for the RPC server to listen on
#tcp6   :::9090          # Address defines the gRPC server address to bind to.
#tcp6   :::9091          # Address defines the gRPC-web server address to bind to.
#tcp6   :::26656         # Address to listen for incoming connections

#    n1        n2       n3
#  26656     26666     26676
#  26657     26667     26677
#   6060      6061      6062
#   9090      9092      9094
#   9091      9093      9095
```



### 3. 启动链

```bash
build/ganymeded start --log_level warn --home n1
build/ganymeded start --log_level warn --home n2
build/ganymeded start --log_level warn --home n3
```



### 4. Relayer配置和启动 

```bash
# 初始化
rm -r ~/.relayer/*
rly config init

# 添加链
rly chains add --file config/test1.json testchain1
rly chains add --file config/test2.json testchain2
rly chains add --file config/test3.json testchain3

# 添加key
rly keys restore testchain1 testchain1_node1_relayer "mnemonic words here"
rly keys restore testchain2 testchain2_node2_relayer "mnemonic words here"
rly keys restore testchain3 testchain3_node3_relayer "mnemonic words here"

# 查询余额
rly q balance testchain1
rly q balance testchain2
rly q balance testchain3

# 添加 path: 1-2 2-3
rly paths new testchain1 testchain2 ibc-path-12
rly paths new testchain2 testchain3 ibc-path-23

# 添加 channel, client, connection
# port-id 和 version 在 x/postoffice/types/keys.go 里定义, 需要保持一致
rly tx clients ibc-path-12 --client-tp 48h
rly tx connection ibc-path-12
rly tx channel ibc-path-12 --src-port postoffice --dst-port postoffice --order unordered --version postoffice-1

rly tx clients ibc-path-23 --client-tp 48h
rly tx connection ibc-path-23
rly tx channel ibc-path-23 --src-port postoffice --dst-port postoffice --order unordered --version postoffice-1

# 查看
rly chains list
rly paths list
rly config show
rly q channels testchain1
rly q channels testchain2
rly q channels testchain3

# 启动 path
rly start --time-threshold 8h
```



> client过期后，需重新建立 path , 要使用 --override

```bash
rly paths new testchain1 testchain2 ibc-path-12-2
rly tx client testchain1 testchain2 ibc-path-12-2 --override
rly tx client testchain2 testchain1 ibc-path-12-2 --override
rly tx connection ibc-path-12-2
rly tx channel ibc-path-12-2 --src-port postoffice --dst-port postoffice --order unordered --version postoffice-1
```



### 5. IBC测试

```bash
# 'channel-0' 从 'rly q channels testchain1' 返回结果中获取
build/ganymeded --home n1 tx postoffice send-ibc-post postoffice channel-0 "hello" "hello test" --from testchain1_node1_relayer --chain-id testchain1

build/ganymeded --home n1 q postoffice list-sent-post
build/ganymeded --home n2 q postoffice list-post
```



## 多节点测试



### 1. 初始化节点数据

```bash
# 此例，为 testchain1 增加节点 node4, chain_id相同
../shell/init_chain testchain1 n4 node4
```



### 2. 复制创世块文件

```bash
# 复制初始节点（n1）的创世块覆盖新节点的创世块文件
cp n1/config/genesis.json n4/config/
```



### 3. 复制用户密钥
> 如果复制节点（n4）不提交交易，可以忽略此步骤
> 生产环境，复制节点间需要建立密钥同步机制（NFS 或 rsync）

```bash
# 复制初始节点（n1）的 keyring 和 sm2 数据
cp n1/keyring-test/* n4/keyring-test/
cp -r n1/sm2/* n4/
```



### 4. 修改 persistent_peers

```bash
# n4/config/config.toml
persistent_peers = "id@127.0.0.1:26656"
# 私有网络或单机测试需要设置这个
addr_book_strict = false
```

其中 id 的获取方法：

```bash
build/ganymeded tendermint show-node-id --home n1
```



### 5. 修改端口配置

```bash
# n4/config/config.toml
# n4/config/app.toml
# n4/config/client.toml

#    n1     n4  
#  26656  26686 
#  26657  26687 
#   6060   6063 
#   9090   9096 
#   9091   9097 
```



### 6. 启动第二节点

```bash
build/ganymeded start --log_level warn --home n4
```



## 管理区块链



### 启动 http API

```bash
build/ganymeded --home n1 http --yaml config/settings1.yaml
build/ganymeded --home n2 http --yaml config/settings2.yaml
build/ganymeded --home n3 http --yaml config/settings3.yaml
```
