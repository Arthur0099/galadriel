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

// PublicparamsABI is the input ABI used to generate the binding from.
const PublicparamsABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"aggSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gVector\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hVector\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"u\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getG\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getH\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getU\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAggGVector\",\"outputs\":[{\"internalType\":\"uint256[256]\",\"name\":\"\",\"type\":\"uint256[256]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAggHVector\",\"outputs\":[{\"internalType\":\"uint256[256]\",\"name\":\"\",\"type\":\"uint256[256]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGVector\",\"outputs\":[{\"internalType\":\"uint256[128]\",\"name\":\"\",\"type\":\"uint256[128]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHVector\",\"outputs\":[{\"internalType\":\"uint256[128]\",\"name\":\"\",\"type\":\"uint256[128]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PublicparamsBin is the compiled bytecode used for deploying new contracts.
const PublicparamsBin = `0x60806040523480156200001157600080fd5b50600160046000018190555060026004600101819055506200006e6040518060400160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c670000000000008152506200015a60201b60201c565b600080820151816000015560208201518160010155905050620000cc6040518060400160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c670000000000008152506200015a60201b60201c565b600260008201518160000155602082015181600101559050506200012b6040518060400160405280601b81526020017f752067656e657261746f72206f6620696e6e657270726f6475637400000000008152506200015a60201b60201c565b60066000820151816000015560208201518160010155905050620001546200023660201b60201c565b62000684565b6200016462000626565b6000620001f2836040516020018082805190602001908083835b60208310620001a357805182526020820191506020810190506020830392506200017e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b90506200022e81600460405180604001604052908160008201548152602001600182015481525050620004cc60201b6200095a1790919060201c565b915050919050565b6000806200029860405160200180807f677673000000000000000000000000000000000000000000000000000000000081525060030190506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b905060008090505b60026040028110156200036157620002ee818301604051602001808281526020019150506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b92506200032a83600460405180604001604052908160008201548152602001600182015481525050620004cc60201b6200095a1790919060201c565b600882608081106200033857fe5b6002020160008201518160000155602082015181600101559050508080600101915050620002a0565b506000620003c360405160200180807f687673000000000000000000000000000000000000000000000000000000000081525060030190506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b905060008090505b60026040028110156200048d5762000419828201604051602001808281526020019150506040516020818303038152906040528051906020012060001c6200049360201b620009221760201c565b93506200045584600460405180604001604052908160008201548152602001600182015481525050620004cc60201b6200095a1790919060201c565b61010882608081106200046457fe5b6002020160008201518160000155602082015181600101559050508080600101915050620003cb565b50505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381620004c357fe5b06915050919050565b620004d662000626565b6001821415620004e9578290506200057f565b60028214156200050d576200050583846200058560201b60201c565b90506200057f565b6200051762000640565b8360000151816000600381106200052a57fe5b6020020181815250508360200151816001600381106200054657fe5b60200201818152505082816002600381106200055e57fe5b6020020181815250506040826060836007600019fa6200057d57600080fd5b505b92915050565b6200058f62000626565b6200059962000662565b836000015181600060048110620005ac57fe5b602002018181525050836020015181600160048110620005c857fe5b602002018181525050826000015181600260048110620005e457fe5b6020020181815250508260200151816003600481106200060057fe5b6020020181815250506040826080836006600019fa6200061f57600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b610b2a80620006946000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806356e2736c116100a2578063c6a898c511610071578063c6a898c5146103ce578063da2b99ce146103f3578063da8972241461043a578063e2179b8e14610458578063ffe237f01461047d57610116565b806356e2736c146102fc5780637982ebcc1461034557806382529fdb14610363578063b8c9d365146103a957610116565b80633e8d3764116100e95780633e8d37641461020e5780633e9552251461022c578063483767f01461024a5780634db118751461029057806352c9b965146102d757610116565b806304c09ce91461011b5780630c00f8a01461016157806324d6147d146101aa5780632e52d606146101f0575b600080fd5b6101236104c3565b6040518082600260200280838360005b8381101561014e578082015181840152602081019050610133565b5050505090500191505060405180910390f35b61018d6004803603602081101561017757600080fd5b8101908080359060200190929190505050610511565b604051808381526020018281526020019250505060405180910390f35b6101b2610538565b6040518082600260200280838360005b838110156101dd5780820151818401526020810190506101c2565b5050505090500191505060405180910390f35b6101f8610587565b6040518082815260200191505060405180910390f35b61021661058c565b6040518082815260200191505060405180910390f35b610234610591565b6040518082815260200191505060405180910390f35b61025261059a565b6040518082608060200280838360005b8381101561027d578082015181840152602081019050610262565b5050505090500191505060405180910390f35b61029861062d565b604051808261010060200280838360005b838110156102c45780820151818401526020810190506102a9565b5050505090500191505060405180910390f35b6102df6106c3565b604051808381526020018281526020019250505060405180910390f35b6103286004803603602081101561031257600080fd5b81019080803590602001909291905050506106d5565b604051808381526020018281526020019250505060405180910390f35b61034d6106fd565b6040518082815260200191505060405180910390f35b61036b610702565b6040518082600260200280838360005b8381101561039657808201518184015260208101905061037b565b5050505090500191505060405180910390f35b6103b1610751565b604051808381526020018281526020019250505060405180910390f35b6103d6610763565b604051808381526020018281526020019250505060405180910390f35b6103fb610775565b604051808261010060200280838360005b8381101561042757808201518184015260208101905061040c565b5050505090500191505060405180910390f35b61044261080d565b6040518082815260200191505060405180910390f35b610460610816565b604051808381526020018281526020019250505060405180910390f35b610485610828565b6040518082608060200280838360005b838110156104b0578082015181840152602081019050610495565b5050505090500191505060405180910390f35b6104cb6108b9565b6104d36108b9565b6000800154816000600281106104e557fe5b6020020181815250506000600101548160016002811061050157fe5b6020020181815250508091505090565b6008816080811061051e57fe5b600202016000915090508060000154908060010154905082565b6105406108b9565b6105486108b9565b6006600001548160006002811061055b57fe5b6020020181815250506006600101548160016002811061057757fe5b6020020181815250508091505090565b600681565b604081565b60006006905090565b6105a26108db565b6105aa6108db565b60008090505b60408110156106255761010881608081106105c757fe5b60020201600001548282600202608081106105de57fe5b60200201818152505061010881608081106105f557fe5b600202016001015482600183600202016080811061060f57fe5b60200201818152505080806001019150506105b0565b508091505090565b6106356108fe565b61063d6108fe565b60008090505b60026040028110156106bb576008816080811061065c57fe5b60020201600001548282600202610100811061067457fe5b6020020181815250506008816080811061068a57fe5b6002020160010154826001836002020161010081106106a557fe5b6020020181815250508080600101915050610643565b508091505090565b60048060000154908060010154905082565b61010881608081106106e357fe5b600202016000915090508060000154908060010154905082565b600281565b61070a6108b9565b6107126108b9565b6002600001548160006002811061072557fe5b6020020181815250506002600101548160016002811061074157fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b61077d6108fe565b6107856108fe565b60008090505b60026040028110156108055761010881608081106107a557fe5b6002020160000154828260020261010081106107bd57fe5b60200201818152505061010881608081106107d457fe5b6002020160010154826001836002020161010081106107ef57fe5b602002018181525050808060010191505061078b565b508091505090565b60006040905090565b60008060000154908060010154905082565b6108306108db565b6108386108db565b60008090505b60408110156108b1576008816080811061085457fe5b600202016000015482826002026080811061086b57fe5b6020020181815250506008816080811061088157fe5b600202016001015482600183600202016080811061089b57fe5b602002018181525050808060010191505061083e565b508091505090565b6040518060400160405280600290602082028038833980820191505090505090565b604051806110000160405280608090602082028038833980820191505090505090565b60405180612000016040528061010090602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161095157fe5b06915050919050565b610962610a97565b6001821415610973578290506109f9565b600282141561098d5761098683846109ff565b90506109f9565b610995610ab1565b8360000151816000600381106109a757fe5b6020020181815250508360200151816001600381106109c257fe5b60200201818152505082816002600381106109d957fe5b6020020181815250506040826060836007600019fa6109f757600080fd5b505b92915050565b610a07610a97565b610a0f610ad3565b836000015181600060048110610a2157fe5b602002018181525050836020015181600160048110610a3c57fe5b602002018181525050826000015181600260048110610a5757fe5b602002018181525050826020015181600360048110610a7257fe5b6020020181815250506040826080836006600019fa610a9057600080fd5b5092915050565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b604051806080016040528060049060208202803883398082019150509050509056fea265627a7a723158208d873b904394af55d1e70458846e343074672366748dc8d0d1f22405d761ef7164736f6c63430005100032`

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

// GVector is a free data retrieval call binding the contract method 0x0c00f8a0.
//
// Solidity: function gVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) GVector(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "gVector", arg0)
	return *ret, err
}

// GVector is a free data retrieval call binding the contract method 0x0c00f8a0.
//
// Solidity: function gVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) GVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GVector(&_Publicparams.CallOpts, arg0)
}

// GVector is a free data retrieval call binding the contract method 0x0c00f8a0.
//
// Solidity: function gVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) GVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.GVector(&_Publicparams.CallOpts, arg0)
}

// GetAggGVector is a free data retrieval call binding the contract method 0x4db11875.
//
// Solidity: function getAggGVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsCaller) GetAggGVector(opts *bind.CallOpts) ([256]*big.Int, error) {
	var (
		ret0 = new([256]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getAggGVector")
	return *ret0, err
}

// GetAggGVector is a free data retrieval call binding the contract method 0x4db11875.
//
// Solidity: function getAggGVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsSession) GetAggGVector() ([256]*big.Int, error) {
	return _Publicparams.Contract.GetAggGVector(&_Publicparams.CallOpts)
}

// GetAggGVector is a free data retrieval call binding the contract method 0x4db11875.
//
// Solidity: function getAggGVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsCallerSession) GetAggGVector() ([256]*big.Int, error) {
	return _Publicparams.Contract.GetAggGVector(&_Publicparams.CallOpts)
}

// GetAggHVector is a free data retrieval call binding the contract method 0xda2b99ce.
//
// Solidity: function getAggHVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsCaller) GetAggHVector(opts *bind.CallOpts) ([256]*big.Int, error) {
	var (
		ret0 = new([256]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getAggHVector")
	return *ret0, err
}

// GetAggHVector is a free data retrieval call binding the contract method 0xda2b99ce.
//
// Solidity: function getAggHVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsSession) GetAggHVector() ([256]*big.Int, error) {
	return _Publicparams.Contract.GetAggHVector(&_Publicparams.CallOpts)
}

// GetAggHVector is a free data retrieval call binding the contract method 0xda2b99ce.
//
// Solidity: function getAggHVector() constant returns(uint256[256])
func (_Publicparams *PublicparamsCallerSession) GetAggHVector() ([256]*big.Int, error) {
	return _Publicparams.Contract.GetAggHVector(&_Publicparams.CallOpts)
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

// GetGVector is a free data retrieval call binding the contract method 0xffe237f0.
//
// Solidity: function getGVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsCaller) GetGVector(opts *bind.CallOpts) ([128]*big.Int, error) {
	var (
		ret0 = new([128]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getGVector")
	return *ret0, err
}

// GetGVector is a free data retrieval call binding the contract method 0xffe237f0.
//
// Solidity: function getGVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsSession) GetGVector() ([128]*big.Int, error) {
	return _Publicparams.Contract.GetGVector(&_Publicparams.CallOpts)
}

// GetGVector is a free data retrieval call binding the contract method 0xffe237f0.
//
// Solidity: function getGVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsCallerSession) GetGVector() ([128]*big.Int, error) {
	return _Publicparams.Contract.GetGVector(&_Publicparams.CallOpts)
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

// GetHVector is a free data retrieval call binding the contract method 0x483767f0.
//
// Solidity: function getHVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsCaller) GetHVector(opts *bind.CallOpts) ([128]*big.Int, error) {
	var (
		ret0 = new([128]*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getHVector")
	return *ret0, err
}

// GetHVector is a free data retrieval call binding the contract method 0x483767f0.
//
// Solidity: function getHVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsSession) GetHVector() ([128]*big.Int, error) {
	return _Publicparams.Contract.GetHVector(&_Publicparams.CallOpts)
}

// GetHVector is a free data retrieval call binding the contract method 0x483767f0.
//
// Solidity: function getHVector() constant returns(uint256[128])
func (_Publicparams *PublicparamsCallerSession) GetHVector() ([128]*big.Int, error) {
	return _Publicparams.Contract.GetHVector(&_Publicparams.CallOpts)
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() constant returns(uint256)
func (_Publicparams *PublicparamsCaller) GetN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Publicparams.contract.Call(opts, out, "getN")
	return *ret0, err
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() constant returns(uint256)
func (_Publicparams *PublicparamsSession) GetN() (*big.Int, error) {
	return _Publicparams.Contract.GetN(&_Publicparams.CallOpts)
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() constant returns(uint256)
func (_Publicparams *PublicparamsCallerSession) GetN() (*big.Int, error) {
	return _Publicparams.Contract.GetN(&_Publicparams.CallOpts)
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

// HVector is a free data retrieval call binding the contract method 0x56e2736c.
//
// Solidity: function hVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCaller) HVector(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Publicparams.contract.Call(opts, out, "hVector", arg0)
	return *ret, err
}

// HVector is a free data retrieval call binding the contract method 0x56e2736c.
//
// Solidity: function hVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsSession) HVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.HVector(&_Publicparams.CallOpts, arg0)
}

// HVector is a free data retrieval call binding the contract method 0x56e2736c.
//
// Solidity: function hVector(uint256 ) constant returns(uint256 X, uint256 Y)
func (_Publicparams *PublicparamsCallerSession) HVector(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Publicparams.Contract.HVector(&_Publicparams.CallOpts, arg0)
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
