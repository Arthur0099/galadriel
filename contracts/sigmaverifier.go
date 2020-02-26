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
const SigmaverifierBin = `0x60806040523480156200001157600080fd5b5060405160208062001480833981018060405260208110156200003357600080fd5b81019080805190602001909291905050506000819050620000536200022c565b8173ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b158015620000b557600080fd5b505afa158015620000ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620000f057600080fd5b81019080919050509050620001046200022c565b8273ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200016657600080fd5b505afa1580156200017b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620001a157600080fd5b81019080919050509050816000600281101515620001bb57fe5b60200201516000800181905550816001600281101515620001d857fe5b6020020151600060010181905550806000600281101515620001f657fe5b60200201516002600001819055508060016002811015156200021457fe5b6020020151600260010181905550505050506200024e565b6040805190810160405280600290602082028038833980820191505090505090565b611222806200025e6000396000f3fe60806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680636e26cc3114610072578063991395ca1461011c578063b8c9d365146101bc578063e1efe766146101ee578063e2179b8e1461028e575b600080fd5b34801561007e57600080fd5b5061010260048036036102e081101561009657600080fd5b810190808061028001906014806020026040519081016040528092919082601460200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080359060200190929190803590602001909291905050506102c0565b604051808215151515815260200191505060405180910390f35b34801561012857600080fd5b506101a2600480360361018081101561014057600080fd5b81019080806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080359060200190929190505050610735565b604051808215151515815260200191505060405180910390f35b3480156101c857600080fd5b506101d1610971565b604051808381526020018281526020019250505060405180910390f35b3480156101fa57600080fd5b50610274600480360361024081101561021257600080fd5b810190808061020001906010806020026040519081016040528092919082601060200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080359060200190929190505050610983565b604051808215151515815260200191505060405180910390f35b34801561029a57600080fd5b506102a3610d13565b604051808381526020018281526020019250505060405180910390f35b600080606060086040519080825280602002602001820160405280156102f55781602001602082028038833980820191505090505b50905060008090505b60088110156103465787600c820160148110151561031857fe5b6020020151828281518110151561032b57fe5b906020019060200201818152505080806001019150506102fe565b5061035081610d25565b915061035a611120565b604080519081016040528089600060148110151561037457fe5b6020020151815260200189600160148110151561038d57fe5b60200201518152508160006003811015156103a457fe5b6020020181905250604080519081016040528089600c6014811015156103c657fe5b6020020151815260200189600d6014811015156103df57fe5b60200201518152508160016003811015156103f657fe5b6020020181905250604080519081016040528089600260148110151561041857fe5b6020020151815260200189600360148110151561043157fe5b602002015181525081600260038110151561044857fe5b602002018190525061045b818885610d8e565b151561046d576000935050505061072d565b610475611120565b60408051908101604052808a600660148110151561048f57fe5b602002015181526020018a60076014811015156104a857fe5b60200201518152508160006003811015156104bf57fe5b602002018190525060408051908101604052808a600e6014811015156104e157fe5b602002015181526020018a600f6014811015156104fa57fe5b602002015181525081600160038110151561051157fe5b602002018190525060408051908101604052808a600860148110151561053357fe5b602002015181526020018a600960148110151561054c57fe5b602002015181525081600260038110151561056357fe5b6020020181905250610576818886610d8e565b151561058957600094505050505061072d565b61059161114e565b60408051908101604052808b60106014811015156105ab57fe5b602002015181526020018b60116014811015156105c457fe5b60200201518152508160006002811015156105db57fe5b602002018190525060408051908101604052808b60046014811015156105fd57fe5b602002015181526020018b600560148110151561061657fe5b602002015181525081600160028110151561062d57fe5b6020020181905250610641818a8988610e3e565b15156106555760009550505050505061072d565b61065d61114e565b60408051908101604052808c601260148110151561067757fe5b602002015181526020018c601360148110151561069057fe5b60200201518152508160006002811015156106a757fe5b602002018190525060408051908101604052808c600a6014811015156106c957fe5b602002015181526020018c600b6014811015156106e257fe5b60200201518152508160016002811015156106f957fe5b602002018190525061070d818a8a89610e3e565b1515610722576000965050505050505061072d565b600196505050505050505b949350505050565b600080610790856006600a8110151561074a57fe5b6020020151866007600a8110151561075e57fe5b6020020151876008600a8110151561077257fe5b6020020151886009600a8110151561078657fe5b6020020151610f40565b905061079a611120565b6040805190810160405280876000600a811015156107b457fe5b60200201518152602001876001600a811015156107cd57fe5b60200201518152508160006003811015156107e457fe5b60200201819052506040805190810160405280876006600a8110151561080657fe5b60200201518152602001876007600a8110151561081f57fe5b602002015181525081600160038110151561083657fe5b60200201819052506040805190810160405280876002600a8110151561085857fe5b60200201518152602001876003600a8110151561087157fe5b602002015181525081600260038110151561088857fe5b602002018190525061089b818684610d8e565b15156108ac5760009250505061096a565b6108b461114e565b6040805190810160405280886008600a811015156108ce57fe5b60200201518152602001886009600a811015156108e757fe5b60200201518152508160006002811015156108fe57fe5b60200201819052506040805190810160405280886004600a8110151561092057fe5b60200201518152602001886005600a8110151561093957fe5b602002015181525081600160028110151561095057fe5b602002018190525061096481878786610e3e565b93505050505b9392505050565b60028060000154908060010154905082565b600080606060066040519080825280602002602001820160405280156109b85781602001602082028038833980820191505090505b50905060008090505b6006811015610a09578681600a016010811015156109db57fe5b602002015182828151811015156109ee57fe5b906020019060200201818152505080806001019150506109c1565b50610a1381610d25565b9150610a1d611120565b6040805190810160405280886000601081101515610a3757fe5b60200201518152602001886001601081101515610a5057fe5b6020020151815250816000600381101515610a6757fe5b6020020181905250604080519081016040528088600a601081101515610a8957fe5b6020020151815260200188600b601081101515610aa257fe5b6020020151815250816001600381101515610ab957fe5b60200201819052506040805190810160405280886004601081101515610adb57fe5b60200201518152602001886005601081101515610af457fe5b6020020151815250816002600381101515610b0b57fe5b6020020181905250610b1e818785610d8e565b1515610b305760009350505050610d0c565b610b38611120565b6040805190810160405280896002601081101515610b5257fe5b60200201518152602001896003601081101515610b6b57fe5b6020020151815250816000600381101515610b8257fe5b6020020181905250604080519081016040528089600c601081101515610ba457fe5b6020020151815260200189600d601081101515610bbd57fe5b6020020151815250816001600381101515610bd457fe5b60200201819052506040805190810160405280896006601081101515610bf657fe5b60200201518152602001896007601081101515610c0f57fe5b6020020151815250816002600381101515610c2657fe5b6020020181905250610c39818886610d8e565b1515610c4c576000945050505050610d0c565b610c5461114e565b60408051908101604052808a600e601081101515610c6e57fe5b602002015181526020018a600f601081101515610c8757fe5b6020020151815250816000600281101515610c9e57fe5b602002018190525060408051908101604052808a6008601081101515610cc057fe5b602002015181526020018a6009601081101515610cd957fe5b6020020151815250816001600281101515610cf057fe5b6020020181905250610d0481898988610e3e565b955050505050505b9392505050565b60008060000154908060010154905082565b6000610d878260405160200180828051906020019060200280838360005b83811015610d5e578082015181840152602081019050610d43565b505050509050019150506040516020818303038152906040528051906020012060019004610f97565b9050919050565b6000610d9861117c565b610de3856001600381101515610daa57fe5b6020020151610dd585886002600381101515610dc257fe5b6020020151610fd190919063ffffffff16565b61107e90919063ffffffff16565b9050610ded61117c565b610e1385876000600381101515610e0057fe5b6020020151610fd190919063ffffffff16565b905080600001518260000151148015610e33575080602001518260200151145b925050509392505050565b6000610e4861117c565b610ebf610e7e85600060408051908101604052908160008201548152602001600182015481525050610fd190919063ffffffff16565b610eb187600260408051908101604052908160008201548152602001600182015481525050610fd190919063ffffffff16565b61107e90919063ffffffff16565b9050610ec961117c565b610f14610ef285896001600281101515610edf57fe5b6020020151610fd190919063ffffffff16565b886000600281101515610f0157fe5b602002015161107e90919063ffffffff16565b905080600001518260000151148015610f34575080602001518260200151145b92505050949350505050565b6000610f8d85858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060019004610f97565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083811515610fc857fe5b06915050919050565b610fd961117c565b6001821415610fea57829050611078565b600282141561100457610ffd838461107e565b9050611078565b61100c611196565b836000015181600060038110151561102057fe5b602002018181525050836020015181600160038110151561103d57fe5b6020020181815250508281600260038110151561105657fe5b6020020181815250506040826060836007600019fa151561107657600080fd5b505b92915050565b61108661117c565b61108e6111b9565b83600001518160006004811015156110a257fe5b60200201818152505083602001518160016004811015156110bf57fe5b60200201818152505082600001518160026004811015156110dc57fe5b60200201818152505082602001518160036004811015156110f957fe5b6020020181815250506040826080836006600019fa151561111957600080fd5b5092915050565b60c0604051908101604052806003905b6111386111dc565b8152602001906001900390816111305790505090565b6080604051908101604052806002905b6111666111dc565b81526020019060019003908161115e5790505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b60408051908101604052806000815260200160008152509056fea165627a7a72305820929004e4bde5b60969b826f88c87dc5b2784ece9d78152386bde1a649165fa6a0029`

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
