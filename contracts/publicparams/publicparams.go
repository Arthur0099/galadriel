// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package publicparams

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
const PublicparamsABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"aggSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"u\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getG\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getH\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getU\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PublicparamsBin is the compiled bytecode used for deploying new contracts.
var PublicparamsBin = "0x608060405234801561001057600080fd5b50600160046000018190555060026004600101819055506100746040518060400160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c6700000000000081525061015d640100000000026401000000009004565b6000808201518160000155602082015181600101559050506100d96040518060400160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c6700000000000081525061015d640100000000026401000000009004565b6002600082015181600001556020820151816001015590505061013f6040518060400160405280601b81526020017f752067656e657261746f72206f6620696e6e657270726f64756374000000000081525061015d640100000000026401000000009004565b60066000820151816000015560208201518160010155905050610424565b6101656103c6565b60006101f8836040516020018082805190602001908083835b602083106101a1578051825260208201915060208101905060208303925061017e565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001206001900461024264010000000002610417176401000000009004565b905061023a8160046040518060400160405290816000820154815260200160018201548152505061027a6401000000000261044f179091906401000000009004565b915050919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161027157fe5b06915050919050565b6102826103c6565b600182141561029357829050610328565b60028214156102bc576102b5838461032e640100000000026401000000009004565b9050610328565b6102c46103e0565b8360000151816000600381106102d657fe5b6020020181815250508360200151816001600381106102f157fe5b602002018181525050828160026003811061030857fe5b6020020181815250506040826060836007600019fa61032657600080fd5b505b92915050565b6103366103c6565b61033e610402565b83600001518160006004811061035057fe5b60200201818152505083602001518160016004811061036b57fe5b60200201818152505082600001518160026004811061038657fe5b6020020181815250508260200151816003600481106103a157fe5b6020020181815250506040826080836006600019fa6103bf57600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b61061f80620004346000396000f3fe608060405234801561001057600080fd5b50600436106100c6576000357c0100000000000000000000000000000000000000000000000000000000900480637982ebcc1161008e5780637982ebcc146101b857806382529fdb146101d6578063b8c9d3651461021c578063c6a898c514610241578063da89722414610266578063e2179b8e14610284576100c6565b806304c09ce9146100cb57806324d6147d146101115780632e52d606146101575780633e8d37641461017557806352c9b96514610193575b600080fd5b6100d36102a9565b6040518082600260200280838360005b838110156100fe5780820151818401526020810190506100e3565b5050505090500191505060405180910390f35b6101196102f7565b6040518082600260200280838360005b83811015610144578082015181840152602081019050610129565b5050505090500191505060405180910390f35b61015f610346565b6040518082815260200191505060405180910390f35b61017d61034b565b6040518082815260200191505060405180910390f35b61019b610350565b604051808381526020018281526020019250505060405180910390f35b6101c0610362565b6040518082815260200191505060405180910390f35b6101de610367565b6040518082600260200280838360005b838110156102095780820151818401526020810190506101ee565b5050505090500191505060405180910390f35b6102246103b6565b604051808381526020018281526020019250505060405180910390f35b6102496103c8565b604051808381526020018281526020019250505060405180910390f35b61026e6103da565b6040518082815260200191505060405180910390f35b61028c6103e3565b604051808381526020018281526020019250505060405180910390f35b6102b16103f5565b6102b96103f5565b6000800154816000600281106102cb57fe5b602002018181525050600060010154816001600281106102e757fe5b6020020181815250508091505090565b6102ff6103f5565b6103076103f5565b6006600001548160006002811061031a57fe5b6020020181815250506006600101548160016002811061033657fe5b6020020181815250508091505090565b600581565b602081565b60048060000154908060010154905082565b600281565b61036f6103f5565b6103776103f5565b6002600001548160006002811061038a57fe5b602002018181525050600260010154816001600281106103a657fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b60006020905090565b60008060000154908060010154905082565b6040518060400160405280600290602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161044657fe5b06915050919050565b61045761058c565b6001821415610468578290506104ee565b60028214156104825761047b83846104f4565b90506104ee565b61048a6105a6565b83600001518160006003811061049c57fe5b6020020181815250508360200151816001600381106104b757fe5b60200201818152505082816002600381106104ce57fe5b6020020181815250506040826060836007600019fa6104ec57600080fd5b505b92915050565b6104fc61058c565b6105046105c8565b83600001518160006004811061051657fe5b60200201818152505083602001518160016004811061053157fe5b60200201818152505082600001518160026004811061054c57fe5b60200201818152505082602001518160036004811061056757fe5b6020020181815250506040826080836006600019fa61058557600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b604051806080016040528060049060208202803883398082019150509050509056fea265627a7a72315820eaa98ca9a5be5d508c8f1a1a8a12a02cb9b3cee9fdb7351fdaf69f04e6537aaa64736f6c63430005100032"

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

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) AggSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "aggSize")
	return *ret0, err
}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() constant returns(uint256)
func (_Publicparams *PublicparamsSession) AggSize() (*big.Int, error) {
	return _Publicparams.Contract.AggSize(&_Publicparams.CallOpts)
}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) AggSize() (*big.Int, error) {
	return _Publicparams.Contract.AggSize(&_Publicparams.CallOpts)
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
