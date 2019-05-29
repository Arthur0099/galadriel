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

// DlesigmaverifierABI is the input ABI used to generate the binding from.
const DlesigmaverifierABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[12]\"},{\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"verifyDLESigmaProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DlesigmaverifierBin is the compiled bytecode used for deploying new contracts.
const DlesigmaverifierBin = `0x608060405234801561001057600080fd5b50610766806100206000396000f3fe608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680639751cb1314610046575b600080fd5b34801561005257600080fd5b506100c260048036036101a081101561006a57600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291905050506100dc565b604051808215151515815260200191505060405180910390f35b600080610137846000600c811015156100f157fe5b6020020151856001600c8110151561010557fe5b6020020151866002600c8110151561011957fe5b6020020151876003600c8110151561012d57fe5b60200201516103a0565b90506101416106b7565b846004600c8110151561015057fe5b602002015181600060068110151561016457fe5b602002018181525050846005600c8110151561017c57fe5b602002015181600160068110151561019057fe5b602002018181525050846000600c811015156101a857fe5b60200201518160026006811015156101bc57fe5b602002018181525050846001600c811015156101d457fe5b60200201518160036006811015156101e857fe5b602002018181525050846006600c8110151561020057fe5b602002015181600460068110151561021457fe5b602002018181525050846007600c8110151561022c57fe5b602002015181600560068110151561024057fe5b6020020181815250506102548185846103f7565b15156102655760009250505061039a565b61026d6106b7565b856008600c8110151561027c57fe5b602002015181600060068110151561029057fe5b602002018181525050856009600c811015156102a857fe5b60200201518160016006811015156102bc57fe5b602002018181525050856002600c811015156102d457fe5b60200201518160026006811015156102e857fe5b602002018181525050856003600c8110151561030057fe5b602002015181600360068110151561031457fe5b60200201818152505085600a600c8110151561032c57fe5b602002015181600460068110151561034057fe5b60200201818152505085600b600c8110151561035857fe5b602002015181600560068110151561036c57fe5b6020020181815250506103808186856103f7565b1515610392576000935050505061039a565b600193505050505b92915050565b60006103ed8585858560405160200180858152602001848152602001838152602001828152602001945050505050604051602081830303815290604052805190602001206001900461052e565b9050949350505050565b60006104016106da565b604080519081016040528086600060068110151561041b57fe5b6020020151815260200186600160068110151561043457fe5b602002015181525090506104466106da565b604080519081016040528087600260068110151561046057fe5b6020020151815260200187600360068110151561047957fe5b6020020151815250905061048b6106da565b60408051908101604052808860046006811015156104a557fe5b602002015181526020018860056006811015156104be57fe5b602002015181525090506104db868461056890919063ffffffff16565b9250610502826104f4878461056890919063ffffffff16565b61061590919063ffffffff16565b905080600001518360000151148015610522575080602001518360200151145b93505050509392505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381151561055f57fe5b06915050919050565b6105706106da565b60018214156105815782905061060f565b600282141561059b576105948384610615565b905061060f565b6105a36106f4565b83600001518160006003811015156105b757fe5b60200201818152505083602001518160016003811015156105d457fe5b602002018181525050828160026003811015156105ed57fe5b6020020181815250506040826060836007600019fa151561060d57600080fd5b505b92915050565b61061d6106da565b610625610717565b836000015181600060048110151561063957fe5b602002018181525050836020015181600160048110151561065657fe5b602002018181525050826000015181600260048110151561067357fe5b602002018181525050826020015181600360048110151561069057fe5b6020020181815250506040826080836006600019fa15156106b057600080fd5b5092915050565b60c060405190810160405280600690602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a72305820f99f7a39d97559b8e52e5204a6cb2ab7117035a4327da5f39aa3c8b5a77c92200029`

// DeployDlesigmaverifier deploys a new Ethereum contract, binding an instance of Dlesigmaverifier to it.
func DeployDlesigmaverifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dlesigmaverifier, error) {
	parsed, err := abi.JSON(strings.NewReader(DlesigmaverifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DlesigmaverifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dlesigmaverifier{DlesigmaverifierCaller: DlesigmaverifierCaller{contract: contract}, DlesigmaverifierTransactor: DlesigmaverifierTransactor{contract: contract}, DlesigmaverifierFilterer: DlesigmaverifierFilterer{contract: contract}}, nil
}

// Dlesigmaverifier is an auto generated Go binding around an Ethereum contract.
type Dlesigmaverifier struct {
	DlesigmaverifierCaller     // Read-only binding to the contract
	DlesigmaverifierTransactor // Write-only binding to the contract
	DlesigmaverifierFilterer   // Log filterer for contract events
}

// DlesigmaverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type DlesigmaverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DlesigmaverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DlesigmaverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DlesigmaverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DlesigmaverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DlesigmaverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DlesigmaverifierSession struct {
	Contract     *Dlesigmaverifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DlesigmaverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DlesigmaverifierCallerSession struct {
	Contract *DlesigmaverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DlesigmaverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DlesigmaverifierTransactorSession struct {
	Contract     *DlesigmaverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DlesigmaverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type DlesigmaverifierRaw struct {
	Contract *Dlesigmaverifier // Generic contract binding to access the raw methods on
}

// DlesigmaverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DlesigmaverifierCallerRaw struct {
	Contract *DlesigmaverifierCaller // Generic read-only contract binding to access the raw methods on
}

// DlesigmaverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DlesigmaverifierTransactorRaw struct {
	Contract *DlesigmaverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDlesigmaverifier creates a new instance of Dlesigmaverifier, bound to a specific deployed contract.
func NewDlesigmaverifier(address common.Address, backend bind.ContractBackend) (*Dlesigmaverifier, error) {
	contract, err := bindDlesigmaverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dlesigmaverifier{DlesigmaverifierCaller: DlesigmaverifierCaller{contract: contract}, DlesigmaverifierTransactor: DlesigmaverifierTransactor{contract: contract}, DlesigmaverifierFilterer: DlesigmaverifierFilterer{contract: contract}}, nil
}

// NewDlesigmaverifierCaller creates a new read-only instance of Dlesigmaverifier, bound to a specific deployed contract.
func NewDlesigmaverifierCaller(address common.Address, caller bind.ContractCaller) (*DlesigmaverifierCaller, error) {
	contract, err := bindDlesigmaverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DlesigmaverifierCaller{contract: contract}, nil
}

// NewDlesigmaverifierTransactor creates a new write-only instance of Dlesigmaverifier, bound to a specific deployed contract.
func NewDlesigmaverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*DlesigmaverifierTransactor, error) {
	contract, err := bindDlesigmaverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DlesigmaverifierTransactor{contract: contract}, nil
}

// NewDlesigmaverifierFilterer creates a new log filterer instance of Dlesigmaverifier, bound to a specific deployed contract.
func NewDlesigmaverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*DlesigmaverifierFilterer, error) {
	contract, err := bindDlesigmaverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DlesigmaverifierFilterer{contract: contract}, nil
}

// bindDlesigmaverifier binds a generic wrapper to an already deployed contract.
func bindDlesigmaverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DlesigmaverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dlesigmaverifier *DlesigmaverifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dlesigmaverifier.Contract.DlesigmaverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dlesigmaverifier *DlesigmaverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dlesigmaverifier.Contract.DlesigmaverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dlesigmaverifier *DlesigmaverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dlesigmaverifier.Contract.DlesigmaverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dlesigmaverifier *DlesigmaverifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dlesigmaverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dlesigmaverifier *DlesigmaverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dlesigmaverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dlesigmaverifier *DlesigmaverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dlesigmaverifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyDLESigmaProof is a free data retrieval call binding the contract method 0x9751cb13.
//
// Solidity: function verifyDLESigmaProof(uint256[12] points, uint256 z) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCaller) VerifyDLESigmaProof(opts *bind.CallOpts, points [12]*big.Int, z *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dlesigmaverifier.contract.Call(opts, out, "verifyDLESigmaProof", points, z)
	return *ret0, err
}

// VerifyDLESigmaProof is a free data retrieval call binding the contract method 0x9751cb13.
//
// Solidity: function verifyDLESigmaProof(uint256[12] points, uint256 z) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierSession) VerifyDLESigmaProof(points [12]*big.Int, z *big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProof(&_Dlesigmaverifier.CallOpts, points, z)
}

// VerifyDLESigmaProof is a free data retrieval call binding the contract method 0x9751cb13.
//
// Solidity: function verifyDLESigmaProof(uint256[12] points, uint256 z) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCallerSession) VerifyDLESigmaProof(points [12]*big.Int, z *big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProof(&_Dlesigmaverifier.CallOpts, points, z)
}
