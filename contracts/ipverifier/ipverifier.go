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
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"gv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"hv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"u\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"l\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"optimizedVerifyIPProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"gv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"hv\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"u\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"l\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"verifyIPProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506118c7806100206000396000f3fe608060405234801561001057600080fd5b5060043610610053576000357c0100000000000000000000000000000000000000000000000000000000900480637b7fc9ad146100585780639761686314610088575b600080fd5b610072600480360361006d9190810190611604565b6100b8565b60405161007f91906117bb565b60405180910390f35b6100a2600480360361009d9190810190611604565b610118565b6040516100af91906117bb565b60405180910390f35b60006100c26113b5565b60006100d58c8c8c8c8c8c8c8c8c610178565b8093508192505050806100ed5760009250505061010b565b6101068c8c8460000151856020015186604001516103d4565b925050505b9998505050505050505050565b60006101226113b5565b60006101358c8c8c8c8c8c8c8c8c610178565b80935081925050508061014d5760009250505061016b565b6101668c8c8460000151856020015186604001516109d8565b925050505b9998505050505050505050565b60006101826113b5565b61018a6113b5565b8a518c5114158061019d57508551875114155b156101af5760008192509250506103c6565b60028c51816101ba57fe5b0460028851816101c657fe5b0460020a146101dc5760008192509250506103c6565b865160405190808252806020026020018201604052801561020c5781602001602082028038833980820191505090505b5081604001516000018190525086516040519080825280602002602001820160405280156102495781602001602082028038833980820191505090505b5081604001516020018190525060008090505b87518110156102da5787818151811061027157fe5b6020026020010151826040015160000151828151811061028d57fe5b6020026020010181815250508681815181106102a557fe5b602002602001015182604001516020015182815181106102c157fe5b602002602001018181525050808060010191505061025c565b5084816040015160400181815250508381604001516060018181525050600061030289610ea4565b905061034d8160405180604001604052808d60006002811061032057fe5b602002015181526020018d60016002811061033757fe5b6020020151815250610ee090919063ffffffff16565b82600001819052506103b460405180604001604052808d60006002811061037057fe5b602002015181526020018d60016002811061038757fe5b60200201518152506103a68b8560000151610ee090919063ffffffff16565b610f8590919063ffffffff16565b82602001819052506001829350935050505b995099975050505050505050565b60006103de6113e8565b60006002905060008090505b6002856000015151816103f957fe5b048110156108845760405180604001604052808660000151836002028151811061041f57fe5b602002602001015181526020018660000151600184600202018151811061044257fe5b6020026020010151815250836000018190525060405180604001604052808660200151836002028151811061047357fe5b602002602001015181526020018660200151600184600202018151811061049657fe5b602002602001015181525083602001819052506000610525866000015183600202815181106104c157fe5b6020026020010151876000015160018560020201815181106104df57fe5b6020026020010151886020015185600202815181106104fa57fe5b60200260200101518960200151600187600202018151811061051857fe5b602002602001015161101d565b9050600061053282611062565b905060008090505b8460028d518161054657fe5b048161054e57fe5b048110156107e757808560028e518161056357fe5b048161056b57fe5b04018660c001818152505060405180604001604052808d836002028151811061059057fe5b602002602001015181526020018d60018460020201815181106105af57fe5b6020026020010151815250866040018190525060405180604001604052808d60028960c0015102815181106105e057fe5b602002602001015181526020018d600160028a60c0015102018151811061060357fe5b60200260200101518152508660600181905250610655610630848860600151610ee090919063ffffffff16565b610647848960400151610ee090919063ffffffff16565b610f8590919063ffffffff16565b86608001819052508560800151600001518c600283028151811061067557fe5b6020026020010181815250508560800151602001518c600160028402018151811061069c57fe5b60200260200101818152505060405180604001604052808c83600202815181106106c257fe5b602002602001015181526020018c60018460020201815181106106e157fe5b6020026020010151815250866040018190525060405180604001604052808c60028960c00151028151811061071257fe5b602002602001015181526020018c600160028a60c0015102018151811061073557fe5b60200260200101518152508660600181905250610787610762838860600151610ee090919063ffffffff16565b610779858960400151610ee090919063ffffffff16565b610f8590919063ffffffff16565b8660a001819052508560a00151600001518b60028302815181106107a757fe5b6020026020010181815250508560a00151602001518b60016002840201815181106107ce57fe5b602002602001018181525050808060010191505061053a565b5061086d8861085f61082061080d610808868761112c90919063ffffffff16565b611166565b8960200151610ee090919063ffffffff16565b61085161083e610839888961112c90919063ffffffff16565b611166565b8a60000151610ee090919063ffffffff16565b610f8590919063ffffffff16565b610f8590919063ffffffff16565b9750600284029350505080806001019150506103ea565b5060006108aa6108a58660600151876040015161112c90919063ffffffff16565b611166565b905060405180604001604052808a6000815181106108c457fe5b602002602001015181526020018a6001815181106108de57fe5b6020026020010151815250836040018190525060405180604001604052808960008151811061090957fe5b602002602001015181526020018960018151811061092357fe5b6020026020010151815250836060018190525061093e611458565b6109a961095c87606001518660600151610ee090919063ffffffff16565b61099b61097a89604001518860400151610ee090919063ffffffff16565b61098d868d610ee090919063ffffffff16565b610f8590919063ffffffff16565b610f8590919063ffffffff16565b9050806000015187600001511480156109c9575080602001518760200151145b94505050505095945050505050565b60006109e26113e8565b6002836000015151816109f157fe5b04604051908082528060200260200182016040528015610a205781602001602082028038833980820191505090505b508160e00181905250600283600001515181610a3857fe5b04604051908082528060200260200182016040528015610a675781602001602082028038833980820191505090505b5081610100018190525060008090505b600284600001515181610a8657fe5b04811015610c8257604051806040016040528085600001518360020281518110610aac57fe5b6020026020010151815260200185600001516001846002020181518110610acf57fe5b60200260200101518152508260000181905250604051806040016040528085602001518360020281518110610b0057fe5b6020026020010151815260200185602001516001846002020181518110610b2357fe5b602002602001015181525082602001819052506000610bb285600001518360020281518110610b4e57fe5b602002602001015186600001516001856002020181518110610b6c57fe5b602002602001015187602001518560020281518110610b8757fe5b602002602001015188602001516001876002020181518110610ba557fe5b602002602001015161101d565b90506000610bbf82611062565b9050818460e001518481518110610bd257fe5b602002602001018181525050808461010001518481518110610bf057fe5b602002602001018181525050610c71610c28610c15838461112c90919063ffffffff16565b8660200151610ee090919063ffffffff16565b610c63610c54610c41868761112c90919063ffffffff16565b8860000151610ee090919063ffffffff16565b8a610f8590919063ffffffff16565b610f8590919063ffffffff16565b965050508080600101915050610a77565b5060606002885181610c9057fe5b04604051908082528060200260200182016040528015610cbf5781602001602082028038833980820191505090505b50905060008090505b6002895181610cd357fe5b04811015610dde5760008090505b600286600001515181610cf057fe5b04811015610dd0576000610d14838360028a600001515181610d0e57fe5b0461119e565b15610d38578460e001518281518110610d2957fe5b60200260200101519050610d54565b8461010001518281518110610d4957fe5b602002602001015190505b6000821415610d7b5780848481518110610d6a57fe5b602002602001018181525050610dc2565b610da9610da482868681518110610d8e57fe5b602002602001015161112c90919063ffffffff16565b611166565b848481518110610db557fe5b6020026020010181815250505b508080600101915050610ce1565b508080600101915050610cc8565b50610de7611458565b610e76610e17610e088760600151886040015161112c90919063ffffffff16565b89610ee090919063ffffffff16565b610e68610e3a8860600151610e2c8d886111fa565b610ee090919063ffffffff16565b610e5a8960400151610e4c8f8961129b565b610ee090919063ffffffff16565b610f8590919063ffffffff16565b610f8590919063ffffffff16565b905085600001518160000151148015610e96575085602001518160200151145b935050505095945050505050565b6000610ed982604051602001610eba9190611752565b6040516020818303038152906040528051906020012060019004611166565b9050919050565b610ee8611458565b6001821415610ef957829050610f7f565b6002821415610f1357610f0c8384610f85565b9050610f7f565b610f1b611472565b836000015181600060038110610f2d57fe5b602002018181525050836020015181600160038110610f4857fe5b6020020181815250508281600260038110610f5f57fe5b6020020181815250506040826060836007600019fa610f7d57600080fd5b505b92915050565b610f8d611458565b610f95611494565b836000015181600060048110610fa757fe5b602002018181525050836020015181600160048110610fc257fe5b602002018181525050826000015181600260048110610fdd57fe5b602002018181525050826020015181600360048110610ff857fe5b6020020181815250506040826080836006600019fa61101657600080fd5b5092915050565b600061105885858585604051602001611039949392919061176d565b6040516020818303038152906040528051906020012060019004611166565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050600083905060008114156110a157600092505050611127565b818111156110b6578181816110b257fe5b0690505b600080600190506000849050600084905060005b600082146110ff578183816110db57fe5b049050838482028603838484028603809550819650829750839850505050506110ca565b600085121561111c57846000038703975050505050505050611127565b849750505050505050505b919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808061115a57fe5b83850991505092915050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161119557fe5b06915050919050565b6000806001830360019060020a02905060008090505b848110156111d857600182908060020a8204915050915080806001019150506111b4565b506000818616146111ed5760019150506111f3565b60009150505b9392505050565b611202611458565b606082516040519080825280602002602001820160405280156112345781602001602082028038833980820191505090505b50905060008090505b83518110156112875761126284828151811061125557fe5b6020026020010151611062565b82828151811061126e57fe5b602002602001018181525050808060010191505061123d565b50611292848261129b565b91505092915050565b6112a3611458565b6112ab611458565b61130e836000815181106112bb57fe5b60200260200101516040518060400160405280876000815181106112db57fe5b60200260200101518152602001876001815181106112f557fe5b6020026020010151815250610ee090919063ffffffff16565b90506000600190505b83518110156113aa5761139b61138c85838151811061133257fe5b6020026020010151604051806040016040528089866002028151811061135457fe5b6020026020010151815260200189600187600202018151811061137357fe5b6020026020010151815250610ee090919063ffffffff16565b83610f8590919063ffffffff16565b91508080600101915050611317565b508091505092915050565b60405180606001604052806113c86114b6565b81526020016113d56114b6565b81526020016113e26114d0565b81525090565b6040518061012001604052806113fc6114b6565b81526020016114096114b6565b81526020016114166114b6565b81526020016114236114b6565b81526020016114306114b6565b815260200161143d6114b6565b81526020016000815260200160608152602001606081525090565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b604051806040016040528060008152602001600081525090565b6040518060800160405280606081526020016060815260200160008152602001600081525090565b600082601f83011261150957600080fd5b600261151c61151782611803565b6117d6565b9150818385602084028201111561153257600080fd5b60005b83811015611562578161154888826115ef565b845260208401935060208301925050600181019050611535565b5050505092915050565b600082601f83011261157d57600080fd5b813561159061158b82611825565b6117d6565b915081818352602084019350602081019050838560208402820111156115b557600080fd5b60005b838110156115e557816115cb88826115ef565b8452602084019350602083019250506001810190506115b8565b5050505092915050565b6000813590506115fe8161186d565b92915050565b60008060008060008060008060006101608a8c03121561162357600080fd5b60008a013567ffffffffffffffff81111561163d57600080fd5b6116498c828d0161156c565b99505060208a013567ffffffffffffffff81111561166657600080fd5b6116728c828d0161156c565b98505060406116838c828d016114f8565b97505060806116948c828d016114f8565b96505060c06116a58c828d016115ef565b95505060e08a013567ffffffffffffffff8111156116c257600080fd5b6116ce8c828d0161156c565b9450506101008a013567ffffffffffffffff8111156116ec57600080fd5b6116f88c828d0161156c565b93505061012061170a8c828d016115ef565b92505061014061171c8c828d016115ef565b9150509295985092959850929598565b6117358161184d565b82525050565b61174c61174782611859565b611863565b82525050565b600061175e828461173b565b60208201915081905092915050565b6000611779828761173b565b602082019150611789828661173b565b602082019150611799828561173b565b6020820191506117a9828461173b565b60208201915081905095945050505050565b60006020820190506117d0600083018461172c565b92915050565b6000604051905081810181811067ffffffffffffffff821117156117f957600080fd5b8060405250919050565b600067ffffffffffffffff82111561181a57600080fd5b602082029050919050565b600067ffffffffffffffff82111561183c57600080fd5b602082029050602081019050919050565b60008115159050919050565b6000819050919050565b6000819050919050565b61187681611859565b811461188157600080fd5b5056fea365627a7a72315820da18401f7c04d2309a5dc0b43bb5d865ae7446a9fabefc0190e4526c0d66c2de6c6578706572696d656e74616cf564736f6c63430005100040",
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
