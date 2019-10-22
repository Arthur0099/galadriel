cd solidity
truffle compile
cd ..
go test -v -run="Generate"
go test -v -run="Local"
