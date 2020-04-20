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
const DlesigmaverifierABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[12]\",\"name\":\"points\",\"type\":\"uint256[12]\"},{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"input\",\"type\":\"uint256[]\"}],\"name\":\"verifyDLESigmaProofWithCustom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[12]\",\"name\":\"points\",\"type\":\"uint256[12]\"},{\"internalType\":\"uint256\",\"name\":\"z\",\"type\":\"uint256\"}],\"name\":\"verifyDLESigmaProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DlesigmaverifierBin is the compiled bytecode used for deploying new contracts.
const DlesigmaverifierBin = `0x608060405234801561001057600080fd5b506109c1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e26d0b81461003b5780639751cb1314610158575b600080fd5b61013e60048036036101c081101561005257600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190803590602001906401000000008111156100bb57600080fd5b8201836020820111156100cd57600080fd5b803590602001918460208302840111640100000000831117156100ef57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506101e1565b604051808215151515815260200191505060405180910390f35b6101c760048036036101a081101561016f57600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919050505061037a565b604051808215151515815260200191505060405180910390f35b6000606060046040519080825280602002602001820160405280156102155781602001602082028038833980820191505090505b50905060008090505b600481101561025d578581600c811061023357fe5b602002015182828151811061024457fe5b602002602001018181525050808060010191505061021e565b5060006102bf8260405160200180828051906020019060200280838360005b8381101561029757808201518184015260208101905061027c565b505050509050019150506040516020818303038152906040528051906020012060001c6103e3565b905060006103228560405160200180828051906020019060200280838360005b838110156102fa5780820151818401526020810190506102df565b505050509050019150506040516020818303038152906040528051906020012060001c6103e3565b90506000610360838360405160200180838152602001828152602001925050506040516020818303038152906040528051906020012060001c6103e3565b905061036d88888361041b565b9450505050509392505050565b6000806103cd846000600c811061038d57fe5b6020020151856001600c811061039f57fe5b6020020151866002600c81106103b157fe5b6020020151876003600c81106103c357fe5b602002015161064e565b90506103da84848361041b565b91505092915050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161041257fe5b06915050919050565b600061042561090c565b846004600c811061043257fe5b60200201518160006006811061044457fe5b602002018181525050846005600c811061045a57fe5b60200201518160016006811061046c57fe5b602002018181525050846000600c811061048257fe5b60200201518160026006811061049457fe5b602002018181525050846001600c81106104aa57fe5b6020020151816003600681106104bc57fe5b602002018181525050846006600c81106104d257fe5b6020020151816004600681106104e457fe5b602002018181525050846007600c81106104fa57fe5b60200201518160056006811061050c57fe5b6020020181815250506105208185856106a4565b61052e576000915050610647565b61053661090c565b856008600c811061054357fe5b60200201518160006006811061055557fe5b602002018181525050856009600c811061056b57fe5b60200201518160016006811061057d57fe5b602002018181525050856002600c811061059357fe5b6020020151816002600681106105a557fe5b602002018181525050856003600c81106105bb57fe5b6020020151816003600681106105cd57fe5b60200201818152505085600a600c81106105e357fe5b6020020151816004600681106105f557fe5b60200201818152505085600b600c811061060b57fe5b60200201518160056006811061061d57fe5b6020020181815250506106318186866106a4565b61064057600092505050610647565b6001925050505b9392505050565b600061069a85858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c6103e3565b9050949350505050565b60006106ae61092e565b6040518060400160405280866000600681106106c657fe5b60200201518152602001866001600681106106dd57fe5b602002015181525090506106ef61092e565b60405180604001604052808760026006811061070757fe5b602002015181526020018760036006811061071e57fe5b6020020151815250905061073061092e565b60405180604001604052808860046006811061074857fe5b602002015181526020018860056006811061075f57fe5b6020020151815250905061077c86846107cf90919063ffffffff16565b92506107a38261079587846107cf90919063ffffffff16565b61087490919063ffffffff16565b9050806000015183600001511480156107c3575080602001518360200151145b93505050509392505050565b6107d761092e565b60018214156107e85782905061086e565b6002821415610802576107fb8384610874565b905061086e565b61080a610948565b83600001518160006003811061081c57fe5b60200201818152505083602001518160016003811061083757fe5b602002018181525050828160026003811061084e57fe5b6020020181815250506040826060836007600019fa61086c57600080fd5b505b92915050565b61087c61092e565b61088461096a565b83600001518160006004811061089657fe5b6020020181815250508360200151816001600481106108b157fe5b6020020181815250508260000151816002600481106108cc57fe5b6020020181815250508260200151816003600481106108e757fe5b6020020181815250506040826080836006600019fa61090557600080fd5b5092915050565b6040518060c00160405280600690602082028038833980820191505090505090565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b604051806080016040528060049060208202803883398082019150509050509056fea265627a7a7231582010b50328b33cc1ebb1333bf4b609350c77fbe96383c9ff7ceb191082dcb9236f64736f6c63430005100032`

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
