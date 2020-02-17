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
const DlesigmaverifierABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[12]\"},{\"name\":\"z\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"verifyDLESigmaProofWithNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[12]\"},{\"name\":\"z\",\"type\":\"uint256\"},{\"name\":\"input\",\"type\":\"uint256[]\"}],\"name\":\"verifyDLESigmaProofWithCustom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[12]\"},{\"name\":\"z\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"addr\",\"type\":\"uint256\"}],\"name\":\"verifyDLESigmaProofWithCustom2\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[12]\"},{\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"verifyDLESigmaProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DlesigmaverifierBin is the compiled bytecode used for deploying new contracts.
const DlesigmaverifierBin = `0x608060405234801561001057600080fd5b50610da3806100206000396000f3fe608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304e72d2b146100675780630581c040146101115780632e26d0b8146101b15780639751cb13146102db575b600080fd5b34801561007357600080fd5b506100f760048036036101e081101561008b57600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291908035906020019092919080359060200190929190505050610371565b604051808215151515815260200191505060405180910390f35b34801561011d57600080fd5b5061019760048036036101c081101561013557600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190803590602001909291905050506104cc565b604051808215151515815260200191505060405180910390f35b3480156101bd57600080fd5b506102c160048036036101c08110156101d557600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291908035906020019064010000000081111561023e57600080fd5b82018360208201111561025057600080fd5b8035906020019184602083028401116401000000008311171561027257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050919291929050505061053f565b604051808215151515815260200191505060405180910390f35b3480156102e757600080fd5b5061035760048036036101a08110156102ff57600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919050505061063c565b604051808215151515815260200191505060405180910390f35b6000606060066040519080825280602002602001820160405280156103a55781602001602082028038833980820191505090505b509050856000600c811015156103b757fe5b60200201518160008151811015156103cb57fe5b9060200190602002018181525050856001600c811015156103e857fe5b60200201518160018151811015156103fc57fe5b9060200190602002018181525050856002600c8110151561041957fe5b602002015181600281518110151561042d57fe5b9060200190602002018181525050856003600c8110151561044a57fe5b602002015181600381518110151561045e57fe5b90602001906020020181815250508381600481518110151561047c57fe5b90602001906020020181815250508281600581518110151561049a57fe5b906020019060200201818152505060006104b3826106ad565b90506104c0878783610716565b92505050949350505050565b600080610528856000600c811015156104e157fe5b6020020151866001600c811015156104f557fe5b6020020151876002600c8110151561050957fe5b6020020151886003600c8110151561051d57fe5b60200201518761097d565b9050610535858583610716565b9150509392505050565b6000606082516004016040519080825280602002602001820160405280156105765781602001602082028038833980820191505090505b50905060008090505b60048110156105c4578581600c8110151561059657fe5b602002015182828151811015156105a957fe5b9060200190602002018181525050808060010191505061057f565b5060008090505b83518110156106185783818151811015156105e257fe5b9060200190602002015182600483018151811015156105fd57fe5b906020019060200201818152505080806001019150506105cb565b506000610624826106ad565b9050610631868683610716565b925050509392505050565b600080610697846000600c8110151561065157fe5b6020020151856001600c8110151561066557fe5b6020020151866002600c8110151561067957fe5b6020020151876003600c8110151561068d57fe5b60200201516109dd565b90506106a4848483610716565b91505092915050565b600061070f8260405160200180828051906020019060200280838360005b838110156106e65780820151818401526020810190506106cb565b505050509050019150506040516020818303038152906040528051906020012060019004610a34565b9050919050565b6000610720610cf4565b846004600c8110151561072f57fe5b602002015181600060068110151561074357fe5b602002018181525050846005600c8110151561075b57fe5b602002015181600160068110151561076f57fe5b602002018181525050846000600c8110151561078757fe5b602002015181600260068110151561079b57fe5b602002018181525050846001600c811015156107b357fe5b60200201518160036006811015156107c757fe5b602002018181525050846006600c811015156107df57fe5b60200201518160046006811015156107f357fe5b602002018181525050846007600c8110151561080b57fe5b602002015181600560068110151561081f57fe5b602002018181525050610833818585610a6e565b1515610843576000915050610976565b61084b610cf4565b856008600c8110151561085a57fe5b602002015181600060068110151561086e57fe5b602002018181525050856009600c8110151561088657fe5b602002015181600160068110151561089a57fe5b602002018181525050856002600c811015156108b257fe5b60200201518160026006811015156108c657fe5b602002018181525050856003600c811015156108de57fe5b60200201518160036006811015156108f257fe5b60200201818152505085600a600c8110151561090a57fe5b602002015181600460068110151561091e57fe5b60200201818152505085600b600c8110151561093657fe5b602002015181600560068110151561094a57fe5b60200201818152505061095e818686610a6e565b151561096f57600092505050610976565b6001925050505b9392505050565b60006109d2868686868660405160200180868152602001858152602001848152602001838152602001828152602001955050505050506040516020818303038152906040528051906020012060019004610a34565b905095945050505050565b6000610a2a85858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060019004610a34565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083811515610a6557fe5b06915050919050565b6000610a78610d17565b6040805190810160405280866000600681101515610a9257fe5b60200201518152602001866001600681101515610aab57fe5b60200201518152509050610abd610d17565b6040805190810160405280876002600681101515610ad757fe5b60200201518152602001876003600681101515610af057fe5b60200201518152509050610b02610d17565b6040805190810160405280886004600681101515610b1c57fe5b60200201518152602001886005600681101515610b3557fe5b60200201518152509050610b528684610ba590919063ffffffff16565b9250610b7982610b6b8784610ba590919063ffffffff16565b610c5290919063ffffffff16565b905080600001518360000151148015610b99575080602001518360200151145b93505050509392505050565b610bad610d17565b6001821415610bbe57829050610c4c565b6002821415610bd857610bd18384610c52565b9050610c4c565b610be0610d31565b8360000151816000600381101515610bf457fe5b6020020181815250508360200151816001600381101515610c1157fe5b60200201818152505082816002600381101515610c2a57fe5b6020020181815250506040826060836007600019fa1515610c4a57600080fd5b505b92915050565b610c5a610d17565b610c62610d54565b8360000151816000600481101515610c7657fe5b6020020181815250508360200151816001600481101515610c9357fe5b6020020181815250508260000151816002600481101515610cb057fe5b6020020181815250508260200151816003600481101515610ccd57fe5b6020020181815250506040826080836006600019fa1515610ced57600080fd5b5092915050565b60c060405190810160405280600690602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a72305820a5082829da0aecc6a0355a897784c1708c61e6b678f1d021d30c3702e6dc1be10029`

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

// VerifyDLESigmaProofWithCustom is a free data retrieval call binding the contract method 0x2e26d0b8.
//
// Solidity: function verifyDLESigmaProofWithCustom(uint256[12] points, uint256 z, uint256[] input) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCaller) VerifyDLESigmaProofWithCustom(opts *bind.CallOpts, points [12]*big.Int, z *big.Int, input []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dlesigmaverifier.contract.Call(opts, out, "verifyDLESigmaProofWithCustom", points, z, input)
	return *ret0, err
}

// VerifyDLESigmaProofWithCustom is a free data retrieval call binding the contract method 0x2e26d0b8.
//
// Solidity: function verifyDLESigmaProofWithCustom(uint256[12] points, uint256 z, uint256[] input) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierSession) VerifyDLESigmaProofWithCustom(points [12]*big.Int, z *big.Int, input []*big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProofWithCustom(&_Dlesigmaverifier.CallOpts, points, z, input)
}

// VerifyDLESigmaProofWithCustom is a free data retrieval call binding the contract method 0x2e26d0b8.
//
// Solidity: function verifyDLESigmaProofWithCustom(uint256[12] points, uint256 z, uint256[] input) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCallerSession) VerifyDLESigmaProofWithCustom(points [12]*big.Int, z *big.Int, input []*big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProofWithCustom(&_Dlesigmaverifier.CallOpts, points, z, input)
}

// VerifyDLESigmaProofWithCustom2 is a free data retrieval call binding the contract method 0x04e72d2b.
//
// Solidity: function verifyDLESigmaProofWithCustom2(uint256[12] points, uint256 z, uint256 nonce, uint256 addr) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCaller) VerifyDLESigmaProofWithCustom2(opts *bind.CallOpts, points [12]*big.Int, z *big.Int, nonce *big.Int, addr *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dlesigmaverifier.contract.Call(opts, out, "verifyDLESigmaProofWithCustom2", points, z, nonce, addr)
	return *ret0, err
}

// VerifyDLESigmaProofWithCustom2 is a free data retrieval call binding the contract method 0x04e72d2b.
//
// Solidity: function verifyDLESigmaProofWithCustom2(uint256[12] points, uint256 z, uint256 nonce, uint256 addr) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierSession) VerifyDLESigmaProofWithCustom2(points [12]*big.Int, z *big.Int, nonce *big.Int, addr *big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProofWithCustom2(&_Dlesigmaverifier.CallOpts, points, z, nonce, addr)
}

// VerifyDLESigmaProofWithCustom2 is a free data retrieval call binding the contract method 0x04e72d2b.
//
// Solidity: function verifyDLESigmaProofWithCustom2(uint256[12] points, uint256 z, uint256 nonce, uint256 addr) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCallerSession) VerifyDLESigmaProofWithCustom2(points [12]*big.Int, z *big.Int, nonce *big.Int, addr *big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProofWithCustom2(&_Dlesigmaverifier.CallOpts, points, z, nonce, addr)
}

// VerifyDLESigmaProofWithNonce is a free data retrieval call binding the contract method 0x0581c040.
//
// Solidity: function verifyDLESigmaProofWithNonce(uint256[12] points, uint256 z, uint256 nonce) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCaller) VerifyDLESigmaProofWithNonce(opts *bind.CallOpts, points [12]*big.Int, z *big.Int, nonce *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dlesigmaverifier.contract.Call(opts, out, "verifyDLESigmaProofWithNonce", points, z, nonce)
	return *ret0, err
}

// VerifyDLESigmaProofWithNonce is a free data retrieval call binding the contract method 0x0581c040.
//
// Solidity: function verifyDLESigmaProofWithNonce(uint256[12] points, uint256 z, uint256 nonce) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierSession) VerifyDLESigmaProofWithNonce(points [12]*big.Int, z *big.Int, nonce *big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProofWithNonce(&_Dlesigmaverifier.CallOpts, points, z, nonce)
}

// VerifyDLESigmaProofWithNonce is a free data retrieval call binding the contract method 0x0581c040.
//
// Solidity: function verifyDLESigmaProofWithNonce(uint256[12] points, uint256 z, uint256 nonce) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCallerSession) VerifyDLESigmaProofWithNonce(points [12]*big.Int, z *big.Int, nonce *big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProofWithNonce(&_Dlesigmaverifier.CallOpts, points, z, nonce)
}
