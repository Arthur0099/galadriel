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
const IpverifierABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ipParams\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"u\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"params\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"uint256[2]\"},{\"name\":\"l\",\"type\":\"uint256[8]\"},{\"name\":\"r\",\"type\":\"uint256[8]\"},{\"name\":\"scalar\",\"type\":\"uint256[3]\"}],\"name\":\"verifyIPProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"gv\",\"type\":\"uint256[32]\"},{\"name\":\"hv\",\"type\":\"uint256[32]\"},{\"name\":\"p\",\"type\":\"uint256[2]\"},{\"name\":\"c\",\"type\":\"uint256\"},{\"name\":\"l\",\"type\":\"uint256[8]\"},{\"name\":\"r\",\"type\":\"uint256[8]\"},{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"verifyIPProofWithCustomParams\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IpverifierBin is the compiled bytecode used for deploying new contracts.
const IpverifierBin = `0x60806040523480156200001157600080fd5b5060405160208062001757833981018060405260208110156200003357600080fd5b810190808051906020019092919050505080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da8972246040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200010a57600080fd5b505afa1580156200011f573d6000803e3d6000fd5b505050506040513d60208110156200013657600080fd5b81019080805190602001909291905050506010141515620001bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6269732073697a65206e6f7420657175616c000000000000000000000000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633e9552256040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200024457600080fd5b505afa15801562000259573d6000803e3d6000fd5b505050506040513d60208110156200027057600080fd5b81019080805190602001909291905050506004141515620002f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6e756d626572206f66206c2c72206e6f7420657175616c00000000000000000081525060200191505060405180910390fd5b620003036200040f565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324d6147d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200038757600080fd5b505afa1580156200039c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620003c257600080fd5b81019080919050509050806000600281101515620003dc57fe5b60200201516000800181905550806001600281101515620003f957fe5b6020020151600060010181905550505062000431565b6040805190810160405280600290602082028038833980820191505090505090565b61131680620004416000396000f3fe608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632cd804211461007d5780632e52d606146101cd5780633e8d3764146101f857806382e1a5cb14610223578063b929d9081461027a578063c6a898c514610429575b600080fd5b34801561008957600080fd5b506101b360048036036102a08110156100a157600080fd5b8101908080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f82011690508083019250505050505091929192908061010001906008806020026040519081016040528092919082600860200280828437600081840152601f19601f82011690508083019250505050505091929192908061010001906008806020026040519081016040528092919082600860200280828437600081840152601f19601f820116905080830192505050505050919291929080606001906003806020026040519081016040528092919082600360200280828437600081840152601f19601f820116905080830192505050505050919291929050505061045b565b604051808215151515815260200191505060405180910390f35b3480156101d957600080fd5b506101e261063c565b6040518082815260200191505060405180910390f35b34801561020457600080fd5b5061020d610641565b6040518082815260200191505060405180910390f35b34801561022f57600080fd5b50610238610646565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561028657600080fd5b5061040f6004803603610aa081101561029e57600080fd5b8101908080610400019060208060200260405190810160405280929190826020800280828437600081840152601f19601f820116905080830192505050505050919291929080610400019060208060200260405190810160405280929190826020800280828437600081840152601f19601f820116905080830192505050505050919291929080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291908061010001906008806020026040519081016040528092919082600860200280828437600081840152601f19601f82011690508083019250505050505091929192908061010001906008806020026040519081016040528092919082600860200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291908035906020019092919050505061066c565b604051808215151515815260200191505060405180910390f35b34801561043557600080fd5b5061043e610894565b604051808381526020018281526020019250505060405180910390f35b6000610632600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ffe237f06040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016104006040518083038186803b1580156104e557600080fd5b505afa1580156104f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061040081101561051f57600080fd5b8101908091905050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663483767f06040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016104006040518083038186803b1580156105ac57600080fd5b505afa1580156105c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506104008110156105e657600080fd5b8101908091905050878560026003811015156105fe57fe5b6020020151888888600060038110151561061457fe5b602002015189600160038110151561062857fe5b602002015161066c565b9050949350505050565b600481565b601081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000610676611124565b61067e611124565b60008090505b601081101561074a5760408051908101604052808d836002026020811015156106a957fe5b602002015181526020018d600184600202016020811015156106c757fe5b602002015181525083826010811015156106dd57fe5b602002018190525060408051908101604052808c8360020260208110151561070157fe5b602002015181526020018c6001846002020160208110151561071f57fe5b6020020151815250828260108110151561073557fe5b60200201819052508080600101915050610684565b50610753611153565b60008090505b60048110156108275760408051908101604052808a8360020260088110151561077e57fe5b602002015181526020018a6001846002020160088110151561079c57fe5b60200201518152508260000151826004811015156107b657fe5b6020020181905250604080519081016040528089836002026008811015156107da57fe5b6020020151815260200189600184600202016008811015156107f857fe5b602002015181525082602001518260048110151561081257fe5b60200201819052508080600101915050610759565b508581604001818152505084816060018181525050610883838360408051908101604052808e600060028110151561085b57fe5b602002015181526020018e600160028110151561087457fe5b60200201518152508c856108a6565b935050505098975050505050505050565b60008060000154908060010154905082565b6000806108b28461092f565b90506108bc611189565b6108ef8260006040805190810160405290816000820154815260200160018201548152505061096b90919063ffffffff16565b905061092288888361091c8a61090e8b8861096b90919063ffffffff16565b610a1890919063ffffffff16565b88610aba565b9250505095945050505050565b600061096482604051602001808281526020019150506040516020818303038152906040528051906020012060019004610df9565b9050919050565b610973611189565b600182141561098457829050610a12565b600282141561099e576109978384610a18565b9050610a12565b6109a66111a3565b83600001518160006003811015156109ba57fe5b60200201818152505083602001518160016003811015156109d757fe5b602002018181525050828160026003811015156109f057fe5b6020020181815250506040826060836007600019fa1515610a1057600080fd5b505b92915050565b610a20611189565b610a286111c6565b8360000151816000600481101515610a3c57fe5b6020020181815250508360200151816001600481101515610a5957fe5b6020020181815250508260000151816002600481101515610a7657fe5b6020020181815250508260200151816003600481101515610a9357fe5b6020020181815250506040826080836006600019fa1515610ab357600080fd5b5092915050565b6000610ac46111e9565b60008090505b6004811015610c3a576000610b49856000015183600481101515610aea57fe5b602002015160000151866000015184600481101515610b0557fe5b602002015160200151876020015185600481101515610b2057fe5b602002015160000151886020015186600481101515610b3b57fe5b602002015160200151610e33565b90506000610b5682610e8a565b905081846040015184600481101515610b6b57fe5b60200201818152505080846060015184600481101515610b8757fe5b602002018181525050610c29610bce610ba98384610f5a90919063ffffffff16565b886020015186600481101515610bbb57fe5b602002015161096b90919063ffffffff16565b610c1b610c0c610be78687610f5a90919063ffffffff16565b8a6000015188600481101515610bf957fe5b602002015161096b90919063ffffffff16565b8a610a1890919063ffffffff16565b610a1890919063ffffffff16565b965050508080600101915050610aca565b50610c4361122b565b60008090505b6010811015610d335760008090505b6004811015610d25576000610c6f83836004610f96565b15610c9257846040015182600481101515610c8657fe5b60200201519050610cac565b846060015182600481101515610ca457fe5b602002015190505b6000821415610cd257808484601081101515610cc457fe5b602002018181525050610d17565b610cff610cfa828686601081101515610ce757fe5b6020020151610f5a90919063ffffffff16565b610df9565b8484601081101515610d0d57fe5b6020020181815250505b508080600101915050610c58565b508080600101915050610c49565b50610d3c611189565b610dcb610d6c610d5d87606001518860400151610f5a90919063ffffffff16565b8961096b90919063ffffffff16565b610dbd610d8f8860600151610d818d88610ff4565b61096b90919063ffffffff16565b610daf8960400151610da18f89611066565b61096b90919063ffffffff16565b610a1890919063ffffffff16565b610a1890919063ffffffff16565b905085600001518160000151148015610deb575085602001518160200151145b935050505095945050505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083811515610e2a57fe5b06915050919050565b6000610e8085858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060019004610df9565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905060008390506000811415610ec957600092505050610f55565b81811115610ee0578181811515610edc57fe5b0690505b600080600190506000849050600084905060005b600082141515610f2d578183811515610f0957fe5b04905083848202860383848402860380955081965082975083985050505050610ef4565b6000851215610f4a57846000038703975050505050505050610f55565b849750505050505050505b919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080801515610f8a57fe5b83850991505092915050565b6000806001830360019060020a02905060008090505b84811015610fd057600182908060020a820491505091508080600101915050610fac565b506000818616141515610fe7576001915050610fed565b60009150505b9392505050565b610ffc611189565b61100461122b565b60008090505b60108110156110525761102e848260108110151561102457fe5b6020020151610e8a565b828260108110151561103c57fe5b602002018181525050808060010191505061100a565b5061105d8482611066565b91505092915050565b61106e611189565b611076611189565b6110af83600060108110151561108857fe5b602002015185600060108110151561109c57fe5b602002015161096b90919063ffffffff16565b90506000600190505b60108110156111195761110a6110fb85836010811015156110d557fe5b602002015187846010811015156110e857fe5b602002015161096b90919063ffffffff16565b83610a1890919063ffffffff16565b915080806001019150506110b8565b508091505092915050565b610400604051908101604052806010905b61113d61124f565b8152602001906001900390816111355790505090565b61024060405190810160405280611168611269565b8152602001611175611269565b815260200160008152602001600081525090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b610900604051908101604052806111fe611298565b815260200161120b611298565b81526020016112186112c7565b81526020016112256112c7565b81525090565b61020060405190810160405280601090602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b610100604051908101604052806004905b61128261124f565b81526020019060019003908161127a5790505090565b610400604051908101604052806010905b6112b161124f565b8152602001906001900390816112a95790505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a72305820f130b9e0a4a9679d33bdfa16c90704a341ae56130fbaa782514fe60d8926be2b0029`

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

// VerifyIPProof is a free data retrieval call binding the contract method 0x2cd80421.
//
// Solidity: function verifyIPProof(uint256[2] p, uint256[8] l, uint256[8] r, uint256[3] scalar) constant returns(bool)
func (_Ipverifier *IpverifierCaller) VerifyIPProof(opts *bind.CallOpts, p [2]*big.Int, l [8]*big.Int, r [8]*big.Int, scalar [3]*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ipverifier.contract.Call(opts, out, "verifyIPProof", p, l, r, scalar)
	return *ret0, err
}

// VerifyIPProof is a free data retrieval call binding the contract method 0x2cd80421.
//
// Solidity: function verifyIPProof(uint256[2] p, uint256[8] l, uint256[8] r, uint256[3] scalar) constant returns(bool)
func (_Ipverifier *IpverifierSession) VerifyIPProof(p [2]*big.Int, l [8]*big.Int, r [8]*big.Int, scalar [3]*big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProof(&_Ipverifier.CallOpts, p, l, r, scalar)
}

// VerifyIPProof is a free data retrieval call binding the contract method 0x2cd80421.
//
// Solidity: function verifyIPProof(uint256[2] p, uint256[8] l, uint256[8] r, uint256[3] scalar) constant returns(bool)
func (_Ipverifier *IpverifierCallerSession) VerifyIPProof(p [2]*big.Int, l [8]*big.Int, r [8]*big.Int, scalar [3]*big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProof(&_Ipverifier.CallOpts, p, l, r, scalar)
}

// VerifyIPProofWithCustomParams is a free data retrieval call binding the contract method 0xb929d908.
//
// Solidity: function verifyIPProofWithCustomParams(uint256[32] gv, uint256[32] hv, uint256[2] p, uint256 c, uint256[8] l, uint256[8] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierCaller) VerifyIPProofWithCustomParams(opts *bind.CallOpts, gv [32]*big.Int, hv [32]*big.Int, p [2]*big.Int, c *big.Int, l [8]*big.Int, r [8]*big.Int, a *big.Int, b *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ipverifier.contract.Call(opts, out, "verifyIPProofWithCustomParams", gv, hv, p, c, l, r, a, b)
	return *ret0, err
}

// VerifyIPProofWithCustomParams is a free data retrieval call binding the contract method 0xb929d908.
//
// Solidity: function verifyIPProofWithCustomParams(uint256[32] gv, uint256[32] hv, uint256[2] p, uint256 c, uint256[8] l, uint256[8] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierSession) VerifyIPProofWithCustomParams(gv [32]*big.Int, hv [32]*big.Int, p [2]*big.Int, c *big.Int, l [8]*big.Int, r [8]*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProofWithCustomParams(&_Ipverifier.CallOpts, gv, hv, p, c, l, r, a, b)
}

// VerifyIPProofWithCustomParams is a free data retrieval call binding the contract method 0xb929d908.
//
// Solidity: function verifyIPProofWithCustomParams(uint256[32] gv, uint256[32] hv, uint256[2] p, uint256 c, uint256[8] l, uint256[8] r, uint256 a, uint256 b) constant returns(bool)
func (_Ipverifier *IpverifierCallerSession) VerifyIPProofWithCustomParams(gv [32]*big.Int, hv [32]*big.Int, p [2]*big.Int, c *big.Int, l [8]*big.Int, r [8]*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProofWithCustomParams(&_Ipverifier.CallOpts, gv, hv, p, c, l, r, a, b)
}
