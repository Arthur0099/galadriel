// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package publicparams

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PublicparamsMetaData contains all meta data concerning the Publicparams contract.
var PublicparamsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pkx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pky\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"aggSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pkauth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"u\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getG\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getH\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getU\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPK\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000bc538038062000bc5833981810160405260408110156200003757600080fd5b810190808051906020019092919080519060200190929190505050600160046000018190555060026004600101819055508160086000018190555080600860010181905550620000cb6040518060400160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c67000000000000815250620001bb640100000000026401000000009004565b600080820151816000015560208201518160010155905050620001326040518060400160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c67000000000000815250620001bb640100000000026401000000009004565b600260008201518160000155602082015181600101559050506200019a6040518060400160405280601b81526020017f752067656e657261746f72206f6620696e6e657270726f647563740000000000815250620001bb640100000000026401000000009004565b600660008201518160000155602082015181600101559050505050620004a4565b620001c562000446565b60006200025d836040516020018082805190602001908083835b60208310620002045780518252602082019150602081019050602083039250620001df565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060019004620002aa6401000000000262000509176401000000009004565b9050620002a281600460405180604001604052908160008201548152602001600182015481525050620002e36401000000000262000541179091906401000000009004565b915050919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381620002da57fe5b06915050919050565b620002ed62000446565b600182141562000300578290506200039f565b60028214156200032d57620003258384620003a5640100000000026401000000009004565b90506200039f565b6200033762000460565b8360000151816000600381106200034a57fe5b6020020181815250508360200151816001600381106200036657fe5b60200201818152505082816002600381106200037e57fe5b6020020181815250506040826060836007600019fa6200039d57600080fd5b505b92915050565b620003af62000446565b620003b962000482565b836000015181600060048110620003cc57fe5b602002018181525050836020015181600160048110620003e857fe5b6020020181815250508260000151816002600481106200040457fe5b6020020181815250508260200151816003600481106200042057fe5b6020020181815250506040826080836006600019fa6200043f57600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b61071180620004b46000396000f3fe608060405234801561001057600080fd5b50600436106100ec576000357c0100000000000000000000000000000000000000000000000000000000900480637982ebcc116100a9578063b8c9d36511610083578063b8c9d365146102ad578063c6a898c5146102d2578063da897224146102f7578063e2179b8e14610315576100ec565b80637982ebcc1461020357806382529fdb14610221578063ac9e14f514610267576100ec565b80630214120b146100f157806304c09ce91461011657806324d6147d1461015c5780632e52d606146101a25780633e8d3764146101c057806352c9b965146101de575b600080fd5b6100f961033a565b604051808381526020018281526020019250505060405180910390f35b61011e61034c565b6040518082600260200280838360005b8381101561014957808201518184015260208101905061012e565b5050505090500191505060405180910390f35b61016461039a565b6040518082600260200280838360005b8381101561018f578082015181840152602081019050610174565b5050505090500191505060405180910390f35b6101aa6103e9565b6040518082815260200191505060405180910390f35b6101c86103ee565b6040518082815260200191505060405180910390f35b6101e66103f3565b604051808381526020018281526020019250505060405180910390f35b61020b610405565b6040518082815260200191505060405180910390f35b61022961040a565b6040518082600260200280838360005b83811015610254578082015181840152602081019050610239565b5050505090500191505060405180910390f35b61026f610459565b6040518082600260200280838360005b8381101561029a57808201518184015260208101905061027f565b5050505090500191505060405180910390f35b6102b56104a8565b604051808381526020018281526020019250505060405180910390f35b6102da6104ba565b604051808381526020018281526020019250505060405180910390f35b6102ff6104cc565b6040518082815260200191505060405180910390f35b61031d6104d5565b604051808381526020018281526020019250505060405180910390f35b60088060000154908060010154905082565b6103546104e7565b61035c6104e7565b60008001548160006002811061036e57fe5b6020020181815250506000600101548160016002811061038a57fe5b6020020181815250508091505090565b6103a26104e7565b6103aa6104e7565b600660000154816000600281106103bd57fe5b602002018181525050600660010154816001600281106103d957fe5b6020020181815250508091505090565b600581565b602081565b60048060000154908060010154905082565b600281565b6104126104e7565b61041a6104e7565b6002600001548160006002811061042d57fe5b6020020181815250506002600101548160016002811061044957fe5b6020020181815250508091505090565b6104616104e7565b6104696104e7565b6008600001548160006002811061047c57fe5b6020020181815250506008600101548160016002811061049857fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b60006020905090565b60008060000154908060010154905082565b6040518060400160405280600290602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161053857fe5b06915050919050565b61054961067e565b600182141561055a578290506105e0565b60028214156105745761056d83846105e6565b90506105e0565b61057c610698565b83600001518160006003811061058e57fe5b6020020181815250508360200151816001600381106105a957fe5b60200201818152505082816002600381106105c057fe5b6020020181815250506040826060836007600019fa6105de57600080fd5b505b92915050565b6105ee61067e565b6105f66106ba565b83600001518160006004811061060857fe5b60200201818152505083602001518160016004811061062357fe5b60200201818152505082600001518160026004811061063e57fe5b60200201818152505082602001518160036004811061065957fe5b6020020181815250506040826080836006600019fa61067757600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b604051806080016040528060049060208202803883398082019150509050509056fea265627a7a72315820d2ac6c724e6e5a319ae5cdfa2037a333ae4d2dca4c4accbee9b17929a2a80f6864736f6c63430005100032",
}

// PublicparamsABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicparamsMetaData.ABI instead.
var PublicparamsABI = PublicparamsMetaData.ABI

// PublicparamsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicparamsMetaData.Bin instead.
var PublicparamsBin = PublicparamsMetaData.Bin

// DeployPublicparams deploys a new Ethereum contract, binding an instance of Publicparams to it.
func DeployPublicparams(auth *bind.TransactOpts, backend bind.ContractBackend, pkx *big.Int, pky *big.Int) (common.Address, *types.Transaction, *Publicparams, error) {
	parsed, err := PublicparamsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicparamsBin), backend, pkx, pky)
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
func (_Publicparams *PublicparamsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Publicparams *PublicparamsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function aggSize() view returns(uint256)
func (_Publicparams *PublicparamsCaller) AggSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "aggSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() view returns(uint256)
func (_Publicparams *PublicparamsSession) AggSize() (*big.Int, error) {
	return _Publicparams.Contract.AggSize(&_Publicparams.CallOpts)
}

// AggSize is a free data retrieval call binding the contract method 0x7982ebcc.
//
// Solidity: function aggSize() view returns(uint256)
func (_Publicparams *PublicparamsCallerSession) AggSize() (*big.Int, error) {
	return _Publicparams.Contract.AggSize(&_Publicparams.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() view returns(uint256)
func (_Publicparams *PublicparamsCaller) BitSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "bitSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() view returns(uint256)
func (_Publicparams *PublicparamsSession) BitSize() (*big.Int, error) {
	return _Publicparams.Contract.BitSize(&_Publicparams.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() view returns(uint256)
func (_Publicparams *PublicparamsCallerSession) BitSize() (*big.Int, error) {
	return _Publicparams.Contract.BitSize(&_Publicparams.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) G(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "g")

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.G(&_Publicparams.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.G(&_Publicparams.CallOpts)
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) GBase(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "gBase")

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) GBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GBase(&_Publicparams.CallOpts)
}

// GBase is a free data retrieval call binding the contract method 0x52c9b965.
//
// Solidity: function gBase() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) GBase() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GBase(&_Publicparams.CallOpts)
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() pure returns(uint256)
func (_Publicparams *PublicparamsCaller) GetBitSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "getBitSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() pure returns(uint256)
func (_Publicparams *PublicparamsSession) GetBitSize() (*big.Int, error) {
	return _Publicparams.Contract.GetBitSize(&_Publicparams.CallOpts)
}

// GetBitSize is a free data retrieval call binding the contract method 0xda897224.
//
// Solidity: function getBitSize() pure returns(uint256)
func (_Publicparams *PublicparamsCallerSession) GetBitSize() (*big.Int, error) {
	return _Publicparams.Contract.GetBitSize(&_Publicparams.CallOpts)
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() view returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetG(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "getG")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() view returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetG() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetG(&_Publicparams.CallOpts)
}

// GetG is a free data retrieval call binding the contract method 0x04c09ce9.
//
// Solidity: function getG() view returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetG() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetG(&_Publicparams.CallOpts)
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() view returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetH(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "getH")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() view returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetH() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetH(&_Publicparams.CallOpts)
}

// GetH is a free data retrieval call binding the contract method 0x82529fdb.
//
// Solidity: function getH() view returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetH() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetH(&_Publicparams.CallOpts)
}

// GetPK is a free data retrieval call binding the contract method 0xac9e14f5.
//
// Solidity: function getPK() view returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetPK(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "getPK")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetPK is a free data retrieval call binding the contract method 0xac9e14f5.
//
// Solidity: function getPK() view returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetPK() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetPK(&_Publicparams.CallOpts)
}

// GetPK is a free data retrieval call binding the contract method 0xac9e14f5.
//
// Solidity: function getPK() view returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetPK() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetPK(&_Publicparams.CallOpts)
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() view returns(uint256[2])
func (_Publicparams *PublicparamsCaller) GetU(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "getU")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() view returns(uint256[2])
func (_Publicparams *PublicparamsSession) GetU() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetU(&_Publicparams.CallOpts)
}

// GetU is a free data retrieval call binding the contract method 0x24d6147d.
//
// Solidity: function getU() view returns(uint256[2])
func (_Publicparams *PublicparamsCallerSession) GetU() ([2]*big.Int, error) {
	return _Publicparams.Contract.GetU(&_Publicparams.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) H(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "h")

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.H(&_Publicparams.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.H(&_Publicparams.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Publicparams *PublicparamsCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "n")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Publicparams *PublicparamsSession) N() (*big.Int, error) {
	return _Publicparams.Contract.N(&_Publicparams.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Publicparams *PublicparamsCallerSession) N() (*big.Int, error) {
	return _Publicparams.Contract.N(&_Publicparams.CallOpts)
}

// Pkauth is a free data retrieval call binding the contract method 0x0214120b.
//
// Solidity: function pkauth() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) Pkauth(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "pkauth")

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pkauth is a free data retrieval call binding the contract method 0x0214120b.
//
// Solidity: function pkauth() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) Pkauth() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.Pkauth(&_Publicparams.CallOpts)
}

// Pkauth is a free data retrieval call binding the contract method 0x0214120b.
//
// Solidity: function pkauth() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) Pkauth() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.Pkauth(&_Publicparams.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) U(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _Publicparams.contract.Call(opts, &out, "u")

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.U(&_Publicparams.CallOpts)
}

// U is a free data retrieval call binding the contract method 0xc6a898c5.
//
// Solidity: function u() view returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) U() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.U(&_Publicparams.CallOpts)
}
