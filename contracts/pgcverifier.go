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

// PgcverifierABI is the input ABI used to generate the binding from.
const PgcverifierABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"aggRangeProofVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dleSigmaVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rangeProofVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sigmaVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"params_\",\"type\":\"address\"},{\"name\":\"dleSigmaVerifier_\",\"type\":\"address\"},{\"name\":\"rangeProofVerifier_\",\"type\":\"address\"},{\"name\":\"aggVerifier_\",\"type\":\"address\"},{\"name\":\"sigmaVerifier_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[36]\"},{\"name\":\"scalar\",\"type\":\"uint256[10]\"},{\"name\":\"l\",\"type\":\"uint256[12]\"},{\"name\":\"r\",\"type\":\"uint256[12]\"},{\"name\":\"ub\",\"type\":\"uint256[4]\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"verifyAggTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"publicKey\",\"type\":\"uint256[2]\"},{\"name\":\"proof\",\"type\":\"uint256[4]\"},{\"name\":\"z\",\"type\":\"uint256\"},{\"name\":\"ub\",\"type\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[]\"}],\"name\":\"verifyBurn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PgcverifierBin is the compiled bytecode used for deploying new contracts.
const PgcverifierBin = `0x60806040523480156200001157600080fd5b5060405160a08062002720833981018060405260a08110156200003357600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919050505084600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001bb62000515565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200023f57600080fd5b505afa15801562000254573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200027a57600080fd5b810190809190505090506200028e62000515565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200031257600080fd5b505afa15801562000327573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200034d57600080fd5b810190809190505090508160006002811015156200036757fe5b602002015160008001819055508160016002811015156200038457fe5b6020020151600060010181905550806000600281101515620003a257fe5b6020020151600260000181905550806001600281101515620003c057fe5b6020020151600260010181905550600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da8972246040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200045357600080fd5b505afa15801562000468573d6000803e3d6000fd5b505050506040513d60208110156200047f57600080fd5b8101908080519060200190929190505050602014151562000508576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f62697473697a65206e6f7420657175616c00000000000000000000000000000081525060200191505060405180910390fd5b5050505050505062000537565b6040805190810160405280600290602082028038833980820191505090505090565b6121d980620005476000396000f3fe6080604052600436106100ba576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806326f9031b146100bf578063293ec7b4146102745780632e52d606146102cb578063361eb474146102f65780633a4f69991461034d5780633e8d37641461037857806361e5af5e146103a357806381b7bc73146103fa578063a89e0e4414610451578063b8c9d365146105ee578063cff0ab9614610620578063e2179b8e14610677575b600080fd5b3480156100cb57600080fd5b5061025a60048036036101a08110156100e357600080fd5b810190808035906020019092919080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001906401000000008111156101d757600080fd5b8201836020820111156101e957600080fd5b8035906020019184602083028401116401000000008311171561020b57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506106a9565b604051808215151515815260200191505060405180910390f35b34801561028057600080fd5b50610289610802565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102d757600080fd5b506102e0610828565b6040518082815260200191505060405180910390f35b34801561030257600080fd5b5061030b61082d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561035957600080fd5b50610362610853565b6040518082815260200191505060405180910390f35b34801561038457600080fd5b5061038d61085b565b6040518082815260200191505060405180910390f35b3480156103af57600080fd5b506103b8610860565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561040657600080fd5b5061040f610886565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561045d57600080fd5b506105d4600480360361096081101561047557600080fd5b810190808061048001906024806020026040519081016040528092919082602460200280828437600081840152601f19601f8201169050808301925050505050509192919290806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291905050506108ac565b604051808215151515815260200191505060405180910390f35b3480156105fa57600080fd5b50610603611459565b604051808381526020018281526020019250505060405180910390f35b34801561062c57600080fd5b5061063561146b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561068357600080fd5b5061068c611491565b604051808381526020018281526020019250505060405180910390f35b60006106b3611e51565b8360006004811015156106c257fe5b6020020151816000015160000181815250508360016004811015156106e357fe5b60200201518160000151602001818152505083600260048110151561070457fe5b60200201518160200151600001818152505083600360048110151561072557fe5b60200201518160200151602001818152505061077f888260408051908101604052808b600060028110151561075657fe5b602002015181526020018b600160028110151561076f57fe5b60200201518152508989886114a3565b15156107f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f646c65207369676d6120766572696679206661696c656400000000000000000081525060200191505060405180910390fd5b60019150509695505050505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600581565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602060020a81565b602081565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006108b6611e51565b8360006004811015156108c557fe5b6020020151816000015160000181815250508360016004811015156108e657fe5b60200201518160000151602001818152505083600260048110151561090757fe5b60200201518160200151600001818152505083600360048110151561092857fe5b602002015181602001516020018181525050610942611e7f565b60008160000181815250505b6010816000015110156109a65788816000015160248110151561096d57fe5b60200201518160400151826000015160108110151561098857fe5b6020020181815250508060000180518091906001018152505061094e565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e1efe76682604001518a6000600a811015156109f857fe5b60200201518b6001600a81101515610a0c57fe5b60200201516040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084601060200280838360005b83811015610a67578082015181840152602081019050610a4c565b505050509050018381526020018281526020019350505050602060405180830381600087803b158015610a9957600080fd5b505af1158015610aad573d6000803e3d6000fd5b505050506040513d6020811015610ac357600080fd5b81019080805190602001909291905050501515610b48576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f70746520657175616c2070726f6f6620696e76616c696400000000000000000081525060200191505060405180910390fd5b60408051908101604052808a6004602481101515610b6257fe5b602002015181526020018a6005602481101515610b7b57fe5b602002015181525081608001516000018190525060408051908101604052808a6008602481101515610ba957fe5b602002015181526020018a6009602481101515610bc257fe5b6020020151815250816080015160200181905250610bfd610bea82608001516000015161183b565b83600001516118cb90919063ffffffff16565b81610100015160000181905250610c31610c1e82608001516020015161183b565b83602001516118cb90919063ffffffff16565b8161010001516020018190525060408051908101604052808a6010602481101515610c5857fe5b602002015181526020018a6011602481101515610c7157fe5b60200201518152508161012001516000018190525060408051908101604052808a6012602481101515610ca057fe5b602002015181526020018a6013602481101515610cb957fe5b60200201518152508161012001516020018190525060408051908101604052808a6018602481101515610ce857fe5b602002015181526020018a6019602481101515610d0157fe5b60200201518152508161014001516000600281101515610d1d57fe5b602002018190525060408051908101604052808a601a602481101515610d3f57fe5b602002015181526020018a601b602481101515610d5857fe5b60200201518152508161014001516001600281101515610d7457fe5b6020020181905250610dd58161010001518261012001518361014001518c6000602481101515610da057fe5b60200201518d6001602481101515610db457fe5b60200201518d6004600a81101515610dc857fe5b602002015160018b61196d565b1515610e49576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f646c65207369676d612070726f6f66206661696c65640000000000000000000081525060200191505060405180910390fd5b886000602481101515610e5857fe5b602002015181606001516000600a81101515610e7057fe5b602002018181525050886001602481101515610e8857fe5b602002015181606001516001600a81101515610ea057fe5b60200201818152505060108160000181815250505b601881600001511015610f1357888160000151602481101515610ed457fe5b602002015181606001516010836000015160020103600a81101515610ef557fe5b60200201818152505080600001805180919060010181525050610eb5565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663991395ca82606001518a6002600a81101515610f6557fe5b60200201518b6003600a81101515610f7957fe5b60200201516040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600a60200280838360005b83811015610fd4578082015181840152602081019050610fb9565b505050509050018381526020018281526020019350505050602060405180830381600087803b15801561100657600080fd5b505af115801561101a573d6000803e3d6000fd5b505050506040513d602081101561103057600080fd5b810190808051906020019092919050505015156110b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f63742076616c69642070726f6f6620696e76616c69640000000000000000000081525060200191505060405180910390fd5b60008160000181815250505b60088160000151101561111d57888160000151601c016024811015156110e357fe5b60200201518161022001518260000151600c811015156110ff57fe5b602002018181525050806000018051809190600101815250506110c1565b88600860248110151561112c57fe5b60200201518161022001516008600c8110151561114557fe5b60200201818152505088600960248110151561115d57fe5b60200201518161022001516009600c8110151561117657fe5b60200201818152505088601260248110151561118e57fe5b6020020151816102200151600a600c811015156111a757fe5b6020020181815250508860136024811015156111bf57fe5b6020020151816102200151600b600c811015156111d857fe5b60200201818152505060008160000181815250505b60058160000151101561124957876005826000015101600a8110151561120f57fe5b6020020151816102400151826000015160058110151561122b57fe5b602002018181525050806000018051809190600101815250506111ed565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d3f6b278261022001518361024001518a8a6040518563ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018085600c60200280838360005b838110156112eb5780820151818401526020810190506112d0565b5050505090500184600560200280838360005b838110156113195780820151818401526020810190506112fe565b5050505090500183600c60200280838360005b8381101561134757808201518184015260208101905061132c565b5050505090500182600c60200280838360005b8381101561137557808201518184015260208101905061135a565b5050505090500194505050505060206040518083038186803b15801561139a57600080fd5b505afa1580156113ae573d6000803e3d6000fd5b505050506040513d60208110156113c457600080fd5b81019080805190602001909291905050501515611449576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f616767726174652072616e67652070726f6f6620696e76616c6964000000000081525060200191505060405180910390fd5b6001925050509695505050505050565b60008060000154908060010154905082565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028060000154908060010154905082565b60006114ad611fdd565b6114fe6114eb6114e68a600260408051908101604052908160008201548152602001600182015481525050611da490919063ffffffff16565b61183b565b88602001516118cb90919063ffffffff16565b9050611508611e7f565b85600060048110151561151757fe5b60200201518161016001516000600c8110151561153057fe5b60200201818152505085600160048110151561154857fe5b60200201518161016001516001600c8110151561156157fe5b60200201818152505085600260048110151561157957fe5b60200201518161016001516002600c8110151561159257fe5b6020020181815250508560036004811015156115aa57fe5b60200201518161016001516003600c811015156115c357fe5b60200201818152505081600001518161016001516004600c811015156115e557fe5b60200201818152505081602001518161016001516005600c8110151561160757fe5b6020020181815250508760000151600001518161016001516006600c8110151561162d57fe5b6020020181815250508760000151602001518161016001516007600c8110151561165357fe5b60200201818152505060008001548161016001516008600c8110151561167557fe5b6020020181815250506000600101548161016001516009600c8110151561169857fe5b6020020181815250508660000151816101600151600a600c811015156116ba57fe5b6020020181815250508660200151816101600151600b600c811015156116dc57fe5b602002018181525050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e26d0b882610160015187876040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600c60200280838360005b83811015611781578082015181840152602081019050611766565b5050505090500183815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156117cd5780820151818401526020810190506117b2565b5050505090500194505050505060206040518083038186803b1580156117f257600080fd5b505afa158015611806573d6000803e3d6000fd5b505050506040513d602081101561181c57600080fd5b8101908080519060200190929190505050925050509695505050505050565b611843611fdd565b6000826000015114801561185b575060008260200151145b1561187e57604080519081016040528060008152602001600081525090506118c6565b60007f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd479050604080519081016040528084600001518152602001846020015183038152509150505b919050565b6118d3611fdd565b6118db611ff7565b83600001518160006004811015156118ef57fe5b602002018181525050836020015181600160048110151561190c57fe5b602002018181525050826000015181600260048110151561192957fe5b602002018181525050826020015181600360048110151561194657fe5b6020020181815250506040826080836006600019fa151561196657600080fd5b5092915050565b6000611977611fdd565b61199a6119878b6020015161183b565b8a602001516118cb90919063ffffffff16565b90506119a4611fdd565b6119c76119b48c6000015161183b565b8b600001516118cb90919063ffffffff16565b90506119d161201a565b8960006002811015156119e057fe5b602002015160000151816000600c811015156119f857fe5b602002018181525050896000600281101515611a1057fe5b602002015160200151816001600c81101515611a2857fe5b602002018181525050896001600281101515611a4057fe5b602002015160000151816002600c81101515611a5857fe5b602002018181525050896001600281101515611a7057fe5b602002015160200151816003600c81101515611a8857fe5b6020020181815250508260000151816004600c81101515611aa557fe5b6020020181815250508260200151816005600c81101515611ac257fe5b6020020181815250508160000151816006600c81101515611adf57fe5b6020020181815250508160200151816007600c81101515611afc57fe5b6020020181815250506000800154816008600c81101515611b1957fe5b602002018181525050600060010154816009600c81101515611b3757fe5b6020020181815250508881600a600c81101515611b5057fe5b6020020181815250508781600b600c81101515611b6957fe5b6020020181815250506000861415611c7f57600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639751cb1382896040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600c60200280838360005b83811015611c11578082015181840152602081019050611bf6565b505050509050018281526020019250505060206040518083038186803b158015611c3a57600080fd5b505afa158015611c4e573d6000803e3d6000fd5b505050506040513d6020811015611c6457600080fd5b81019080805190602001909291905050509350505050611d98565b6001861415611d9457600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630581c0408289886040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600c60200280838360005b83811015611d1f578082015181840152602081019050611d04565b50505050905001838152602001828152602001935050505060206040518083038186803b158015611d4f57600080fd5b505afa158015611d63573d6000803e3d6000fd5b505050506040513d6020811015611d7957600080fd5b81019080805190602001909291905050509350505050611d98565b5050505b98975050505050505050565b611dac611fdd565b6001821415611dbd57829050611e4b565b6002821415611dd757611dd083846118cb565b9050611e4b565b611ddf61203e565b8360000151816000600381101515611df357fe5b6020020181815250508360200151816001600381101515611e1057fe5b60200201818152505082816002600381101515611e2957fe5b6020020181815250506040826060836007600019fa1515611e4957600080fd5b505b92915050565b60a060405190810160405280611e65612061565b8152602001611e72612061565b8152602001600081525090565b6116a06040519081016040528060008152602001611e9b61207b565b8152602001611ea861209f565b8152602001611eb56120c3565b8152602001611ec26120e7565b8152602001611ecf6120e7565b8152602001611edc612115565b8152602001611ee9612115565b8152602001611ef66120e7565b8152602001611f036120e7565b8152602001611f10612138565b8152602001611f1d612166565b8152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001611f55612115565b8152602001611f626120c3565b8152602001611f6f612166565b8152602001611f7c61218a565b8152602001611f896120c3565b8152602001611f966120c3565b8152602001611fa3612166565b8152602001611fb0612166565b8152602001611fbd612115565b8152602001611fca612115565b8152602001611fd7612115565b81525090565b604080519081016040528060008152602001600081525090565b608060405190810160405280600490602082028038833980820191505090505090565b61018060405190810160405280600c90602082028038833980820191505090505090565b606060405190810160405280600390602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b61028060405190810160405280601490602082028038833980820191505090505090565b61020060405190810160405280601090602082028038833980820191505090505090565b61014060405190810160405280600a90602082028038833980820191505090505090565b60a0604051908101604052806120fb612061565b8152602001612108612061565b8152602001600081525090565b608060405190810160405280600490602082028038833980820191505090505090565b6080604051908101604052806002905b612150612061565b8152602001906001900390816121485790505090565b61018060405190810160405280600c90602082028038833980820191505090505090565b60a06040519081016040528060059060208202803883398082019150509050509056fea165627a7a72305820311d2bdcf0331bf9e907a6edf75140624635cfe8c1823601f932e366a47a17fb0029`

// DeployPgcverifier deploys a new Ethereum contract, binding an instance of Pgcverifier to it.
func DeployPgcverifier(auth *bind.TransactOpts, backend bind.ContractBackend, params_ common.Address, dleSigmaVerifier_ common.Address, rangeProofVerifier_ common.Address, aggVerifier_ common.Address, sigmaVerifier_ common.Address) (common.Address, *types.Transaction, *Pgcverifier, error) {
	parsed, err := abi.JSON(strings.NewReader(PgcverifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PgcverifierBin), backend, params_, dleSigmaVerifier_, rangeProofVerifier_, aggVerifier_, sigmaVerifier_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pgcverifier{PgcverifierCaller: PgcverifierCaller{contract: contract}, PgcverifierTransactor: PgcverifierTransactor{contract: contract}, PgcverifierFilterer: PgcverifierFilterer{contract: contract}}, nil
}

// Pgcverifier is an auto generated Go binding around an Ethereum contract.
type Pgcverifier struct {
	PgcverifierCaller     // Read-only binding to the contract
	PgcverifierTransactor // Write-only binding to the contract
	PgcverifierFilterer   // Log filterer for contract events
}

// PgcverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type PgcverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgcverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PgcverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgcverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PgcverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgcverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PgcverifierSession struct {
	Contract     *Pgcverifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PgcverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PgcverifierCallerSession struct {
	Contract *PgcverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PgcverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PgcverifierTransactorSession struct {
	Contract     *PgcverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PgcverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type PgcverifierRaw struct {
	Contract *Pgcverifier // Generic contract binding to access the raw methods on
}

// PgcverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PgcverifierCallerRaw struct {
	Contract *PgcverifierCaller // Generic read-only contract binding to access the raw methods on
}

// PgcverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PgcverifierTransactorRaw struct {
	Contract *PgcverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPgcverifier creates a new instance of Pgcverifier, bound to a specific deployed contract.
func NewPgcverifier(address common.Address, backend bind.ContractBackend) (*Pgcverifier, error) {
	contract, err := bindPgcverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pgcverifier{PgcverifierCaller: PgcverifierCaller{contract: contract}, PgcverifierTransactor: PgcverifierTransactor{contract: contract}, PgcverifierFilterer: PgcverifierFilterer{contract: contract}}, nil
}

// NewPgcverifierCaller creates a new read-only instance of Pgcverifier, bound to a specific deployed contract.
func NewPgcverifierCaller(address common.Address, caller bind.ContractCaller) (*PgcverifierCaller, error) {
	contract, err := bindPgcverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PgcverifierCaller{contract: contract}, nil
}

// NewPgcverifierTransactor creates a new write-only instance of Pgcverifier, bound to a specific deployed contract.
func NewPgcverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*PgcverifierTransactor, error) {
	contract, err := bindPgcverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PgcverifierTransactor{contract: contract}, nil
}

// NewPgcverifierFilterer creates a new log filterer instance of Pgcverifier, bound to a specific deployed contract.
func NewPgcverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*PgcverifierFilterer, error) {
	contract, err := bindPgcverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PgcverifierFilterer{contract: contract}, nil
}

// bindPgcverifier binds a generic wrapper to an already deployed contract.
func bindPgcverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PgcverifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pgcverifier *PgcverifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pgcverifier.Contract.PgcverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pgcverifier *PgcverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pgcverifier.Contract.PgcverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pgcverifier *PgcverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pgcverifier.Contract.PgcverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pgcverifier *PgcverifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pgcverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pgcverifier *PgcverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pgcverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pgcverifier *PgcverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pgcverifier.Contract.contract.Transact(opts, method, params...)
}

// AggRangeProofVerifier is a free data retrieval call binding the contract method 0x293ec7b4.
//
// Solidity: function aggRangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) AggRangeProofVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "aggRangeProofVerifier")
	return *ret0, err
}

// AggRangeProofVerifier is a free data retrieval call binding the contract method 0x293ec7b4.
//
// Solidity: function aggRangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierSession) AggRangeProofVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.AggRangeProofVerifier(&_Pgcverifier.CallOpts)
}

// AggRangeProofVerifier is a free data retrieval call binding the contract method 0x293ec7b4.
//
// Solidity: function aggRangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) AggRangeProofVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.AggRangeProofVerifier(&_Pgcverifier.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Pgcverifier *PgcverifierCaller) BitSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "bitSize")
	return *ret0, err
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Pgcverifier *PgcverifierSession) BitSize() (*big.Int, error) {
	return _Pgcverifier.Contract.BitSize(&_Pgcverifier.CallOpts)
}

// BitSize is a free data retrieval call binding the contract method 0x3e8d3764.
//
// Solidity: function bitSize() constant returns(uint256)
func (_Pgcverifier *PgcverifierCallerSession) BitSize() (*big.Int, error) {
	return _Pgcverifier.Contract.BitSize(&_Pgcverifier.CallOpts)
}

// DleSigmaVerifier is a free data retrieval call binding the contract method 0x361eb474.
//
// Solidity: function dleSigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) DleSigmaVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "dleSigmaVerifier")
	return *ret0, err
}

// DleSigmaVerifier is a free data retrieval call binding the contract method 0x361eb474.
//
// Solidity: function dleSigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierSession) DleSigmaVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.DleSigmaVerifier(&_Pgcverifier.CallOpts)
}

// DleSigmaVerifier is a free data retrieval call binding the contract method 0x361eb474.
//
// Solidity: function dleSigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) DleSigmaVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.DleSigmaVerifier(&_Pgcverifier.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierCaller) G(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Pgcverifier.contract.Call(opts, out, "g")
	return *ret, err
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Pgcverifier.Contract.G(&_Pgcverifier.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierCallerSession) G() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Pgcverifier.Contract.G(&_Pgcverifier.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierCaller) H(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _Pgcverifier.contract.Call(opts, out, "h")
	return *ret, err
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Pgcverifier.Contract.H(&_Pgcverifier.CallOpts)
}

// H is a free data retrieval call binding the contract method 0xb8c9d365.
//
// Solidity: function h() constant returns(uint256 X, uint256 Y)
func (_Pgcverifier *PgcverifierCallerSession) H() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Pgcverifier.Contract.H(&_Pgcverifier.CallOpts)
}

// MaxNumber is a free data retrieval call binding the contract method 0x3a4f6999.
//
// Solidity: function maxNumber() constant returns(uint256)
func (_Pgcverifier *PgcverifierCaller) MaxNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "maxNumber")
	return *ret0, err
}

// MaxNumber is a free data retrieval call binding the contract method 0x3a4f6999.
//
// Solidity: function maxNumber() constant returns(uint256)
func (_Pgcverifier *PgcverifierSession) MaxNumber() (*big.Int, error) {
	return _Pgcverifier.Contract.MaxNumber(&_Pgcverifier.CallOpts)
}

// MaxNumber is a free data retrieval call binding the contract method 0x3a4f6999.
//
// Solidity: function maxNumber() constant returns(uint256)
func (_Pgcverifier *PgcverifierCallerSession) MaxNumber() (*big.Int, error) {
	return _Pgcverifier.Contract.MaxNumber(&_Pgcverifier.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Pgcverifier *PgcverifierCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "n")
	return *ret0, err
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Pgcverifier *PgcverifierSession) N() (*big.Int, error) {
	return _Pgcverifier.Contract.N(&_Pgcverifier.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() constant returns(uint256)
func (_Pgcverifier *PgcverifierCallerSession) N() (*big.Int, error) {
	return _Pgcverifier.Contract.N(&_Pgcverifier.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) Params(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "params")
	return *ret0, err
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() constant returns(address)
func (_Pgcverifier *PgcverifierSession) Params() (common.Address, error) {
	return _Pgcverifier.Contract.Params(&_Pgcverifier.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) Params() (common.Address, error) {
	return _Pgcverifier.Contract.Params(&_Pgcverifier.CallOpts)
}

// RangeProofVerifier is a free data retrieval call binding the contract method 0x61e5af5e.
//
// Solidity: function rangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) RangeProofVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "rangeProofVerifier")
	return *ret0, err
}

// RangeProofVerifier is a free data retrieval call binding the contract method 0x61e5af5e.
//
// Solidity: function rangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierSession) RangeProofVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.RangeProofVerifier(&_Pgcverifier.CallOpts)
}

// RangeProofVerifier is a free data retrieval call binding the contract method 0x61e5af5e.
//
// Solidity: function rangeProofVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) RangeProofVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.RangeProofVerifier(&_Pgcverifier.CallOpts)
}

// SigmaVerifier is a free data retrieval call binding the contract method 0x81b7bc73.
//
// Solidity: function sigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCaller) SigmaVerifier(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pgcverifier.contract.Call(opts, out, "sigmaVerifier")
	return *ret0, err
}

// SigmaVerifier is a free data retrieval call binding the contract method 0x81b7bc73.
//
// Solidity: function sigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierSession) SigmaVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.SigmaVerifier(&_Pgcverifier.CallOpts)
}

// SigmaVerifier is a free data retrieval call binding the contract method 0x81b7bc73.
//
// Solidity: function sigmaVerifier() constant returns(address)
func (_Pgcverifier *PgcverifierCallerSession) SigmaVerifier() (common.Address, error) {
	return _Pgcverifier.Contract.SigmaVerifier(&_Pgcverifier.CallOpts)
}

// VerifyAggTransfer is a paid mutator transaction binding the contract method 0xa89e0e44.
//
// Solidity: function verifyAggTransfer(uint256[36] points, uint256[10] scalar, uint256[12] l, uint256[12] r, uint256[4] ub, uint256 nonce) returns(bool)
func (_Pgcverifier *PgcverifierTransactor) VerifyAggTransfer(opts *bind.TransactOpts, points [36]*big.Int, scalar [10]*big.Int, l [12]*big.Int, r [12]*big.Int, ub [4]*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Pgcverifier.contract.Transact(opts, "verifyAggTransfer", points, scalar, l, r, ub, nonce)
}

// VerifyAggTransfer is a paid mutator transaction binding the contract method 0xa89e0e44.
//
// Solidity: function verifyAggTransfer(uint256[36] points, uint256[10] scalar, uint256[12] l, uint256[12] r, uint256[4] ub, uint256 nonce) returns(bool)
func (_Pgcverifier *PgcverifierSession) VerifyAggTransfer(points [36]*big.Int, scalar [10]*big.Int, l [12]*big.Int, r [12]*big.Int, ub [4]*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Pgcverifier.Contract.VerifyAggTransfer(&_Pgcverifier.TransactOpts, points, scalar, l, r, ub, nonce)
}

// VerifyAggTransfer is a paid mutator transaction binding the contract method 0xa89e0e44.
//
// Solidity: function verifyAggTransfer(uint256[36] points, uint256[10] scalar, uint256[12] l, uint256[12] r, uint256[4] ub, uint256 nonce) returns(bool)
func (_Pgcverifier *PgcverifierTransactorSession) VerifyAggTransfer(points [36]*big.Int, scalar [10]*big.Int, l [12]*big.Int, r [12]*big.Int, ub [4]*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Pgcverifier.Contract.VerifyAggTransfer(&_Pgcverifier.TransactOpts, points, scalar, l, r, ub, nonce)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x26f9031b.
//
// Solidity: function verifyBurn(uint256 amount, uint256[2] publicKey, uint256[4] proof, uint256 z, uint256[4] ub, uint256[] input) returns(bool)
func (_Pgcverifier *PgcverifierTransactor) VerifyBurn(opts *bind.TransactOpts, amount *big.Int, publicKey [2]*big.Int, proof [4]*big.Int, z *big.Int, ub [4]*big.Int, input []*big.Int) (*types.Transaction, error) {
	return _Pgcverifier.contract.Transact(opts, "verifyBurn", amount, publicKey, proof, z, ub, input)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x26f9031b.
//
// Solidity: function verifyBurn(uint256 amount, uint256[2] publicKey, uint256[4] proof, uint256 z, uint256[4] ub, uint256[] input) returns(bool)
func (_Pgcverifier *PgcverifierSession) VerifyBurn(amount *big.Int, publicKey [2]*big.Int, proof [4]*big.Int, z *big.Int, ub [4]*big.Int, input []*big.Int) (*types.Transaction, error) {
	return _Pgcverifier.Contract.VerifyBurn(&_Pgcverifier.TransactOpts, amount, publicKey, proof, z, ub, input)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x26f9031b.
//
// Solidity: function verifyBurn(uint256 amount, uint256[2] publicKey, uint256[4] proof, uint256 z, uint256[4] ub, uint256[] input) returns(bool)
func (_Pgcverifier *PgcverifierTransactorSession) VerifyBurn(amount *big.Int, publicKey [2]*big.Int, proof [4]*big.Int, z *big.Int, ub [4]*big.Int, input []*big.Int) (*types.Transaction, error) {
	return _Pgcverifier.Contract.VerifyBurn(&_Pgcverifier.TransactOpts, amount, publicKey, proof, z, ub, input)
}
