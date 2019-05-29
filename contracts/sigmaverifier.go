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

// SigmaverifierABI is the input ABI used to generate the binding from.
const SigmaverifierABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"params\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[20]\"},{\"name\":\"z1\",\"type\":\"uint256\"},{\"name\":\"z2\",\"type\":\"uint256\"},{\"name\":\"z3\",\"type\":\"uint256\"}],\"name\":\"verifySigmaProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SigmaverifierBin is the compiled bytecode used for deploying new contracts.
const SigmaverifierBin = `0x60806040523480156200001157600080fd5b5060405160208062000ff88339810180604052620000339190810190620002d9565b6000819050620000426200020b565b8173ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b158015620000a457600080fd5b505afa158015620000b9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250620000df919081019062000305565b9050620000eb6200020b565b8273ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200014d57600080fd5b505afa15801562000162573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525062000188919081019062000305565b90508160006002811015156200019a57fe5b60200201516000800181905550816001600281101515620001b757fe5b6020020151600060010181905550806000600281101515620001d557fe5b6020020151600260000181905550806001600281101515620001f357fe5b602002015160026001018190555050505050620003c0565b6040805190810160405280600290602082028038833980820191505090505090565b60006200023b8251620003a2565b905092915050565b600082601f83011215156200025757600080fd5b60026200026e62000268826200035f565b62000331565b915081838560208402820111156200028557600080fd5b60005b83811015620002b957816200029e8882620002c3565b84526020840193506020830192505060018101905062000288565b5050505092915050565b6000620002d18251620003b6565b905092915050565b600060208284031215620002ec57600080fd5b6000620002fc848285016200022d565b91505092915050565b6000604082840312156200031857600080fd5b6000620003288482850162000243565b91505092915050565b6000604051905081810181811067ffffffffffffffff821117156200035557600080fd5b8060405250919050565b600067ffffffffffffffff8211156200037757600080fd5b602082029050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003af8262000382565b9050919050565b6000819050919050565b610c2880620003d06000396000f3fe608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680636e26cc311461005c578063b8c9d36514610099578063e2179b8e146100c5575b600080fd5b34801561006857600080fd5b50610083600480360361007e9190810190610ab6565b6100f1565b6040516100909190610b3b565b60405180910390f35b3480156100a557600080fd5b506100ae61057c565b6040516100bc929190610b56565b60405180910390f35b3480156100d157600080fd5b506100da61058e565b6040516100e8929190610b56565b60405180910390f35b60008061019c86600c60148110151561010657fe5b602002015187600d60148110151561011a57fe5b602002015188600e60148110151561012e57fe5b602002015189600f60148110151561014257fe5b60200201518a601060148110151561015657fe5b60200201518b601160148110151561016a57fe5b60200201518c601260148110151561017e57fe5b60200201518d601360148110151561019257fe5b60200201516105a0565b90506101a6610956565b60408051908101604052808860006014811015156101c057fe5b602002015181526020018860016014811015156101d957fe5b60200201518152508160006003811015156101f057fe5b6020020181905250604080519081016040528088600c60148110151561021257fe5b6020020151815260200188600d60148110151561022b57fe5b602002015181525081600160038110151561024257fe5b6020020181905250604080519081016040528088600260148110151561026457fe5b6020020151815260200188600360148110151561027d57fe5b602002015181525081600260038110151561029457fe5b60200201819052506102a781878461061b565b15156102b857600092505050610574565b6102c0610956565b60408051908101604052808960066014811015156102da57fe5b602002015181526020018960076014811015156102f357fe5b602002015181525081600060038110151561030a57fe5b6020020181905250604080519081016040528089600e60148110151561032c57fe5b6020020151815260200189600f60148110151561034557fe5b602002015181525081600160038110151561035c57fe5b6020020181905250604080519081016040528089600860148110151561037e57fe5b6020020151815260200189600960148110151561039757fe5b60200201518152508160026003811015156103ae57fe5b60200201819052506103c181878561061b565b15156103d35760009350505050610574565b6103db610984565b60408051908101604052808a60106014811015156103f557fe5b602002015181526020018a601160148110151561040e57fe5b602002015181525081600060028110151561042557fe5b602002018190525060408051908101604052808a600460148110151561044757fe5b602002015181526020018a600560148110151561046057fe5b602002015181525081600160028110151561047757fe5b602002018190525061048b818988876106cb565b151561049e576000945050505050610574565b6104a6610984565b60408051908101604052808b60126014811015156104c057fe5b602002015181526020018b60136014811015156104d957fe5b60200201518152508160006002811015156104f057fe5b602002018190525060408051908101604052808b600a60148110151561051257fe5b602002015181526020018b600b60148110151561052b57fe5b602002015181525081600160028110151561054257fe5b6020020181905250610556818989886106cb565b151561056a57600095505050505050610574565b6001955050505050505b949350505050565b60028060000154908060010154905082565b60008060000154908060010154905082565b600061060d8989898989898989604051602001808981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020019850505050505050505060405160208183030381529060405280519060200120600190046107cd565b905098975050505050505050565b60006106256109b2565b61067085600160038110151561063757fe5b60200201516106628588600260038110151561064f57fe5b602002015161080790919063ffffffff16565b6108b490919063ffffffff16565b905061067a6109b2565b6106a08587600060038110151561068d57fe5b602002015161080790919063ffffffff16565b9050806000015182600001511480156106c0575080602001518260200151145b925050509392505050565b60006106d56109b2565b61074c61070b8560006040805190810160405290816000820154815260200160018201548152505061080790919063ffffffff16565b61073e8760026040805190810160405290816000820154815260200160018201548152505061080790919063ffffffff16565b6108b490919063ffffffff16565b90506107566109b2565b6107a161077f8589600160028110151561076c57fe5b602002015161080790919063ffffffff16565b88600060028110151561078e57fe5b60200201516108b490919063ffffffff16565b9050806000015182600001511480156107c1575080602001518260200151145b92505050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156107fe57fe5b06915050919050565b61080f6109b2565b6001821415610820578290506108ae565b600282141561083a5761083383846108b4565b90506108ae565b6108426109cc565b836000015181600060038110151561085657fe5b602002018181525050836020015181600160038110151561087357fe5b6020020181815250508281600260038110151561088c57fe5b6020020181815250506040826060836007600019fa15156108ac57600080fd5b505b92915050565b6108bc6109b2565b6108c46109ef565b83600001518160006004811015156108d857fe5b60200201818152505083602001518160016004811015156108f557fe5b602002018181525050826000015181600260048110151561091257fe5b602002018181525050826020015181600360048110151561092f57fe5b6020020181815250506040826080836006600019fa151561094f57600080fd5b5092915050565b60c0604051908101604052806003905b61096e610a12565b8152602001906001900390816109665790505090565b6080604051908101604052806002905b61099c610a12565b8152602001906001900390816109945790505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b600082601f8301121515610a3f57600080fd5b6014610a52610a4d82610bac565b610b7f565b91508183856020840282011115610a6857600080fd5b60005b83811015610a985781610a7e8882610aa2565b845260208401935060208301925050600181019050610a6b565b5050505092915050565b6000610aae8235610be4565b905092915050565b6000806000806102e08587031215610acd57600080fd5b6000610adb87828801610a2c565b945050610280610aed87828801610aa2565b9350506102a0610aff87828801610aa2565b9250506102c0610b1187828801610aa2565b91505092959194509250565b610b2681610bce565b82525050565b610b3581610bda565b82525050565b6000602082019050610b506000830184610b1d565b92915050565b6000604082019050610b6b6000830185610b2c565b610b786020830184610b2c565b9392505050565b6000604051905081810181811067ffffffffffffffff82111715610ba257600080fd5b8060405250919050565b600067ffffffffffffffff821115610bc357600080fd5b602082029050919050565b60008115159050919050565b6000819050919050565b600081905091905056fea265627a7a723058205129e5fbac0eac44709663e9ab142c24b1380cd79d11b6295a923886cd3dbf226c6578706572696d656e74616cf50037`

// DeploySigmaverifier deploys a new Ethereum contract, binding an instance of Sigmaverifier to it.
func DeploySigmaverifier(auth *bind.TransactOpts, backend bind.ContractBackend, params common.Address) (common.Address, *types.Transaction, *Sigmaverifier, error) {
	parsed, err := abi.JSON(strings.NewReader(SigmaverifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigmaverifierBin), backend, params)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sigmaverifier{SigmaverifierCaller: SigmaverifierCaller{contract: contract}, SigmaverifierTransactor: SigmaverifierTransactor{contract: contract}, SigmaverifierFilterer: SigmaverifierFilterer{contract: contract}}, nil
}

// Sigmaverifier is an auto generated Go binding around an Ethereum contract.
type Sigmaverifier struct {
	SigmaverifierCaller     // Read-only binding to the contract
	SigmaverifierTransactor // Write-only binding to the contract
	SigmaverifierFilterer   // Log filterer for contract events
}

// SigmaverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigmaverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigmaverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigmaverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigmaverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigmaverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigmaverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigmaverifierSession struct {
	Contract     *Sigmaverifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigmaverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigmaverifierCallerSession struct {
	Contract *SigmaverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SigmaverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigmaverifierTransactorSession struct {
	Contract     *SigmaverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SigmaverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigmaverifierRaw struct {
	Contract *Sigmaverifier // Generic contract binding to access the raw methods on
}

// SigmaverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigmaverifierCallerRaw struct {
	Contract *SigmaverifierCaller // Generic read-only contract binding to access the raw methods on
}

// SigmaverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigmaverifierTransactorRaw struct {
	Contract *SigmaverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigmaverifier creates a new instance of Sigmaverifier, bound to a specific deployed contract.
func NewSigmaverifier(address common.Address, backend bind.ContractBackend) (*Sigmaverifier, error) {
	contract, err := bindSigmaverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sigmaverifier{SigmaverifierCaller: SigmaverifierCaller{contract: contract}, SigmaverifierTransactor: SigmaverifierTransactor{contract: contract}, SigmaverifierFilterer: SigmaverifierFilterer{contract: contract}}, nil
}

// NewSigmaverifierCaller creates a new read-only instance of Sigmaverifier, bound to a specific deployed contract.
func NewSigmaverifierCaller(address common.Address, caller bind.ContractCaller) (*SigmaverifierCaller, error) {
	contract, err := bindSigmaverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigmaverifierCaller{contract: contract}, nil
}

// NewSigmaverifierTransactor creates a new write-only instance of Sigmaverifier, bound to a specific deployed contract.
func NewSigmaverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*SigmaverifierTransactor, error) {
	contract, err := bindSigmaverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigmaverifierTransactor{contract: contract}, nil
}

// NewSigmaverifierFilterer creates a new log filterer instance of Sigmaverifier, bound to a specific deployed contract.
func NewSigmaverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*SigmaverifierFilterer, error) {
	contract, err := bindSigmaverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigmaverifierFilterer{contract: contract}, nil
}

// bindSigmaverifier binds a generic wrapper to an already deployed contract.
func bindSigmaverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigmaverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sigmaverifier *SigmaverifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sigmaverifier.Contract.SigmaverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sigmaverifier *SigmaverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.SigmaverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sigmaverifier *SigmaverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.SigmaverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sigmaverifier *SigmaverifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sigmaverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sigmaverifier *SigmaverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sigmaverifier *SigmaverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.contract.Transact(opts, method, params...)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierCaller) G(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Sigmaverifier.contract.Call(opts, out, "g")
	return *ret, err
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sigmaverifier.Contract.G(&_Sigmaverifier.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierCallerSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sigmaverifier.Contract.G(&_Sigmaverifier.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierCaller) H(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Sigmaverifier.contract.Call(opts, out, "h")
	return *ret, err
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sigmaverifier.Contract.H(&_Sigmaverifier.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Sigmaverifier *SigmaverifierCallerSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Sigmaverifier.Contract.H(&_Sigmaverifier.CallOpts)
}

// VerifySigmaProof is a free data retrieval call binding the contract method 0x6e26cc31.
//
// Solidity: function verifySigmaProof(uint256[20] points, uint256 z1, uint256 z2, uint256 z3) constant returns(bool)
func (_Sigmaverifier *SigmaverifierCaller) VerifySigmaProof(opts *bind.CallOpts, points [20]*big.Int, z1 *big.Int, z2 *big.Int, z3 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Sigmaverifier.contract.Call(opts, out, "verifySigmaProof", points, z1, z2, z3)
	return *ret0, err
}

// VerifySigmaProof is a free data retrieval call binding the contract method 0x6e26cc31.
//
// Solidity: function verifySigmaProof(uint256[20] points, uint256 z1, uint256 z2, uint256 z3) constant returns(bool)
func (_Sigmaverifier *SigmaverifierSession) VerifySigmaProof(points [20]*big.Int, z1 *big.Int, z2 *big.Int, z3 *big.Int) (bool, error) {
	return _Sigmaverifier.Contract.VerifySigmaProof(&_Sigmaverifier.CallOpts, points, z1, z2, z3)
}

// VerifySigmaProof is a free data retrieval call binding the contract method 0x6e26cc31.
//
// Solidity: function verifySigmaProof(uint256[20] points, uint256 z1, uint256 z2, uint256 z3) constant returns(bool)
func (_Sigmaverifier *SigmaverifierCallerSession) VerifySigmaProof(points [20]*big.Int, z1 *big.Int, z2 *big.Int, z3 *big.Int) (bool, error) {
	return _Sigmaverifier.Contract.VerifySigmaProof(&_Sigmaverifier.CallOpts, points, z1, z2, z3)
}
