// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IpverifierABI is the input ABI used to generate the binding from.
const IpverifierABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ipParams\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"u\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"params\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"uint256[2]\"},{\"name\":\"l\",\"type\":\"uint256[10]\"},{\"name\":\"r\",\"type\":\"uint256[10]\"},{\"name\":\"scalar\",\"type\":\"uint256[3]\"}],\"name\":\"verifyIPProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"gv\",\"type\":\"uint256[64]\"},{\"name\":\"hv\",\"type\":\"uint256[64]\"},{\"name\":\"p\",\"type\":\"uint256[2]\"},{\"name\":\"c\",\"type\":\"uint256\"},{\"name\":\"l\",\"type\":\"uint256[10]\"},{\"name\":\"r\",\"type\":\"uint256[10]\"},{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"verifyIPProofWithCustomParams\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IpverifierBin is the compiled bytecode used for deploying new contracts.
const IpverifierBin = `0x60806040523480156200001157600080fd5b5060405160208062001759833981018060405260208110156200003357600080fd5b810190808051906020019092919050505080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da8972246040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200010a57600080fd5b505afa1580156200011f573d6000803e3d6000fd5b505050506040513d60208110156200013657600080fd5b81019080805190602001909291905050506020141515620001bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6269732073697a65206e6f7420657175616c000000000000000000000000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633e9552256040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200024457600080fd5b505afa15801562000259573d6000803e3d6000fd5b505050506040513d60208110156200027057600080fd5b81019080805190602001909291905050506005141515620002f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6e756d626572206f66206c2c72206e6f7420657175616c00000000000000000081525060200191505060405180910390fd5b620003036200040f565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324d6147d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200038757600080fd5b505afa1580156200039c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620003c257600080fd5b81019080919050509050806000600281101515620003dc57fe5b60200201516000800181905550806001600281101515620003f957fe5b6020020151600060010181905550505062000431565b6040805190810160405280600290602082028038833980820191505090505090565b61131880620004416000396000f3fe608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631982df651461007d5780632e52d6061461022e5780633e8d37641461025957806382e1a5cb14610284578063a93eef6e146102db578063c6a898c51461042b575b600080fd5b34801561008957600080fd5b5061021460048036036113208110156100a157600080fd5b810190808061080001906040806020026040519081016040528092919082604060200280828437600081840152601f19601f82011690508083019250505050505091929192908061080001906040806020026040519081016040528092919082604060200280828437600081840152601f19601f820116905080830192505050505050919291929080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291908035906020019092919050505061045d565b604051808215151515815260200191505060405180910390f35b34801561023a57600080fd5b50610243610685565b6040518082815260200191505060405180910390f35b34801561026557600080fd5b5061026e61068a565b6040518082815260200191505060405180910390f35b34801561029057600080fd5b5061029961068f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102e757600080fd5b5061041160048036036103208110156102ff57600080fd5b8101908080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f8201169050808301925050505050509192919290806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f820116905080830192505050505050919291929080606001906003806020026040519081016040528092919082600360200280828437600081840152601f19601f82011690508083019250505050505091929192905050506106b5565b604051808215151515815260200191505060405180910390f35b34801561043757600080fd5b50610440610896565b604051808381526020018281526020019250505060405180910390f35b6000610467611126565b61046f611126565b60008090505b602081101561053b5760408051908101604052808d8360020260408110151561049a57fe5b602002015181526020018d600184600202016040811015156104b857fe5b602002015181525083826020811015156104ce57fe5b602002018190525060408051908101604052808c836002026040811015156104f257fe5b602002015181526020018c6001846002020160408110151561051057fe5b6020020151815250828260208110151561052657fe5b60200201819052508080600101915050610475565b50610544611155565b60008090505b60058110156106185760408051908101604052808a83600202600a8110151561056f57fe5b602002015181526020018a60018460020201600a8110151561058d57fe5b60200201518152508260000151826005811015156105a757fe5b602002018190525060408051908101604052808983600202600a811015156105cb57fe5b602002015181526020018960018460020201600a811015156105e957fe5b602002015181525082602001518260058110151561060357fe5b6020020181905250808060010191505061054a565b508581604001818152505084816060018181525050610674838360408051908101604052808e600060028110151561064c57fe5b602002015181526020018e600160028110151561066557fe5b60200201518152508c856108a8565b935050505098975050505050505050565b600581565b602081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061088c600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ffe237f06040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016108006040518083038186803b15801561073f57600080fd5b505afa158015610753573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061080081101561077957600080fd5b8101908091905050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663483767f06040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016108006040518083038186803b15801561080657600080fd5b505afa15801561081a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061080081101561084057600080fd5b81019080919050508785600260038110151561085857fe5b6020020151888888600060038110151561086e57fe5b602002015189600160038110151561088257fe5b602002015161045d565b9050949350505050565b60008060000154908060010154905082565b6000806108b484610931565b90506108be61118b565b6108f18260006040805190810160405290816000820154815260200160018201548152505061096d90919063ffffffff16565b905061092488888361091e8a6109108b8861096d90919063ffffffff16565b610a1a90919063ffffffff16565b88610abc565b9250505095945050505050565b600061096682604051602001808281526020019150506040516020818303038152906040528051906020012060019004610dfb565b9050919050565b61097561118b565b600182141561098657829050610a14565b60028214156109a0576109998384610a1a565b9050610a14565b6109a86111a5565b83600001518160006003811015156109bc57fe5b60200201818152505083602001518160016003811015156109d957fe5b602002018181525050828160026003811015156109f257fe5b6020020181815250506040826060836007600019fa1515610a1257600080fd5b505b92915050565b610a2261118b565b610a2a6111c8565b8360000151816000600481101515610a3e57fe5b6020020181815250508360200151816001600481101515610a5b57fe5b6020020181815250508260000151816002600481101515610a7857fe5b6020020181815250508260200151816003600481101515610a9557fe5b6020020181815250506040826080836006600019fa1515610ab557600080fd5b5092915050565b6000610ac66111eb565b60008090505b6005811015610c3c576000610b4b856000015183600581101515610aec57fe5b602002015160000151866000015184600581101515610b0757fe5b602002015160200151876020015185600581101515610b2257fe5b602002015160000151886020015186600581101515610b3d57fe5b602002015160200151610e35565b90506000610b5882610e8c565b905081846040015184600581101515610b6d57fe5b60200201818152505080846060015184600581101515610b8957fe5b602002018181525050610c2b610bd0610bab8384610f5c90919063ffffffff16565b886020015186600581101515610bbd57fe5b602002015161096d90919063ffffffff16565b610c1d610c0e610be98687610f5c90919063ffffffff16565b8a6000015188600581101515610bfb57fe5b602002015161096d90919063ffffffff16565b8a610a1a90919063ffffffff16565b610a1a90919063ffffffff16565b965050508080600101915050610acc565b50610c4561122d565b60008090505b6020811015610d355760008090505b6005811015610d27576000610c7183836005610f98565b15610c9457846040015182600581101515610c8857fe5b60200201519050610cae565b846060015182600581101515610ca657fe5b602002015190505b6000821415610cd457808484602081101515610cc657fe5b602002018181525050610d19565b610d01610cfc828686602081101515610ce957fe5b6020020151610f5c90919063ffffffff16565b610dfb565b8484602081101515610d0f57fe5b6020020181815250505b508080600101915050610c5a565b508080600101915050610c4b565b50610d3e61118b565b610dcd610d6e610d5f87606001518860400151610f5c90919063ffffffff16565b8961096d90919063ffffffff16565b610dbf610d918860600151610d838d88610ff6565b61096d90919063ffffffff16565b610db18960400151610da38f89611068565b61096d90919063ffffffff16565b610a1a90919063ffffffff16565b610a1a90919063ffffffff16565b905085600001518160000151148015610ded575085602001518160200151145b935050505095945050505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083811515610e2c57fe5b06915050919050565b6000610e8285858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060019004610dfb565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905060008390506000811415610ecb57600092505050610f57565b81811115610ee2578181811515610ede57fe5b0690505b600080600190506000849050600084905060005b600082141515610f2f578183811515610f0b57fe5b04905083848202860383848402860380955081965082975083985050505050610ef6565b6000851215610f4c57846000038703975050505050505050610f57565b849750505050505050505b919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080801515610f8c57fe5b83850991505092915050565b6000806001830360019060020a02905060008090505b84811015610fd257600182908060020a820491505091508080600101915050610fae565b506000818616141515610fe9576001915050610fef565b60009150505b9392505050565b610ffe61118b565b61100661122d565b60008090505b602081101561105457611030848260208110151561102657fe5b6020020151610e8c565b828260208110151561103e57fe5b602002018181525050808060010191505061100c565b5061105f8482611068565b91505092915050565b61107061118b565b61107861118b565b6110b183600060208110151561108a57fe5b602002015185600060208110151561109e57fe5b602002015161096d90919063ffffffff16565b90506000600190505b602081101561111b5761110c6110fd85836020811015156110d757fe5b602002015187846020811015156110ea57fe5b602002015161096d90919063ffffffff16565b83610a1a90919063ffffffff16565b915080806001019150506110ba565b508091505092915050565b610800604051908101604052806020905b61113f611251565b8152602001906001900390816111375790505090565b6102c06040519081016040528061116a61126b565b815260200161117761126b565b815260200160008152602001600081525090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b6111406040519081016040528061120061129a565b815260200161120d61129a565b815260200161121a6112c9565b81526020016112276112c9565b81525090565b61040060405190810160405280602090602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b610140604051908101604052806005905b611284611251565b81526020019060019003908161127c5790505090565b610800604051908101604052806020905b6112b3611251565b8152602001906001900390816112ab5790505090565b60a06040519081016040528060059060208202803883398082019150509050509056fea165627a7a72305820b5496eac11542db9351d0286f8f8671acea0b04a1c9c3344befc85cc96b379980029`

// DeployIpverifier deploys a new Ethereum contract, binding an instance of Ipverifier to it.
func DeployIpverifier(auth *bind.TransactOpts, backend bind.ContractBackend, params common.Address) (common.Address, *types.Transaction, *Ipverifier, error) {
	parsed, err := abi.JSON(strings.NewReader(IpverifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IpverifierBin), backend, params)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ipverifier{IpverifierCaller: IpverifierCaller{contract: contract}, IpverifierTransactor: IpverifierTransactor{contract: contract}, IpverifierFilterer: IpverifierFilterer{contract: contract}}, nil
}

// Ipverifier is an auto generated Go binding around an Ethereum contract.
type Ipverifier struct {
	IpverifierCaller     // Read-only binding to the contract
	IpverifierTransactor // Write-only binding to the contract
	IpverifierFilterer   // Log filterer for contract events
}

// IpverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpverifierSession struct {
	Contract     *Ipverifier       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpverifierCallerSession struct {
	Contract *IpverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IpverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpverifierTransactorSession struct {
	Contract     *IpverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IpverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpverifierRaw struct {
	Contract *Ipverifier // Generic contract binding to access the raw methods on
}

// IpverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpverifierCallerRaw struct {
	Contract *IpverifierCaller // Generic read-only contract binding to access the raw methods on
}

// IpverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpverifierTransactorRaw struct {
	Contract *IpverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpverifier creates a new instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifier(address common.Address, backend bind.ContractBackend) (*Ipverifier, error) {
	contract, err := bindIpverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipverifier{IpverifierCaller: IpverifierCaller{contract: contract}, IpverifierTransactor: IpverifierTransactor{contract: contract}, IpverifierFilterer: IpverifierFilterer{contract: contract}}, nil
}

// NewIpverifierCaller creates a new read-only instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifierCaller(address common.Address, caller bind.ContractCaller) (*IpverifierCaller, error) {
	contract, err := bindIpverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpverifierCaller{contract: contract}, nil
}

// NewIpverifierTransactor creates a new write-only instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IpverifierTransactor, error) {
	contract, err := bindIpverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpverifierTransactor{contract: contract}, nil
}

// NewIpverifierFilterer creates a new log filterer instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IpverifierFilterer, error) {
	contract, err := bindIpverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpverifierFilterer{contract: contract}, nil
}

// bindIpverifier binds a generic wrapper to an already deployed contract.
func bindIpverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IpverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipverifier *IpverifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ipverifier.Contract.IpverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipverifier *IpverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipverifier.Contract.IpverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipverifier *IpverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipverifier.Contract.IpverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipverifier *IpverifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ipverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipverifier *IpverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipverifier *IpverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipverifier.Contract.contract.Transact(opts, method, params...)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Ipverifier *IpverifierCaller) BitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ipverifier.contract.Call(opts, out, "bitSize")
	return *ret0, err
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Ipverifier *IpverifierSession) BitSize() (*big.Int, error) {
	return _Ipverifier.Contract.BitSize(&_Ipverifier.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Ipverifier *IpverifierCallerSession) BitSize() (*big.Int, error) {
	return _Ipverifier.Contract.BitSize(&_Ipverifier.CallOpts)
}

// IpParams is a free data retrieval call binding the contract method 0x82e1a5cb.
//
// Solidity: function ipParams() constant returns(address)
func (_Ipverifier *IpverifierCaller) IpParams(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ipverifier.contract.Call(opts, out, "ipParams")
	return *ret0, err
}

// IpParams is a free data retrieval call binding the contract method 0x82e1a5cb.
//
// Solidity: function ipParams() constant returns(address)
func (_Ipverifier *IpverifierSession) IpParams() (common.Address, error) {
	return _Ipverifier.Contract.IpParams(&_Ipverifier.CallOpts)
}

// IpParams is a free data retrieval call binding the contract method 0x82e1a5cb.
//
// Solidity: function ipParams() constant returns(address)
func (_Ipverifier *IpverifierCallerSession) IpParams() (common.Address, error) {
	return _Ipverifier.Contract.IpParams(&_Ipverifier.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Ipverifier *IpverifierCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ipverifier.contract.Call(opts, out, "n")
	return *ret0, err
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Ipverifier *IpverifierSession) N() (*big.Int, error) {
	return _Ipverifier.Contract.N(&_Ipverifier.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Ipverifier *IpverifierCallerSession) N() (*big.Int, error) {
	return _Ipverifier.Contract.N(&_Ipverifier.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() constant returns(uint256 X, uint256 Y)
func (_Ipverifier *IpverifierCaller) U(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Ipverifier.contract.Call(opts, out, "u")
	return *ret, err
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() constant returns(uint256 X, uint256 Y)
func (_Ipverifier *IpverifierSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Ipverifier.Contract.U(&_Ipverifier.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() constant returns(uint256 X, uint256 Y)
func (_Ipverifier *IpverifierCallerSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Ipverifier.Contract.U(&_Ipverifier.CallOpts)
}

// VerifyIPProof is a free data retrieval call binding the contract method 0xa93eef6e.
//
// Solidity: function verifyIPProof(uint256[2] p, uint256[10] l, uint256[10] r, uint256[3] scalar) constant returns(bool)
func (_Ipverifier *IpverifierCaller) VerifyIPProof(opts *bind.CallOpts, p [2]*big.Int, l [10]*big.Int, r [10]*big.Int, scalar [3]*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ipverifier.contract.Call(opts, out, "verifyIPProof", p, l, r, scalar)
	return *ret0, err
}

// VerifyIPProof is a free data retrieval call binding the contract method 0xa93eef6e.
//
// Solidity: function verifyIPProof(uint256[2] p, uint256[10] l, uint256[10] r, uint256[3] scalar) constant returns(bool)
func (_Ipverifier *IpverifierSession) VerifyIPProof(p [2]*big.Int, l [10]*big.Int, r [10]*big.Int, scalar [3]*big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProof(&_Ipverifier.CallOpts, p, l, r, scalar)
}

// VerifyIPProof is a free data retrieval call binding the contract method 0xa93eef6e.
//
// Solidity: function verifyIPProof(uint256[2] p, uint256[10] l, uint256[10] r, uint256[3] scalar) constant returns(bool)
func (_Ipverifier *IpverifierCallerSession) VerifyIPProof(p [2]*big.Int, l [10]*big.Int, r [10]*big.Int, scalar [3]*big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProof(&_Ipverifier.CallOpts, p, l, r, scalar)
}

// VerifyIPProofWithCustomParams is a free data retrieval call binding the contract method 0x1982df65.
//
// Solidity: function verifyIPProofWithCustomParams(uint256[64] gv, uint256[64] hv, uint256[2] p, uint256 c, uint256[10] l, uint256[10] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierCaller) VerifyIPProofWithCustomParams(opts *bind.CallOpts, gv [64]*big.Int, hv [64]*big.Int, p [2]*big.Int, c *big.Int, l [10]*big.Int, r [10]*big.Int, a *big.Int, b *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ipverifier.contract.Call(opts, out, "verifyIPProofWithCustomParams", gv, hv, p, c, l, r, a, b)
	return *ret0, err
}

// VerifyIPProofWithCustomParams is a free data retrieval call binding the contract method 0x1982df65.
//
// Solidity: function verifyIPProofWithCustomParams(uint256[64] gv, uint256[64] hv, uint256[2] p, uint256 c, uint256[10] l, uint256[10] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierSession) VerifyIPProofWithCustomParams(gv [64]*big.Int, hv [64]*big.Int, p [2]*big.Int, c *big.Int, l [10]*big.Int, r [10]*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProofWithCustomParams(&_Ipverifier.CallOpts, gv, hv, p, c, l, r, a, b)
}

// VerifyIPProofWithCustomParams is a free data retrieval call binding the contract method 0x1982df65.
//
// Solidity: function verifyIPProofWithCustomParams(uint256[64] gv, uint256[64] hv, uint256[2] p, uint256 c, uint256[10] l, uint256[10] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierCallerSession) VerifyIPProofWithCustomParams(gv [64]*big.Int, hv [64]*big.Int, p [2]*big.Int, c *big.Int, l [10]*big.Int, r [10]*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProofWithCustomParams(&_Ipverifier.CallOpts, gv, hv, p, c, l, r, a, b)
}
