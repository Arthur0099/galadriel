// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipverifier

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IpverifierMetaData contains all meta data concerning the Ipverifier contract.
var IpverifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"gv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"hv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"u\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"l\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"optimizedVerifyIPProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"gv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"hv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"u\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"l\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"verifyIPProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611919806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637b7fc9ad1461003b578063976168631461006b575b600080fd5b61005560048036038101906100509190611655565b61009b565b604051610062919061180c565b60405180910390f35b61008560048036038101906100809190611655565b6100fb565b604051610092919061180c565b60405180910390f35b60006100a56113f7565b60006100b88c8c8c8c8c8c8c8c8c61015b565b8093508192505050806100d0576000925050506100ee565b6100e98c8c8460000151856020015186604001516103e2565b925050505b9998505050505050505050565b60006101056113f7565b60006101188c8c8c8c8c8c8c8c8c61015b565b8093508192505050806101305760009250505061014e565b6101498c8c8460000151856020015186604001516109da565b925050505b9998505050505050505050565b60006101656113f7565b61016d6113f7565b8a518c5114158061018057508551875114155b156101925760008192509250506103d4565b60028c518161019d57fe5b0460028851816101a957fe5b0460020a146101bf5760008192509250506103d4565b865167ffffffffffffffff811180156101d757600080fd5b506040519080825280602002602001820160405280156102065781602001602082028036833780820191505090505b50816040015160000181905250865167ffffffffffffffff8111801561022b57600080fd5b5060405190808252806020026020018201604052801561025a5781602001602082028036833780820191505090505b5081604001516020018190525060005b87518110156102e85787818151811061027f57fe5b6020026020010151826040015160000151828151811061029b57fe5b6020026020010181815250508681815181106102b357fe5b602002602001015182604001516020015182815181106102cf57fe5b602002602001018181525050808060010191505061026a565b5084816040015160400181815250508381604001516060018181525050600061031089610ee2565b905061035b8160405180604001604052808d60006002811061032e57fe5b602002015181526020018d60016002811061034557fe5b6020020151815250610f1d90919063ffffffff16565b82600001819052506103c260405180604001604052808d60006002811061037e57fe5b602002015181526020018d60016002811061039557fe5b60200201518152506103b48b8560000151610f1d90919063ffffffff16565b610fc290919063ffffffff16565b82602001819052506001829350935050505b995099975050505050505050565b60006103ec61142a565b60006002905060005b60028560000151518161040457fe5b0481101561088c5760405180604001604052808660000151836002028151811061042a57fe5b602002602001015181526020018660000151600184600202018151811061044d57fe5b6020026020010151815250836000018190525060405180604001604052808660200151836002028151811061047e57fe5b60200260200101518152602001866020015160018460020201815181106104a157fe5b602002602001015181525083602001819052506000610530866000015183600202815181106104cc57fe5b6020026020010151876000015160018560020201815181106104ea57fe5b60200260200101518860200151856002028151811061050557fe5b60200260200101518960200151600187600202018151811061052357fe5b602002602001015161105a565b9050600061053d8261109e565b905060005b8460028d518161054e57fe5b048161055657fe5b048110156107ef57808560028e518161056b57fe5b048161057357fe5b04018660c001818152505060405180604001604052808d836002028151811061059857fe5b602002602001015181526020018d60018460020201815181106105b757fe5b6020026020010151815250866040018190525060405180604001604052808d60028960c0015102815181106105e857fe5b602002602001015181526020018d600160028a60c0015102018151811061060b57fe5b6020026020010151815250866060018190525061065d610638848860600151610f1d90919063ffffffff16565b61064f848960400151610f1d90919063ffffffff16565b610fc290919063ffffffff16565b86608001819052508560800151600001518c600283028151811061067d57fe5b6020026020010181815250508560800151602001518c60016002840201815181106106a457fe5b60200260200101818152505060405180604001604052808c83600202815181106106ca57fe5b602002602001015181526020018c60018460020201815181106106e957fe5b6020026020010151815250866040018190525060405180604001604052808c60028960c00151028151811061071a57fe5b602002602001015181526020018c600160028a60c0015102018151811061073d57fe5b6020026020010151815250866060018190525061078f61076a838860600151610f1d90919063ffffffff16565b610781858960400151610f1d90919063ffffffff16565b610fc290919063ffffffff16565b8660a001819052508560a00151600001518b60028302815181106107af57fe5b6020026020010181815250508560a00151602001518b60016002840201815181106107d657fe5b6020026020010181815250508080600101915050610542565b5061087588610867610828610815610810868761116890919063ffffffff16565b6111a2565b8960200151610f1d90919063ffffffff16565b610859610846610841888961116890919063ffffffff16565b6111a2565b8a60000151610f1d90919063ffffffff16565b610fc290919063ffffffff16565b610fc290919063ffffffff16565b9750600284029350505080806001019150506103f5565b5060006108b26108ad8660600151876040015161116890919063ffffffff16565b6111a2565b905060405180604001604052808a6000815181106108cc57fe5b602002602001015181526020018a6001815181106108e657fe5b6020026020010151815250836040018190525060405180604001604052808960008151811061091157fe5b602002602001015181526020018960018151811061092b57fe5b6020026020010151815250836060018190525060006109ab61095e87606001518660600151610f1d90919063ffffffff16565b61099d61097c89604001518860400151610f1d90919063ffffffff16565b61098f868d610f1d90919063ffffffff16565b610fc290919063ffffffff16565b610fc290919063ffffffff16565b9050806000015187600001511480156109cb575080602001518760200151145b94505050505095945050505050565b60006109e461142a565b6002836000015151816109f357fe5b0467ffffffffffffffff81118015610a0a57600080fd5b50604051908082528060200260200182016040528015610a395781602001602082028036833780820191505090505b508160e00181905250600283600001515181610a5157fe5b0467ffffffffffffffff81118015610a6857600080fd5b50604051908082528060200260200182016040528015610a975781602001602082028036833780820191505090505b5081610100018190525060005b600284600001515181610ab357fe5b04811015610caf57604051806040016040528085600001518360020281518110610ad957fe5b6020026020010151815260200185600001516001846002020181518110610afc57fe5b60200260200101518152508260000181905250604051806040016040528085602001518360020281518110610b2d57fe5b6020026020010151815260200185602001516001846002020181518110610b5057fe5b602002602001015181525082602001819052506000610bdf85600001518360020281518110610b7b57fe5b602002602001015186600001516001856002020181518110610b9957fe5b602002602001015187602001518560020281518110610bb457fe5b602002602001015188602001516001876002020181518110610bd257fe5b602002602001015161105a565b90506000610bec8261109e565b9050818460e001518481518110610bff57fe5b602002602001018181525050808461010001518481518110610c1d57fe5b602002602001018181525050610c9e610c55610c42838461116890919063ffffffff16565b8660200151610f1d90919063ffffffff16565b610c90610c81610c6e868761116890919063ffffffff16565b8860000151610f1d90919063ffffffff16565b8a610fc290919063ffffffff16565b610fc290919063ffffffff16565b965050508080600101915050610aa4565b5060006002885181610cbd57fe5b0467ffffffffffffffff81118015610cd457600080fd5b50604051908082528060200260200182016040528015610d035781602001602082028036833780820191505090505b50905060005b6002895181610d1457fe5b04811015610e1c5760005b600286600001515181610d2e57fe5b04811015610e0e576000610d52838360028a600001515181610d4c57fe5b046111da565b15610d76578460e001518281518110610d6757fe5b60200260200101519050610d92565b8461010001518281518110610d8757fe5b602002602001015190505b6000821415610db95780848481518110610da857fe5b602002602001018181525050610e00565b610de7610de282868681518110610dcc57fe5b602002602001015161116890919063ffffffff16565b6111a2565b848481518110610df357fe5b6020026020010181815250505b508080600101915050610d1f565b508080600101915050610d09565b50610e2561149a565b610eb4610e55610e468760600151886040015161116890919063ffffffff16565b89610f1d90919063ffffffff16565b610ea6610e788860600151610e6a8d88611228565b610f1d90919063ffffffff16565b610e988960400151610e8a8f896112dd565b610f1d90919063ffffffff16565b610fc290919063ffffffff16565b610fc290919063ffffffff16565b905085600001518160000151148015610ed4575085602001518160200151145b935050505095945050505050565b6000610f1682604051602001610ef891906117a3565b6040516020818303038152906040528051906020012060001c6111a2565b9050919050565b610f2561149a565b6001821415610f3657829050610fbc565b6002821415610f5057610f498384610fc2565b9050610fbc565b610f586114b4565b836000015181600060038110610f6a57fe5b602002018181525050836020015181600160038110610f8557fe5b6020020181815250508281600260038110610f9c57fe5b6020020181815250506040826060836007600019fa610fba57600080fd5b505b92915050565b610fca61149a565b610fd26114d6565b836000015181600060048110610fe457fe5b602002018181525050836020015181600160048110610fff57fe5b60200201818152505082600001518160026004811061101a57fe5b60200201818152505082602001518160036004811061103557fe5b6020020181815250506040826080836006600019fa61105357600080fd5b5092915050565b60006110948585858560405160200161107694939291906117be565b6040516020818303038152906040528051906020012060001c6111a2565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050600083905060008114156110dd57600092505050611163565b818111156110f2578181816110ee57fe5b0690505b600080600190506000849050600084905060005b6000821461113b5781838161111757fe5b04905083848202860383848402860380955081965082975083985050505050611106565b600085121561115857846000038703975050505050505050611163565b849750505050505050505b919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808061119657fe5b83850991505092915050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508083816111d157fe5b06915050919050565b600080600183036001901b905060005b8481101561120657600182901c915080806001019150506111ea565b5060008186161461121b576001915050611221565b60009150505b9392505050565b61123061149a565b6000825167ffffffffffffffff8111801561124a57600080fd5b506040519080825280602002602001820160405280156112795781602001602082028036833780820191505090505b50905060005b83518110156112c9576112a484828151811061129757fe5b602002602001015161109e565b8282815181106112b057fe5b602002602001018181525050808060010191505061127f565b506112d484826112dd565b91505092915050565b6112e561149a565b6112ed61149a565b611350836000815181106112fd57fe5b602002602001015160405180604001604052808760008151811061131d57fe5b602002602001015181526020018760018151811061133757fe5b6020026020010151815250610f1d90919063ffffffff16565b90506000600190505b83518110156113ec576113dd6113ce85838151811061137457fe5b6020026020010151604051806040016040528089866002028151811061139657fe5b602002602001015181526020018960018760020201815181106113b557fe5b6020026020010151815250610f1d90919063ffffffff16565b83610fc290919063ffffffff16565b91508080600101915050611359565b508091505092915050565b604051806060016040528061140a61149a565b815260200161141761149a565b81526020016114246114f8565b81525090565b60405180610120016040528061143e61149a565b815260200161144b61149a565b815260200161145861149a565b815260200161146561149a565b815260200161147261149a565b815260200161147f61149a565b81526020016000815260200160608152602001606081525090565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028036833780820191505090505090565b6040518060800160405280600490602082028036833780820191505090505090565b6040518060800160405280606081526020016060815260200160008152602001600081525090565b600061153361152e84611858565b611827565b9050808285602086028201111561154957600080fd5b60005b85811015611579578161155f8882611640565b84526020840193506020830192505060018101905061154c565b5050509392505050565b60006115966115918461187e565b611827565b905080838252602082019050828560208602820111156115b557600080fd5b60005b858110156115e557816115cb8882611640565b8452602084019350602083019250506001810190506115b8565b5050509392505050565b600082601f83011261160057600080fd5b600261160d848285611520565b91505092915050565b600082601f83011261162757600080fd5b8135611637848260208601611583565b91505092915050565b60008135905061164f816118cc565b92915050565b60008060008060008060008060006101608a8c03121561167457600080fd5b60008a013567ffffffffffffffff81111561168e57600080fd5b61169a8c828d01611616565b99505060208a013567ffffffffffffffff8111156116b757600080fd5b6116c38c828d01611616565b98505060406116d48c828d016115ef565b97505060806116e58c828d016115ef565b96505060c06116f68c828d01611640565b95505060e08a013567ffffffffffffffff81111561171357600080fd5b61171f8c828d01611616565b9450506101008a013567ffffffffffffffff81111561173d57600080fd5b6117498c828d01611616565b93505061012061175b8c828d01611640565b92505061014061176d8c828d01611640565b9150509295985092959850929598565b611786816118aa565b82525050565b61179d611798826118b6565b6118c0565b82525050565b60006117af828461178c565b60208201915081905092915050565b60006117ca828761178c565b6020820191506117da828661178c565b6020820191506117ea828561178c565b6020820191506117fa828461178c565b60208201915081905095945050505050565b6000602082019050611821600083018461177d565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561184e5761184d6118ca565b5b8060405250919050565b600067ffffffffffffffff821115611873576118726118ca565b5b602082029050919050565b600067ffffffffffffffff821115611899576118986118ca565b5b602082029050602081019050919050565b60008115159050919050565b6000819050919050565b6000819050919050565bfe5b6118d5816118b6565b81146118e057600080fd5b5056fea26469706673582212208ee6d23a20edbe8a56bfb13931d05c15abcbd78b94581bbe400b338a160cb8ec64736f6c63430007060033",
}

// IpverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IpverifierMetaData.ABI instead.
var IpverifierABI = IpverifierMetaData.ABI

// IpverifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IpverifierMetaData.Bin instead.
var IpverifierBin = IpverifierMetaData.Bin

// DeployIpverifier deploys a new Ethereum contract, binding an instance of Ipverifier to it.
func DeployIpverifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ipverifier, error) {
	parsed, err := IpverifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IpverifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ipverifier{IpverifierCaller: IpverifierCaller{contract: contract}, IpverifierTransactor: IpverifierTransactor{contract: contract}, IpverifierFilterer: IpverifierFilterer{contract: contract}}, nil
}

// Ipverifier is an auto generated Go binding around an Ethereum contract.
type Ipverifier struct {
	IpverifierCaller     // Read-only binding to the contract
	IpverifierTransactor // Write-only binding to the contract
	IpverifierFilterer   // Log filterer for contract events
}

// IpverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpverifierSession struct {
	Contract     *Ipverifier       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpverifierCallerSession struct {
	Contract *IpverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IpverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpverifierTransactorSession struct {
	Contract     *IpverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IpverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpverifierRaw struct {
	Contract *Ipverifier // Generic contract binding to access the raw methods on
}

// IpverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpverifierCallerRaw struct {
	Contract *IpverifierCaller // Generic read-only contract binding to access the raw methods on
}

// IpverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpverifierTransactorRaw struct {
	Contract *IpverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpverifier creates a new instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifier(address common.Address, backend bind.ContractBackend) (*Ipverifier, error) {
	contract, err := bindIpverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipverifier{IpverifierCaller: IpverifierCaller{contract: contract}, IpverifierTransactor: IpverifierTransactor{contract: contract}, IpverifierFilterer: IpverifierFilterer{contract: contract}}, nil
}

// NewIpverifierCaller creates a new read-only instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifierCaller(address common.Address, caller bind.ContractCaller) (*IpverifierCaller, error) {
	contract, err := bindIpverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpverifierCaller{contract: contract}, nil
}

// NewIpverifierTransactor creates a new write-only instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IpverifierTransactor, error) {
	contract, err := bindIpverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpverifierTransactor{contract: contract}, nil
}

// NewIpverifierFilterer creates a new log filterer instance of Ipverifier, bound to a specific deployed contract.
func NewIpverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IpverifierFilterer, error) {
	contract, err := bindIpverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpverifierFilterer{contract: contract}, nil
}

// bindIpverifier binds a generic wrapper to an already deployed contract.
func bindIpverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IpverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipverifier *IpverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipverifier.Contract.IpverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipverifier *IpverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipverifier.Contract.IpverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipverifier *IpverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipverifier.Contract.IpverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipverifier *IpverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipverifier *IpverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipverifier *IpverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipverifier.Contract.contract.Transact(opts, method, params...)
}

// OptimizedVerifyIPProof is a free data retrieval call binding the contract method 0x97616863.
//
// Solidity: function optimizedVerifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) view returns(bool)
func (_Ipverifier *IpverifierCaller) OptimizedVerifyIPProof(opts *bind.CallOpts, gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _Ipverifier.contract.Call(opts, &out, "optimizedVerifyIPProof", gv, hv, p, u, c, l, r, a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OptimizedVerifyIPProof is a free data retrieval call binding the contract method 0x97616863.
//
// Solidity: function optimizedVerifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) view returns(bool)
func (_Ipverifier *IpverifierSession) OptimizedVerifyIPProof(gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.OptimizedVerifyIPProof(&_Ipverifier.CallOpts, gv, hv, p, u, c, l, r, a, b)
}

// OptimizedVerifyIPProof is a free data retrieval call binding the contract method 0x97616863.
//
// Solidity: function optimizedVerifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) view returns(bool)
func (_Ipverifier *IpverifierCallerSession) OptimizedVerifyIPProof(gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.OptimizedVerifyIPProof(&_Ipverifier.CallOpts, gv, hv, p, u, c, l, r, a, b)
}

// VerifyIPProof is a free data retrieval call binding the contract method 0x7b7fc9ad.
//
// Solidity: function verifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) view returns(bool)
func (_Ipverifier *IpverifierCaller) VerifyIPProof(opts *bind.CallOpts, gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _Ipverifier.contract.Call(opts, &out, "verifyIPProof", gv, hv, p, u, c, l, r, a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyIPProof is a free data retrieval call binding the contract method 0x7b7fc9ad.
//
// Solidity: function verifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) view returns(bool)
func (_Ipverifier *IpverifierSession) VerifyIPProof(gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProof(&_Ipverifier.CallOpts, gv, hv, p, u, c, l, r, a, b)
}

// VerifyIPProof is a free data retrieval call binding the contract method 0x7b7fc9ad.
//
// Solidity: function verifyIPProof(uint256[] gv, uint256[] hv, uint256[2] p, uint256[2] u, uint256 c, uint256[] l, uint256[] r, uint256 a, uint256 b) view returns(bool)
func (_Ipverifier *IpverifierCallerSession) VerifyIPProof(gv []*big.Int, hv []*big.Int, p [2]*big.Int, u [2]*big.Int, c *big.Int, l []*big.Int, r []*big.Int, a *big.Int, b *big.Int) (bool, error) {
	return _Ipverifier.Contract.VerifyIPProof(&_Ipverifier.CallOpts, gv, hv, p, u, c, l, r, a, b)
}
