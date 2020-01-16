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
const SigmaverifierABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"params\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[16]\"},{\"name\":\"z1\",\"type\":\"uint256\"},{\"name\":\"z2\",\"type\":\"uint256\"}],\"name\":\"verifyPTEProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[10]\"},{\"name\":\"z1\",\"type\":\"uint256\"},{\"name\":\"z2\",\"type\":\"uint256\"}],\"name\":\"verifyCTValidProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[20]\"},{\"name\":\"z1\",\"type\":\"uint256\"},{\"name\":\"z2\",\"type\":\"uint256\"},{\"name\":\"z3\",\"type\":\"uint256\"}],\"name\":\"verifySigmaProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SigmaverifierBin is the compiled bytecode used for deploying new contracts.
const SigmaverifierBin = `0x60806040523480156200001157600080fd5b5060405160208062001470833981018060405260208110156200003357600080fd5b81019080805190602001909291905050506000819050620000536200022c565b8173ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b158015620000b557600080fd5b505afa158015620000ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620000f057600080fd5b81019080919050509050620001046200022c565b8273ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200016657600080fd5b505afa1580156200017b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620001a157600080fd5b81019080919050509050816000600281101515620001bb57fe5b60200201516000800181905550816001600281101515620001d857fe5b6020020151600060010181905550806000600281101515620001f657fe5b60200201516002600001819055508060016002811015156200021457fe5b6020020151600260010181905550505050506200024e565b6040805190810160405280600290602082028038833980820191505090505090565b611212806200025e6000396000f3fe60806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680636e26cc3114610072578063991395ca1461011c578063b8c9d365146101bc578063e1efe766146101ee578063e2179b8e1461028e575b600080fd5b34801561007e57600080fd5b5061010260048036036102e081101561009657600080fd5b810190808061028001906014806020026040519081016040528092919082601460200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080359060200190929190803590602001909291905050506102c0565b604051808215151515815260200191505060405180910390f35b34801561012857600080fd5b506101a2600480360361018081101561014057600080fd5b81019080806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291908035906020019092919050505061074b565b604051808215151515815260200191505060405180910390f35b3480156101c857600080fd5b506101d1610987565b604051808381526020018281526020019250505060405180910390f35b3480156101fa57600080fd5b50610274600480360361024081101561021257600080fd5b810190808061020001906010806020026040519081016040528092919082601060200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080359060200190929190505050610999565b604051808215151515815260200191505060405180910390f35b34801561029a57600080fd5b506102a3610cf1565b604051808381526020018281526020019250505060405180910390f35b60008061036b86600c6014811015156102d557fe5b602002015187600d6014811015156102e957fe5b602002015188600e6014811015156102fd57fe5b602002015189600f60148110151561031157fe5b60200201518a601060148110151561032557fe5b60200201518b601160148110151561033957fe5b60200201518c601260148110151561034d57fe5b60200201518d601360148110151561036157fe5b6020020151610d03565b9050610375611110565b604080519081016040528088600060148110151561038f57fe5b602002015181526020018860016014811015156103a857fe5b60200201518152508160006003811015156103bf57fe5b6020020181905250604080519081016040528088600c6014811015156103e157fe5b6020020151815260200188600d6014811015156103fa57fe5b602002015181525081600160038110151561041157fe5b6020020181905250604080519081016040528088600260148110151561043357fe5b6020020151815260200188600360148110151561044c57fe5b602002015181525081600260038110151561046357fe5b6020020181905250610476818784610d7e565b151561048757600092505050610743565b61048f611110565b60408051908101604052808960066014811015156104a957fe5b602002015181526020018960076014811015156104c257fe5b60200201518152508160006003811015156104d957fe5b6020020181905250604080519081016040528089600e6014811015156104fb57fe5b6020020151815260200189600f60148110151561051457fe5b602002015181525081600160038110151561052b57fe5b6020020181905250604080519081016040528089600860148110151561054d57fe5b6020020151815260200189600960148110151561056657fe5b602002015181525081600260038110151561057d57fe5b6020020181905250610590818785610d7e565b15156105a25760009350505050610743565b6105aa61113e565b60408051908101604052808a60106014811015156105c457fe5b602002015181526020018a60116014811015156105dd57fe5b60200201518152508160006002811015156105f457fe5b602002018190525060408051908101604052808a600460148110151561061657fe5b602002015181526020018a600560148110151561062f57fe5b602002015181525081600160028110151561064657fe5b602002018190525061065a81898887610e2e565b151561066d576000945050505050610743565b61067561113e565b60408051908101604052808b601260148110151561068f57fe5b602002015181526020018b60136014811015156106a857fe5b60200201518152508160006002811015156106bf57fe5b602002018190525060408051908101604052808b600a6014811015156106e157fe5b602002015181526020018b600b6014811015156106fa57fe5b602002015181525081600160028110151561071157fe5b602002018190525061072581898988610e2e565b151561073957600095505050505050610743565b6001955050505050505b949350505050565b6000806107a6856006600a8110151561076057fe5b6020020151866007600a8110151561077457fe5b6020020151876008600a8110151561078857fe5b6020020151886009600a8110151561079c57fe5b6020020151610f30565b90506107b0611110565b6040805190810160405280876000600a811015156107ca57fe5b60200201518152602001876001600a811015156107e357fe5b60200201518152508160006003811015156107fa57fe5b60200201819052506040805190810160405280876006600a8110151561081c57fe5b60200201518152602001876007600a8110151561083557fe5b602002015181525081600160038110151561084c57fe5b60200201819052506040805190810160405280876002600a8110151561086e57fe5b60200201518152602001876003600a8110151561088757fe5b602002015181525081600260038110151561089e57fe5b60200201819052506108b1818684610d7e565b15156108c257600092505050610980565b6108ca61113e565b6040805190810160405280886008600a811015156108e457fe5b60200201518152602001886009600a811015156108fd57fe5b602002015181525081600060028110151561091457fe5b60200201819052506040805190810160405280886004600a8110151561093657fe5b60200201518152602001886005600a8110151561094f57fe5b602002015181525081600160028110151561096657fe5b602002018190525061097a81878786610e2e565b93505050505b9392505050565b60028060000154908060010154905082565b6000806109f485600a6010811015156109ae57fe5b602002015186600b6010811015156109c257fe5b602002015187600c6010811015156109d657fe5b602002015188600d6010811015156109ea57fe5b6020020151610f30565b90506109fe611110565b6040805190810160405280876000601081101515610a1857fe5b60200201518152602001876001601081101515610a3157fe5b6020020151815250816000600381101515610a4857fe5b6020020181905250604080519081016040528087600a601081101515610a6a57fe5b6020020151815260200187600b601081101515610a8357fe5b6020020151815250816001600381101515610a9a57fe5b60200201819052506040805190810160405280876004601081101515610abc57fe5b60200201518152602001876005601081101515610ad557fe5b6020020151815250816002600381101515610aec57fe5b6020020181905250610aff818684610d7e565b1515610b1057600092505050610cea565b610b18611110565b6040805190810160405280886002601081101515610b3257fe5b60200201518152602001886003601081101515610b4b57fe5b6020020151815250816000600381101515610b6257fe5b6020020181905250604080519081016040528088600c601081101515610b8457fe5b6020020151815260200188600d601081101515610b9d57fe5b6020020151815250816001600381101515610bb457fe5b60200201819052506040805190810160405280886006601081101515610bd657fe5b60200201518152602001886007601081101515610bef57fe5b6020020151815250816002600381101515610c0657fe5b6020020181905250610c19818785610d7e565b1515610c2b5760009350505050610cea565b610c3361113e565b604080519081016040528089600e601081101515610c4d57fe5b6020020151815260200189600f601081101515610c6657fe5b6020020151815250816000600281101515610c7d57fe5b60200201819052506040805190810160405280896008601081101515610c9f57fe5b60200201518152602001896009601081101515610cb857fe5b6020020151815250816001600281101515610ccf57fe5b6020020181905250610ce381888887610e2e565b9450505050505b9392505050565b60008060000154908060010154905082565b6000610d70898989898989898960405160200180898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001985050505050505050506040516020818303038152906040528051906020012060019004610f87565b905098975050505050505050565b6000610d8861116c565b610dd3856001600381101515610d9a57fe5b6020020151610dc585886002600381101515610db257fe5b6020020151610fc190919063ffffffff16565b61106e90919063ffffffff16565b9050610ddd61116c565b610e0385876000600381101515610df057fe5b6020020151610fc190919063ffffffff16565b905080600001518260000151148015610e23575080602001518260200151145b925050509392505050565b6000610e3861116c565b610eaf610e6e85600060408051908101604052908160008201548152602001600182015481525050610fc190919063ffffffff16565b610ea187600260408051908101604052908160008201548152602001600182015481525050610fc190919063ffffffff16565b61106e90919063ffffffff16565b9050610eb961116c565b610f04610ee285896001600281101515610ecf57fe5b6020020151610fc190919063ffffffff16565b886000600281101515610ef157fe5b602002015161106e90919063ffffffff16565b905080600001518260000151148015610f24575080602001518260200151145b92505050949350505050565b6000610f7d85858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060019004610f87565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083811515610fb857fe5b06915050919050565b610fc961116c565b6001821415610fda57829050611068565b6002821415610ff457610fed838461106e565b9050611068565b610ffc611186565b836000015181600060038110151561101057fe5b602002018181525050836020015181600160038110151561102d57fe5b6020020181815250508281600260038110151561104657fe5b6020020181815250506040826060836007600019fa151561106657600080fd5b505b92915050565b61107661116c565b61107e6111a9565b836000015181600060048110151561109257fe5b60200201818152505083602001518160016004811015156110af57fe5b60200201818152505082600001518160026004811015156110cc57fe5b60200201818152505082602001518160036004811015156110e957fe5b6020020181815250506040826080836006600019fa151561110957600080fd5b5092915050565b60c0604051908101604052806003905b6111286111cc565b8152602001906001900390816111205790505090565b6080604051908101604052806002905b6111566111cc565b81526020019060019003908161114e5790505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b60408051908101604052806000815260200160008152509056fea165627a7a72305820f7d991b240b19da43e67f1c3e542ba834a6302c2a9ba06c2f9f344f524d6b3150029`

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

// VerifyCTValidProof is a paid mutator transaction binding the contract method 0x991395ca.
//
// Solidity: function verifyCTValidProof(uint256[10] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierTransactor) VerifyCTValidProof(opts *bind.TransactOpts, points [10]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.contract.Transact(opts, "verifyCTValidProof", points, z1, z2)
}

// VerifyCTValidProof is a paid mutator transaction binding the contract method 0x991395ca.
//
// Solidity: function verifyCTValidProof(uint256[10] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierSession) VerifyCTValidProof(points [10]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.VerifyCTValidProof(&_Sigmaverifier.TransactOpts, points, z1, z2)
}

// VerifyCTValidProof is a paid mutator transaction binding the contract method 0x991395ca.
//
// Solidity: function verifyCTValidProof(uint256[10] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierTransactorSession) VerifyCTValidProof(points [10]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.VerifyCTValidProof(&_Sigmaverifier.TransactOpts, points, z1, z2)
}

// VerifyPTEProof is a paid mutator transaction binding the contract method 0xe1efe766.
//
// Solidity: function verifyPTEProof(uint256[16] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierTransactor) VerifyPTEProof(opts *bind.TransactOpts, points [16]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.contract.Transact(opts, "verifyPTEProof", points, z1, z2)
}

// VerifyPTEProof is a paid mutator transaction binding the contract method 0xe1efe766.
//
// Solidity: function verifyPTEProof(uint256[16] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierSession) VerifyPTEProof(points [16]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.VerifyPTEProof(&_Sigmaverifier.TransactOpts, points, z1, z2)
}

// VerifyPTEProof is a paid mutator transaction binding the contract method 0xe1efe766.
//
// Solidity: function verifyPTEProof(uint256[16] points, uint256 z1, uint256 z2) returns(bool)
func (_Sigmaverifier *SigmaverifierTransactorSession) VerifyPTEProof(points [16]*big.Int, z1 *big.Int, z2 *big.Int) (*types.Transaction, error) {
	return _Sigmaverifier.Contract.VerifyPTEProof(&_Sigmaverifier.TransactOpts, points, z1, z2)
}
