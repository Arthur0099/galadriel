# Galadriel

Galadriel is an account-based implementation of a supervisable cryptographic transaction system that supports anonymity, confidentiality (hiding transaction amounts) and supervisability.

The system allows ordinary users to use the system to make private transactions in digital currencies without fear of information leakage.

The regulator (the institution that owns `Global Key`) can supervise and audit all transactions. Also, this system supports anti-money laundering, transaction amount range, and other (linear audit strategy) strategy audit functions, and transaction users can prove to anyone that one or more transactions comply with the above strategy.


## build & test

Before testing, you need to install `golang` and `ganache` and set the corresponding configuration in `.env`, which can be found in `.env.example`.

The detailed configuration of `.env` is as follows:

- TestMemo: Mnemonics used to start the local chain
- DeployerKey: Contract Deployer Private Keys
- AliceKey: Test user alice private key
- BobKey: Test user bob private key
- GlobalKey: Regulator Private Key

### start local chain

```bash
ganache -m $your_memo
```

### run test

```bash
cd test

go test -v -timeout 1h -run="TestPGCSystemContractTokenLocal"
```

### test result

![](assets/pictures/test-output.png)
