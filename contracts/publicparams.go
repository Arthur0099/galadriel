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
const PublicparamsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gVector\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gBase\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hVector\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"aggSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"u\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"getG\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAggGVector\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[256]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAggHVector\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[256]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGVector\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[128]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHVector\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[128]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBitSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getN\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PublicparamsBin is the compiled bytecode used for deploying new contracts.
const PublicparamsBin = `0x60806040523480156200001157600080fd5b5060016004600001819055506002600460010181905550620000776040805190810160405280601a81526020017f672067656e657261746f72206f66207477697374656420656c670000000000008152506200017e640100000000026401000000009004565b600080820151816000015560208201518160010155905050620000de6040805190810160405280601a81526020017f682067656e657261746f72206f66207477697374656420656c670000000000008152506200017e640100000000026401000000009004565b60026000820151816000015560208201518160010155905050620001466040805190810160405280601b81526020017f752067656e657261746f72206f6620696e6e657270726f6475637400000000008152506200017e640100000000026401000000009004565b60066000820151816000015560208201518160010155905050620001786200026f640100000000026401000000009004565b6200071a565b62000188620006ba565b600062000222836040516020018082805190602001908083835b602083101515620001c95780518252602082019150602081019050602083039250620001a2565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b90506200026781600460408051908101604052908160008201548152602001600182015481525050620005456401000000000262000a5e179091906401000000009004565b915050919050565b600080620002db60405160200180807f6776730000000000000000000000000000000000000000000000000000000000815250600301905060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b905060008090505b6002604002811015620003b9576200033b8183016040516020018082815260200191505060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b92506200038083600460408051908101604052908160008201548152602001600182015481525050620005456401000000000262000a5e179091906401000000009004565b6008826080811015156200039057fe5b6002020160008201518160000155602082015181600101559050508080600101915050620002e3565b5060006200042560405160200180807f6876730000000000000000000000000000000000000000000000000000000000815250600301905060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b905060008090505b60026040028110156200050457620004858282016040516020018082815260200191505060405160208183030381529060405280519060200120600190046200050a6401000000000262000a24176401000000009004565b9350620004ca84600460408051908101604052908160008201548152602001600182015481525050620005456401000000000262000a5e179091906401000000009004565b61010882608081101515620004db57fe5b60020201600082015181600001556020820151816001015590505080806001019150506200042d565b50505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838115156200053c57fe5b06915050919050565b6200054f620006ba565b6001821415620005625782905062000609565b60028214156200058f576200058783846200060f640100000000026401000000009004565b905062000609565b62000599620006d4565b8360000151816000600381101515620005ae57fe5b6020020181815250508360200151816001600381101515620005cc57fe5b60200201818152505082816002600381101515620005e657fe5b6020020181815250506040826060836007600019fa15156200060757600080fd5b505b92915050565b62000619620006ba565b62000623620006f7565b83600001518160006004811015156200063857fe5b60200201818152505083602001518160016004811015156200065657fe5b60200201818152505082600001518160026004811015156200067457fe5b60200201818152505082602001518160036004811015156200069257fe5b6020020181815250506040826080836006600019fa1515620006b357600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b610c39806200072a6000396000f3fe6080604052600436106100fc576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304c09ce9146101015780630c00f8a01461015457806324d6147d146101aa5780632e52d606146101fd5780633e8d3764146102285780633e95522514610253578063483767f01461027e5780634db11875146102d157806352c9b9651461032557806356e2736c146103575780637982ebcc146103ad57806382529fdb146103d8578063b8c9d3651461042b578063c6a898c51461045d578063da2b99ce1461048f578063da897224146104e3578063e2179b8e1461050e578063ffe237f014610540575b600080fd5b34801561010d57600080fd5b50610116610593565b6040518082600260200280838360005b83811015610141578082015181840152602081019050610126565b5050505090500191505060405180910390f35b34801561016057600080fd5b5061018d6004803603602081101561017757600080fd5b81019080803590602001909291905050506105e5565b604051808381526020018281526020019250505060405180910390f35b3480156101b657600080fd5b506101bf61060e565b6040518082600260200280838360005b838110156101ea5780820151818401526020810190506101cf565b5050505090500191505060405180910390f35b34801561020957600080fd5b50610212610661565b6040518082815260200191505060405180910390f35b34801561023457600080fd5b5061023d610666565b6040518082815260200191505060405180910390f35b34801561025f57600080fd5b5061026861066b565b6040518082815260200191505060405180910390f35b34801561028a57600080fd5b50610293610674565b6040518082608060200280838360005b838110156102be5780820151818401526020810190506102a3565b5050505090500191505060405180910390f35b3480156102dd57600080fd5b506102e661070f565b604051808261010060200280838360005b838110156103125780820151818401526020810190506102f7565b5050505090500191505060405180910390f35b34801561033157600080fd5b5061033a6107ad565b604051808381526020018281526020019250505060405180910390f35b34801561036357600080fd5b506103906004803603602081101561037a57600080fd5b81019080803590602001909291905050506107bf565b604051808381526020018281526020019250505060405180910390f35b3480156103b957600080fd5b506103c26107e9565b6040518082815260200191505060405180910390f35b3480156103e457600080fd5b506103ed6107ee565b6040518082600260200280838360005b838110156104185780820151818401526020810190506103fd565b5050505090500191505060405180910390f35b34801561043757600080fd5b50610440610841565b604051808381526020018281526020019250505060405180910390f35b34801561046957600080fd5b50610472610853565b604051808381526020018281526020019250505060405180910390f35b34801561049b57600080fd5b506104a4610865565b604051808261010060200280838360005b838110156104d05780820151818401526020810190506104b5565b5050505090500191505060405180910390f35b3480156104ef57600080fd5b506104f8610905565b6040518082815260200191505060405180910390f35b34801561051a57600080fd5b5061052361090e565b604051808381526020018281526020019250505060405180910390f35b34801561054c57600080fd5b50610555610920565b6040518082608060200280838360005b83811015610580578082015181840152602081019050610565565b5050505090500191505060405180910390f35b61059b6109b9565b6105a36109b9565b60008001548160006002811015156105b757fe5b6020020181815250506000600101548160016002811015156105d557fe5b6020020181815250508091505090565b6008816080811015156105f457fe5b600202016000915090508060000154908060010154905082565b6106166109b9565b61061e6109b9565b60066000015481600060028110151561063357fe5b60200201818152505060066001015481600160028110151561065157fe5b6020020181815250508091505090565b600681565b604081565b60006006905090565b61067c6109db565b6106846109db565b60008090505b604081101561070757610108816080811015156106a357fe5b600202016000015482826002026080811015156106bc57fe5b602002018181525050610108816080811015156106d557fe5b600202016001015482600183600202016080811015156106f157fe5b602002018181525050808060010191505061068a565b508091505090565b6107176109ff565b61071f6109ff565b60008090505b60026040028110156107a55760088160808110151561074057fe5b600202016000015482826002026101008110151561075a57fe5b60200201818152505060088160808110151561077257fe5b600202016001015482600183600202016101008110151561078f57fe5b6020020181815250508080600101915050610725565b508091505090565b60048060000154908060010154905082565b610108816080811015156107cf57fe5b600202016000915090508060000154908060010154905082565b600281565b6107f66109b9565b6107fe6109b9565b60026000015481600060028110151561081357fe5b60200201818152505060026001015481600160028110151561083157fe5b6020020181815250508091505090565b60028060000154908060010154905082565b60068060000154908060010154905082565b61086d6109ff565b6108756109ff565b60008090505b60026040028110156108fd576101088160808110151561089757fe5b60020201600001548282600202610100811015156108b157fe5b602002018181525050610108816080811015156108ca57fe5b60020201600101548260018360020201610100811015156108e757fe5b602002018181525050808060010191505061087b565b508091505090565b60006040905090565b60008060000154908060010154905082565b6109286109db565b6109306109db565b60008090505b60408110156109b15760088160808110151561094e57fe5b6002020160000154828260020260808110151561096757fe5b60200201818152505060088160808110151561097f57fe5b6002020160010154826001836002020160808110151561099b57fe5b6020020181815250508080600101915050610936565b508091505090565b6040805190810160405280600290602082028038833980820191505090505090565b61100060405190810160405280608090602082028038833980820191505090505090565b6120006040519081016040528061010090602082028038833980820191505090505090565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083811515610a5557fe5b06915050919050565b610a66610bad565b6001821415610a7757829050610b05565b6002821415610a9157610a8a8384610b0b565b9050610b05565b610a99610bc7565b8360000151816000600381101515610aad57fe5b6020020181815250508360200151816001600381101515610aca57fe5b60200201818152505082816002600381101515610ae357fe5b6020020181815250506040826060836007600019fa1515610b0357600080fd5b505b92915050565b610b13610bad565b610b1b610bea565b8360000151816000600481101515610b2f57fe5b6020020181815250508360200151816001600481101515610b4c57fe5b6020020181815250508260000151816002600481101515610b6957fe5b6020020181815250508260200151816003600481101515610b8657fe5b6020020181815250506040826080836006600019fa1515610ba657600080fd5b5092915050565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a7230582021eaab2949956f42de51fea10cba06158adbc245f9121c871f2847c8fac218290029`

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
