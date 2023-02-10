# PGC

PGC(Pretty Good Confidential Transaction System with Accountability)是一种基于账户实现的可监管加密交易系统，支持匿名性、机密性（隐藏交易金额）和可监管性。
普通用户可以使用该系统进行数字货币的隐私交易转账而不用担心信息被泄露，
监管方（拥有`Global Key`的机构）可以对所有交易进行监管、审计。同时，本系统还支持反洗钱、交易金额范围等（线性审计策略）策略审计功能，交易用户可向任何人证明一笔或多笔交易符合上述策略。

## 部署测试

在测试之前，需要安装`golang`和`ganache`，同时在`.env`中设置好相应的配置，可参照`.env.example`。

`.env`详细配置如下:

- TestMemo: 启动本地链使用的助记词
- DeployerKey: 合约部署者私钥
- AliceKey: 测试用户 alice 私钥
- BobKey: 测试用户 bob 私钥
- GlobalKey: 监管者私钥

启动本地区块链

```bash
ganache -m $your_memo
```

执行测试

```bash
cd test

go test -v -timeout 1h -run="TestPGCSystemContractTokenLocal"
```

执行结果

![](assets/pictures/test-output.png)
