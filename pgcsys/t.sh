cd ../solidity
truffle compile
cd ..
go test -v -run="Generate"
cd pgcsys
go test -v -run="TestPGCSystemContractETH"
go test -v -run="TestPGCSystemContractToken"
