## 安装ipfs

1. 第一个节点

```bash
mkdir ~/.ipfs/
cp shell/swarm.key ~/.ipfs/
ipfs init
```

查看```ipfs id```，修改```bootstrap.txt```内容里的id和ip

```bash
shell/config_ipfs
```

2. 其他节点

复制新的```bootstrap.txt```，然后执行 ```config_ipfs```

3. 在各个节点

```bash
shell/run_ipfs
```

4. 查看节点状态

```bash
ipfs swarm peers
```
