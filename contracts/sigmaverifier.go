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
const SigmaverifierBin = `0x608060405234801561001057600080fd5b50604051602080610d1c8339810180604052602081101561003057600080fd5b8101908080519060200190929190505050600081905061004e61021a565b8173ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156100af57600080fd5b505afa1580156100c3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156100e857600080fd5b810190809190505090506100fa61021a565b8273ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b15801561015b57600080fd5b505afa15801561016f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250604081101561019457600080fd5b810190809190505090508160006002811015156101ad57fe5b602002015160008001819055508160016002811015156101c957fe5b60200201516000600101819055508060006002811015156101e657fe5b602002015160026000018190555080600160028110151561020357fe5b60200201516002600101819055505050505061023c565b6040805190810160405280600290602082028038833980820191505090505090565b610ad18061024b6000396000f3fe608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680636e26cc311461005c578063b8c9d36514610106578063e2179b8e14610138575b600080fd5b34801561006857600080fd5b506100ec60048036036102e081101561008057600080fd5b810190808061028001906014806020026040519081016040528092919082601460200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190803590602001909291908035906020019092919050505061016a565b604051808215151515815260200191505060405180910390f35b34801561011257600080fd5b5061011b6105f5565b604051808381526020018281526020019250505060405180910390f35b34801561014457600080fd5b5061014d610607565b604051808381526020018281526020019250505060405180910390f35b60008061021586600c60148110151561017f57fe5b602002015187600d60148110151561019357fe5b602002015188600e6014811015156101a757fe5b602002015189600f6014811015156101bb57fe5b60200201518a60106014811015156101cf57fe5b60200201518b60116014811015156101e357fe5b60200201518c60126014811015156101f757fe5b60200201518d601360148110151561020b57fe5b6020020151610619565b905061021f6109cf565b604080519081016040528088600060148110151561023957fe5b6020020151815260200188600160148110151561025257fe5b602002015181525081600060038110151561026957fe5b6020020181905250604080519081016040528088600c60148110151561028b57fe5b6020020151815260200188600d6014811015156102a457fe5b60200201518152508160016003811015156102bb57fe5b602002018190525060408051908101604052808860026014811015156102dd57fe5b602002015181526020018860036014811015156102f657fe5b602002015181525081600260038110151561030d57fe5b6020020181905250610320818784610694565b1515610331576000925050506105ed565b6103396109cf565b604080519081016040528089600660148110151561035357fe5b6020020151815260200189600760148110151561036c57fe5b602002015181525081600060038110151561038357fe5b6020020181905250604080519081016040528089600e6014811015156103a557fe5b6020020151815260200189600f6014811015156103be57fe5b60200201518152508160016003811015156103d557fe5b602002018190525060408051908101604052808960086014811015156103f757fe5b6020020151815260200189600960148110151561041057fe5b602002015181525081600260038110151561042757fe5b602002018190525061043a818785610694565b151561044c57600093505050506105ed565b6104546109fd565b60408051908101604052808a601060148110151561046e57fe5b602002015181526020018a601160148110151561048757fe5b602002015181525081600060028110151561049e57fe5b602002018190525060408051908101604052808a60046014811015156104c057fe5b602002015181526020018a60056014811015156104d957fe5b60200201518152508160016002811015156104f057fe5b602002018190525061050481898887610744565b15156105175760009450505050506105ed565b61051f6109fd565b60408051908101604052808b601260148110151561053957fe5b602002015181526020018b601360148110151561055257fe5b602002015181525081600060028110151561056957fe5b602002018190525060408051908101604052808b600a60148110151561058b57fe5b602002015181526020018b600b6014811015156105a457fe5b60200201518152508160016002811015156105bb57fe5b60200201819052506105cf81898988610744565b15156105e3576000955050505050506105ed565b6001955050505050505b949350505050565b60028060000154908060010154905082565b60008060000154908060010154905082565b6000610686898989898989898960405160200180898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001985050505050505050506040516020818303038152906040528051906020012060019004610846565b905098975050505050505050565b600061069e610a2b565b6106e98560016003811015156106b057fe5b60200201516106db858860026003811015156106c857fe5b602002015161088090919063ffffffff16565b61092d90919063ffffffff16565b90506106f3610a2b565b6107198587600060038110151561070657fe5b602002015161088090919063ffffffff16565b905080600001518260000151148015610739575080602001518260200151145b925050509392505050565b600061074e610a2b565b6107c56107848560006040805190810160405290816000820154815260200160018201548152505061088090919063ffffffff16565b6107b78760026040805190810160405290816000820154815260200160018201548152505061088090919063ffffffff16565b61092d90919063ffffffff16565b90506107cf610a2b565b61081a6107f8858960016002811015156107e557fe5b602002015161088090919063ffffffff16565b88600060028110151561080757fe5b602002015161092d90919063ffffffff16565b90508060000151826000015114801561083a575080602001518260200151145b92505050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381151561087757fe5b06915050919050565b610888610a2b565b600182141561089957829050610927565b60028214156108b3576108ac838461092d565b9050610927565b6108bb610a45565b83600001518160006003811015156108cf57fe5b60200201818152505083602001518160016003811015156108ec57fe5b6020020181815250508281600260038110151561090557fe5b6020020181815250506040826060836007600019fa151561092557600080fd5b505b92915050565b610935610a2b565b61093d610a68565b836000015181600060048110151561095157fe5b602002018181525050836020015181600160048110151561096e57fe5b602002018181525050826000015181600260048110151561098b57fe5b60200201818152505082602001518160036004811015156109a857fe5b6020020181815250506040826080836006600019fa15156109c857600080fd5b5092915050565b60c0604051908101604052806003905b6109e7610a8b565b8152602001906001900390816109df5790505090565b6080604051908101604052806002905b610a15610a8b565b815260200190600190039081610a0d5790505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b60408051908101604052806000815260200160008152509056fea165627a7a72305820a90b73380ce442a40e7021c99a7ebedcc8566c806e1888e19852a9c56b849c490029`

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
