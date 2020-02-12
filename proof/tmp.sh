cd ../solidity
truffle compile
cd ../
go test -v -run="Generate"
cd proof
go test -v -run="TestInnerProductContractVerify"
