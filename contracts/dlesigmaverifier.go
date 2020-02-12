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
const DlesigmaverifierABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[12]\"},{\"name\":\"z\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"verifyDLESigmaProofWithNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[12]\"},{\"name\":\"z\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"addr\",\"type\":\"uint256\"}],\"name\":\"verifyDLESigmaProofWithCustom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[12]\"},{\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"verifyDLESigmaProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DlesigmaverifierBin is the compiled bytecode used for deploying new contracts.
const DlesigmaverifierBin = `0x608060405234801561001057600080fd5b50610a8b806100206000396000f3fe608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630581c0401461005c5780636907b66d146100fc5780639751cb13146101a6575b600080fd5b34801561006857600080fd5b506100e260048036036101c081101561008057600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291908035906020019092919050505061023c565b604051808215151515815260200191505060405180910390f35b34801561010857600080fd5b5061018c60048036036101e081101561012057600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080359060200190929190803590602001909291905050506102af565b604051808215151515815260200191505060405180910390f35b3480156101b257600080fd5b5061022260048036036101a08110156101ca57600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190505050610324565b604051808215151515815260200191505060405180910390f35b600080610298856000600c8110151561025157fe5b6020020151866001600c8110151561026557fe5b6020020151876002600c8110151561027957fe5b6020020151886003600c8110151561028d57fe5b602002015187610395565b90506102a58585836103f5565b9150509392505050565b60008061030c866000600c811015156102c457fe5b6020020151876001600c811015156102d857fe5b6020020151886002600c811015156102ec57fe5b6020020151896003600c8110151561030057fe5b6020020151888861065c565b90506103198686836103f5565b915050949350505050565b60008061037f846000600c8110151561033957fe5b6020020151856001600c8110151561034d57fe5b6020020151866002600c8110151561036157fe5b6020020151876003600c8110151561037557fe5b60200201516106c5565b905061038c8484836103f5565b91505092915050565b60006103ea86868686866040516020018086815260200185815260200184815260200183815260200182815260200195505050505050604051602081830303815290604052805190602001206001900461071c565b905095945050505050565b60006103ff6109dc565b846004600c8110151561040e57fe5b602002015181600060068110151561042257fe5b602002018181525050846005600c8110151561043a57fe5b602002015181600160068110151561044e57fe5b602002018181525050846000600c8110151561046657fe5b602002015181600260068110151561047a57fe5b602002018181525050846001600c8110151561049257fe5b60200201518160036006811015156104a657fe5b602002018181525050846006600c811015156104be57fe5b60200201518160046006811015156104d257fe5b602002018181525050846007600c811015156104ea57fe5b60200201518160056006811015156104fe57fe5b602002018181525050610512818585610756565b1515610522576000915050610655565b61052a6109dc565b856008600c8110151561053957fe5b602002015181600060068110151561054d57fe5b602002018181525050856009600c8110151561056557fe5b602002015181600160068110151561057957fe5b602002018181525050856002600c8110151561059157fe5b60200201518160026006811015156105a557fe5b602002018181525050856003600c811015156105bd57fe5b60200201518160036006811015156105d157fe5b60200201818152505085600a600c811015156105e957fe5b60200201518160046006811015156105fd57fe5b60200201818152505085600b600c8110151561061557fe5b602002015181600560068110151561062957fe5b60200201818152505061063d818686610756565b151561064e57600092505050610655565b6001925050505b9392505050565b60006106b9878787878787604051602001808781526020018681526020018581526020018481526020018381526020018281526020019650505050505050604051602081830303815290604052805190602001206001900461071c565b90509695505050505050565b60006107128585858560405160200180858152602001848152602001838152602001828152602001945050505050604051602081830303815290604052805190602001206001900461071c565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381151561074d57fe5b06915050919050565b60006107606109ff565b604080519081016040528086600060068110151561077a57fe5b6020020151815260200186600160068110151561079357fe5b602002015181525090506107a56109ff565b60408051908101604052808760026006811015156107bf57fe5b602002015181526020018760036006811015156107d857fe5b602002015181525090506107ea6109ff565b604080519081016040528088600460068110151561080457fe5b6020020151815260200188600560068110151561081d57fe5b6020020151815250905061083a868461088d90919063ffffffff16565b925061086182610853878461088d90919063ffffffff16565b61093a90919063ffffffff16565b905080600001518360000151148015610881575080602001518360200151145b93505050509392505050565b6108956109ff565b60018214156108a657829050610934565b60028214156108c0576108b9838461093a565b9050610934565b6108c8610a19565b83600001518160006003811015156108dc57fe5b60200201818152505083602001518160016003811015156108f957fe5b6020020181815250508281600260038110151561091257fe5b6020020181815250506040826060836007600019fa151561093257600080fd5b505b92915050565b6109426109ff565b61094a610a3c565b836000015181600060048110151561095e57fe5b602002018181525050836020015181600160048110151561097b57fe5b602002018181525050826000015181600260048110151561099857fe5b60200201818152505082602001518160036004811015156109b557fe5b6020020181815250506040826080836006600019fa15156109d557600080fd5b5092915050565b60c060405190810160405280600690602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a72305820ecab467d0b2c07b416ca832d0eb2414653e1a6cabe89b751bd9fcf2fd965724f0029`

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

// VerifyDLESigmaProofWithCustom is a free data retrieval call binding the contract method 0x6907b66d.
//
// Solidity: function verifyDLESigmaProofWithCustom(uint256[12] points, uint256 z, uint256 nonce, uint256 addr) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCaller) VerifyDLESigmaProofWithCustom(opts *bind.CallOpts, points [12]*big.Int, z *big.Int, nonce *big.Int, addr *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dlesigmaverifier.contract.Call(opts, out, "verifyDLESigmaProofWithCustom", points, z, nonce, addr)
	return *ret0, err
}

// VerifyDLESigmaProofWithCustom is a free data retrieval call binding the contract method 0x6907b66d.
//
// Solidity: function verifyDLESigmaProofWithCustom(uint256[12] points, uint256 z, uint256 nonce, uint256 addr) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierSession) VerifyDLESigmaProofWithCustom(points [12]*big.Int, z *big.Int, nonce *big.Int, addr *big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProofWithCustom(&_Dlesigmaverifier.CallOpts, points, z, nonce, addr)
}

// VerifyDLESigmaProofWithCustom is a free data retrieval call binding the contract method 0x6907b66d.
//
// Solidity: function verifyDLESigmaProofWithCustom(uint256[12] points, uint256 z, uint256 nonce, uint256 addr) constant returns(bool)
func (_Dlesigmaverifier *DlesigmaverifierCallerSession) VerifyDLESigmaProofWithCustom(points [12]*big.Int, z *big.Int, nonce *big.Int, addr *big.Int) (bool, error) {
	return _Dlesigmaverifier.Contract.VerifyDLESigmaProofWithCustom(&_Dlesigmaverifier.CallOpts, points, z, nonce, addr)
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
