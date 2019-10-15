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

// PublicparamsABI is the input ABI used to generate the binding from.
const PublicparamsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gVector\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gBase\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hVector\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"u\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"getG\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGVector\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[64]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHVector\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[64]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getN\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PublicparamsBin is the compiled bytecode used for deploying new contracts.
const PublicparamsBin = `0x60806040523480156200001157600080fd5b5060016004600001819055506002600460010181905550620000776040805190810160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c6700000000000081525062000133640100000000026401000000009004565b600080820151816000015560208201518160010155905050620000de6040805190810160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c6700000000000081525062000133640100000000026401000000009004565b6002600082015181600001556020820151816001015590505060008001546006600001819055506000600101546006600101819055506200012d62000224640100000000026401000000009004565b620005ec565b6200013d6200058c565b6000620001d7836040516020018082805190602001908083835b6020831015156200017e578051825260208201915060208101905060208303925062000157565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060019004620003dc64010000000002620007c5176401000000009004565b90506200021c816004604080519081016040529081600082015481526020016001820154815250506200041764010000000002620007ff179091906401000000009004565b915050919050565b600080600090505b6020811015620002fd576200027f81604051602001808281526020019150506040516020818303038152906040528051906020012060019004620003dc64010000000002620007c5176401000000009004565b9150620002c4826004604080519081016040529081600082015481526020016001820154815250506200041764010000000002620007ff179091906401000000009004565b600882602081101515620002d457fe5b60020201600082015181600001556020820151816001015590505080806001019150506200022c565b5060008090505b6020811015620003d8576200035a60208201604051602001808281526020019150506040516020818303038152906040528051906020012060019004620003dc64010000000002620007c5176401000000009004565b91506200039f826004604080519081016040529081600082015481526020016001820154815250506200041764010000000002620007ff179091906401000000009004565b604882602081101515620003af57fe5b600202016000820151816000015560208201518160010155905050808060010191505062000304565b5050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156200040e57fe5b06915050919050565b620004216200058c565b60018214156200043457829050620004db565b60028214156200046157620004598384620004e1640100000000026401000000009004565b9050620004db565b6200046b620005a6565b83600001518160006003811015156200048057fe5b60200201818152505083602001518160016003811015156200049e57fe5b60200201818152505082816002600381101515620004b857fe5b6020020181815250506040826060836007600019fa1515620004d957600080fd5b505b92915050565b620004eb6200058c565b620004f5620005c9565b83600001518160006004811015156200050a57fe5b60200201818152505083602001518160016004811015156200052857fe5b60200201818152505082600001518160026004811015156200054657fe5b60200201818152505082602001518160036004811015156200056457fe5b6020020181815250506040826080836006600019fa15156200058557600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b6109da80620005fc6000396000f3fe6080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304c09ce9146100e05780630c00f8a01461013357806324d6147d146101895780632e52d606146101dc5780633e8d3764146102075780633e95522514610232578063483767f01461025d57806352c9b965146102b057806356e2736c146102e257806382529fdb14610338578063b8c9d3651461038b578063c6a898c5146103bd578063da897224146103ef578063e2179b8e1461041a578063ffe237f01461044c575b600080fd5b3480156100ec57600080fd5b506100f561049f565b6040518082600260200280838360005b83811015610120578082015181840152602081019050610105565b5050505090500191505060405180910390f35b34801561013f57600080fd5b5061016c6004803603602081101561015657600080fd5b81019080803590602001909291905050506104f1565b604051808381526020018281526020019250505060405180910390f35b34801561019557600080fd5b5061019e61051a565b6040518082600260200280838360005b838110156101c95780820151818401526020810190506101ae565b5050505090500191505060405180910390f35b3480156101e857600080fd5b506101f161056d565b6040518082815260200191505060405180910390f35b34801561021357600080fd5b5061021c610572565b6040518082815260200191505060405180910390f35b34801561023e57600080fd5b50610247610577565b6040518082815260200191505060405180910390f35b34801561026957600080fd5b50610272610580565b6040518082604060200280838360005b8381101561029d578082015181840152602081019050610282565b5050505090500191505060405180910390f35b3480156102bc57600080fd5b506102c5610619565b604051808381526020018281526020019250505060405180910390f35b3480156102ee57600080fd5b5061031b6004803603602081101561030557600080fd5b810190808035906020019092919050505061062b565b604051808381526020018281526020019250505060405180910390f35b34801561034457600080fd5b5061034d610654565b6040518082600260200280838360005b8381101561037857808201518184015260208101905061035d565b5050505090500191505060405180910390f35b34801561039757600080fd5b506103a06106a7565b604051808381526020018281526020019250505060405180910390f35b3480156103c957600080fd5b506103d26106b9565b604051808381526020018281526020019250505060405180910390f35b3480156103fb57600080fd5b506104046106cb565b6040518082815260200191505060405180910390f35b34801561042657600080fd5b5061042f6106d4565b604051808381526020018281526020019250505060405180910390f35b34801561045857600080fd5b506104616106e6565b6040518082604060200280838360005b8381101561048c578082015181840152602081019050610471565b5050505090500191505060405180910390f35b6104a761077f565b6104af61077f565b60008001548160006002811015156104c357fe5b6020020181815250506000600101548160016002811015156104e157fe5b6020020181815250508091505090565b60088160208110151561050057fe5b600202016000915090508060000154908060010154905082565b61052261077f565b61052a61077f565b60066000015481600060028110151561053f57fe5b60200201818152505060066001015481600160028110151561055d57fe5b6020020181815250508091505090565b600581565b602081565b60006005905090565b6105886107a1565b6105906107a1565b60008090505b6020811015610611576048816020811015156105ae57fe5b600202016000015482826002026040811015156105c757fe5b6020020181815250506048816020811015156105df57fe5b600202016001015482600183600202016040811015156105fb57fe5b6020020181815250508080600101915050610596565b508091505090565b60048060000154908060010154905082565b60488160208110151561063a57fe5b600202016000915090508060000154908060010154905082565b61065c61077f565b61066461077f565b60026000015481600060028110151561067957fe5b60200201818152505060026001015481600160028110151561069757fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b60006020905090565b60008060000154908060010154905082565b6106ee6107a1565b6106f66107a1565b60008090505b60208110156107775760088160208110151561071457fe5b6002020160000154828260020260408110151561072d57fe5b60200201818152505060088160208110151561074557fe5b6002020160010154826001836002020160408110151561076157fe5b60200201818152505080806001019150506106fc565b508091505090565b6040805190810160405280600290602082028038833980820191505090505090565b61080060405190810160405280604090602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156107f657fe5b06915050919050565b61080761094e565b6001821415610818578290506108a6565b60028214156108325761082b83846108ac565b90506108a6565b61083a610968565b836000015181600060038110151561084e57fe5b602002018181525050836020015181600160038110151561086b57fe5b6020020181815250508281600260038110151561088457fe5b6020020181815250506040826060836007600019fa15156108a457600080fd5b505b92915050565b6108b461094e565b6108bc61098b565b83600001518160006004811015156108d057fe5b60200201818152505083602001518160016004811015156108ed57fe5b602002018181525050826000015181600260048110151561090a57fe5b602002018181525050826020015181600360048110151561092757fe5b6020020181815250506040826080836006600019fa151561094757600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a72305820222c6c0a9ef08751befa4f66866ca191afe142a8a057ecca044954bc0151372e0029`

// DeployPublicparams deploys a new Ethereum contract, binding an instance of Publicparams to it.
func DeployPublicparams(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Publicparams, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicparamsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PublicparamsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Publicparams{PublicparamsCaller: PublicparamsCaller{contract: contract}, PublicparamsTransactor: PublicparamsTransactor{contract: contract}, PublicparamsFilterer: PublicparamsFilterer{contract: contract}}, nil
}

// Publicparams is an auto generated Go binding around an Ethereum contract.
type Publicparams struct {
	PublicparamsCaller     // Read-only binding to the contract
	PublicparamsTransactor // Write-only binding to the contract
	PublicparamsFilterer   // Log filterer for contract events
}

// PublicparamsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicparamsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicparamsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicparamsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicparamsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicparamsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicparamsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicparamsSession struct {
	Contract     *Publicparams     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicparamsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicparamsCallerSession struct {
	Contract *PublicparamsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PublicparamsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicparamsTransactorSession struct {
	Contract     *PublicparamsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PublicparamsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicparamsRaw struct {
	Contract *Publicparams // Generic contract binding to access the raw methods on
}

// PublicparamsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicparamsCallerRaw struct {
	Contract *PublicparamsCaller // Generic read-only contract binding to access the raw methods on
}

// PublicparamsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicparamsTransactorRaw struct {
	Contract *PublicparamsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicparams creates a new instance of Publicparams, bound to a specific deployed contract.
func NewPublicparams(address common.Address, backend bind.ContractBackend) (*Publicparams, error) {
	contract, err := bindPublicparams(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Publicparams{PublicparamsCaller: PublicparamsCaller{contract: contract}, PublicparamsTransactor: PublicparamsTransactor{contract: contract}, PublicparamsFilterer: PublicparamsFilterer{contract: contract}}, nil
}

// NewPublicparamsCaller creates a new read-only instance of Publicparams, bound to a specific deployed contract.
func NewPublicparamsCaller(address common.Address, caller bind.ContractCaller) (*PublicparamsCaller, error) {
	contract, err := bindPublicparams(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicparamsCaller{contract: contract}, nil
}

// NewPublicparamsTransactor creates a new write-only instance of Publicparams, bound to a specific deployed contract.
func NewPublicparamsTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicparamsTransactor, error) {
	contract, err := bindPublicparams(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicparamsTransactor{contract: contract}, nil
}

// NewPublicparamsFilterer creates a new log filterer instance of Publicparams, bound to a specific deployed contract.
func NewPublicparamsFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicparamsFilterer, error) {
	contract, err := bindPublicparams(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicparamsFilterer{contract: contract}, nil
}

// bindPublicparams binds a generic wrapper to an already deployed contract.
func bindPublicparams(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicparamsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publicparams *PublicparamsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Publicparams.Contract.PublicparamsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publicparams *PublicparamsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publicparams.Contract.PublicparamsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publicparams *PublicparamsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publicparams.Contract.PublicparamsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Publicparams *PublicparamsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Publicparams.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Publicparams *PublicparamsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Publicparams.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Publicparams *PublicparamsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Publicparams.Contract.contract.Transact(opts, method, params...)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) BitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "bitSize")
	return *ret0, err
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Publicparams *PublicparamsSession) BitSize() (*big.Int, error) {
	return _Publicparams.Contract.BitSize(&_Publicparams.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) BitSize() (*big.Int, error) {
	return _Publicparams.Contract.BitSize(&_Publicparams.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) G(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "g")
	return *ret, err
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.G(&_Publicparams.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.G(&_Publicparams.CallOpts)
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) GBase(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "gBase")
	return *ret, err
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) GBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GBase(&_Publicparams.CallOpts)
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) GBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GBase(&_Publicparams.CallOpts)
}

// GVector is a free data retrieval call binding the contract method 0x0c00f8a0.
//
// Solidity: function gVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) GVector(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "gVector", arg0)
	return *ret, err
}

// GVector is a free data retrieval call binding the contract method 0x0c00f8a0.
//
// Solidity: function gVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) GVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GVector(&_Publicparams.CallOpts, arg0)
}

// GVector is a free data retrieval call binding the contract method 0x0c00f8a0.
//
// Solidity: function gVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) GVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GVector(&_Publicparams.CallOpts, arg0)
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) GetBitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getBitSize")
	return *ret0, err
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() constant returns(uint256)
func (_Publicparams *PublicparamsSession) GetBitSize() (*big.Int, error) {
	return _Publicparams.Contract.GetBitSize(&_Publicparams.CallOpts)
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) GetBitSize() (*big.Int, error) {
	return _Publicparams.Contract.GetBitSize(&_Publicparams.CallOpts)
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() constant returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetG(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getG")
	return *ret0, err
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() constant returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetG() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetG(&_Publicparams.CallOpts)
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() constant returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetG() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetG(&_Publicparams.CallOpts)
}

// GetGVector is a free data retrieval call binding the contract method 0xffe237f0.
//
// Solidity: function getGVector() constant returns(uint256[64])
func (_Publicparams *PublicparamsCaller) GetGVector(opts *bind.CallOpts) ([64]*big.Int, error) {
	var (
		ret0 = new([64]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getGVector")
	return *ret0, err
}

// GetGVector is a free data retrieval call binding the contract method 0xffe237f0.
//
// Solidity: function getGVector() constant returns(uint256[64])
func (_Publicparams *PublicparamsSession) GetGVector() ([64]*big.Int, error) {
	return _Publicparams.Contract.GetGVector(&_Publicparams.CallOpts)
}

// GetGVector is a free data retrieval call binding the contract method 0xffe237f0.
//
// Solidity: function getGVector() constant returns(uint256[64])
func (_Publicparams *PublicparamsCallerSession) GetGVector() ([64]*big.Int, error) {
	return _Publicparams.Contract.GetGVector(&_Publicparams.CallOpts)
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() constant returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetH(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getH")
	return *ret0, err
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() constant returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetH() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetH(&_Publicparams.CallOpts)
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() constant returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetH() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetH(&_Publicparams.CallOpts)
}

// GetHVector is a free data retrieval call binding the contract method 0x483767f0.
//
// Solidity: function getHVector() constant returns(uint256[64])
func (_Publicparams *PublicparamsCaller) GetHVector(opts *bind.CallOpts) ([64]*big.Int, error) {
	var (
		ret0 = new([64]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getHVector")
	return *ret0, err
}

// GetHVector is a free data retrieval call binding the contract method 0x483767f0.
//
// Solidity: function getHVector() constant returns(uint256[64])
func (_Publicparams *PublicparamsSession) GetHVector() ([64]*big.Int, error) {
	return _Publicparams.Contract.GetHVector(&_Publicparams.CallOpts)
}

// GetHVector is a free data retrieval call binding the contract method 0x483767f0.
//
// Solidity: function getHVector() constant returns(uint256[64])
func (_Publicparams *PublicparamsCallerSession) GetHVector() ([64]*big.Int, error) {
	return _Publicparams.Contract.GetHVector(&_Publicparams.CallOpts)
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) GetN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getN")
	return *ret0, err
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() constant returns(uint256)
func (_Publicparams *PublicparamsSession) GetN() (*big.Int, error) {
	return _Publicparams.Contract.GetN(&_Publicparams.CallOpts)
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) GetN() (*big.Int, error) {
	return _Publicparams.Contract.GetN(&_Publicparams.CallOpts)
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() constant returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetU(opts *bind.CallOpts) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getU")
	return *ret0, err
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() constant returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetU() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetU(&_Publicparams.CallOpts)
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() constant returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetU() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetU(&_Publicparams.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) H(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "h")
	return *ret, err
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.H(&_Publicparams.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.H(&_Publicparams.CallOpts)
}

// HVector is a free data retrieval call binding the contract method 0x56e2736c.
//
// Solidity: function hVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) HVector(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "hVector", arg0)
	return *ret, err
}

// HVector is a free data retrieval call binding the contract method 0x56e2736c.
//
// Solidity: function hVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) HVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.HVector(&_Publicparams.CallOpts, arg0)
}

// HVector is a free data retrieval call binding the contract method 0x56e2736c.
//
// Solidity: function hVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) HVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.HVector(&_Publicparams.CallOpts, arg0)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "n")
	return *ret0, err
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Publicparams *PublicparamsSession) N() (*big.Int, error) {
	return _Publicparams.Contract.N(&_Publicparams.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) N() (*big.Int, error) {
	return _Publicparams.Contract.N(&_Publicparams.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) U(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "u")
	return *ret, err
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.U(&_Publicparams.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.U(&_Publicparams.CallOpts)
}
