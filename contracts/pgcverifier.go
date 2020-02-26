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
const PgcverifierABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"aggRangeProofVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dleSigmaVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bitSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rangeProofVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sigmaVerifier\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"h\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"params_\",\"type\":\"address\"},{\"name\":\"dleSigmaVerifier_\",\"type\":\"address\"},{\"name\":\"rangeProofVerifier_\",\"type\":\"address\"},{\"name\":\"aggVerifier_\",\"type\":\"address\"},{\"name\":\"sigmaVerifier_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"points\",\"type\":\"uint256[36]\"},{\"name\":\"scalar\",\"type\":\"uint256[10]\"},{\"name\":\"l\",\"type\":\"uint256[14]\"},{\"name\":\"r\",\"type\":\"uint256[14]\"},{\"name\":\"ub\",\"type\":\"uint256[4]\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"verifyAggTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"publicKey\",\"type\":\"uint256[2]\"},{\"name\":\"proof\",\"type\":\"uint256[4]\"},{\"name\":\"z\",\"type\":\"uint256\"},{\"name\":\"ub\",\"type\":\"uint256[4]\"},{\"name\":\"input\",\"type\":\"uint256[]\"}],\"name\":\"verifyBurn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PgcverifierBin is the compiled bytecode used for deploying new contracts.
const PgcverifierBin = `0x60806040523480156200001157600080fd5b5060405160a08062002720833981018060405260a08110156200003357600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919050505084600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001bb62000515565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200023f57600080fd5b505afa15801562000254573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200027a57600080fd5b810190809190505090506200028e62000515565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156200031257600080fd5b505afa15801562000327573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200034d57600080fd5b810190809190505090508160006002811015156200036757fe5b602002015160008001819055508160016002811015156200038457fe5b6020020151600060010181905550806000600281101515620003a257fe5b6020020151600260000181905550806001600281101515620003c057fe5b6020020151600260010181905550600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da8972246040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156200045357600080fd5b505afa15801562000468573d6000803e3d6000fd5b505050506040513d60208110156200047f57600080fd5b8101908080519060200190929190505050604014151562000508576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f62697473697a65206e6f7420657175616c00000000000000000000000000000081525060200191505060405180910390fd5b5050505050505062000537565b6040805190810160405280600290602082028038833980820191505090505090565b6121d980620005476000396000f3fe6080604052600436106100ba576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806323757b64146100bf57806326f9031b1461025c578063293ec7b4146104115780632e52d60614610468578063361eb474146104935780633a4f6999146104ea5780633e8d37641461051557806361e5af5e1461054057806381b7bc7314610597578063b8c9d365146105ee578063cff0ab9614610620578063e2179b8e14610677575b600080fd5b3480156100cb57600080fd5b5061024260048036036109e08110156100e357600080fd5b810190808061048001906024806020026040519081016040528092919082602460200280828437600081840152601f19601f8201169050808301925050505050509192919290806101400190600a806020026040519081016040528092919082600a60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101c00190600e806020026040519081016040528092919082600e60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101c00190600e806020026040519081016040528092919082600e60200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291905050506106a9565b604051808215151515815260200191505060405180910390f35b34801561026857600080fd5b506103f760048036036101a081101561028057600080fd5b810190808035906020019092919080604001906002806020026040519081016040528092919082600260200280828437600081840152601f19601f820116905080830192505050505050919291929080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080608001906004806020026040519081016040528092919082600460200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019064010000000081111561037457600080fd5b82018360208201111561038657600080fd5b803590602001918460208302840111640100000000831117156103a857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050611256565b604051808215151515815260200191505060405180910390f35b34801561041d57600080fd5b506104266113af565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561047457600080fd5b5061047d6113d5565b6040518082815260200191505060405180910390f35b34801561049f57600080fd5b506104a86113da565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f657600080fd5b506104ff611400565b6040518082815260200191505060405180910390f35b34801561052157600080fd5b5061052a611408565b6040518082815260200191505060405180910390f35b34801561054c57600080fd5b5061055561140d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105a357600080fd5b506105ac611433565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105fa57600080fd5b50610603611459565b604051808381526020018281526020019250505060405180910390f35b34801561062c57600080fd5b5061063561146b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561068357600080fd5b5061068c611491565b604051808381526020018281526020019250505060405180910390f35b60006106b3611e51565b8360006004811015156106c257fe5b6020020151816000015160000181815250508360016004811015156106e357fe5b60200201518160000151602001818152505083600260048110151561070457fe5b60200201518160200151600001818152505083600360048110151561072557fe5b60200201518160200151602001818152505061073f611e7f565b60008160000181815250505b6010816000015110156107a35788816000015160248110151561076a57fe5b60200201518160400151826000015160108110151561078557fe5b6020020181815250508060000180518091906001018152505061074b565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e1efe76682604001518a6000600a811015156107f557fe5b60200201518b6001600a8110151561080957fe5b60200201516040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084601060200280838360005b83811015610864578082015181840152602081019050610849565b505050509050018381526020018281526020019350505050602060405180830381600087803b15801561089657600080fd5b505af11580156108aa573d6000803e3d6000fd5b505050506040513d60208110156108c057600080fd5b81019080805190602001909291905050501515610945576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f70746520657175616c2070726f6f6620696e76616c696400000000000000000081525060200191505060405180910390fd5b60408051908101604052808a600460248110151561095f57fe5b602002015181526020018a600560248110151561097857fe5b602002015181525081608001516000018190525060408051908101604052808a60086024811015156109a657fe5b602002015181526020018a60096024811015156109bf57fe5b60200201518152508160800151602001819052506109fa6109e78260800151600001516114a3565b836000015161153390919063ffffffff16565b81610100015160000181905250610a2e610a1b8260800151602001516114a3565b836020015161153390919063ffffffff16565b8161010001516020018190525060408051908101604052808a6010602481101515610a5557fe5b602002015181526020018a6011602481101515610a6e57fe5b60200201518152508161012001516000018190525060408051908101604052808a6012602481101515610a9d57fe5b602002015181526020018a6013602481101515610ab657fe5b60200201518152508161012001516020018190525060408051908101604052808a6018602481101515610ae557fe5b602002015181526020018a6019602481101515610afe57fe5b60200201518152508161014001516000600281101515610b1a57fe5b602002018190525060408051908101604052808a601a602481101515610b3c57fe5b602002015181526020018a601b602481101515610b5557fe5b60200201518152508161014001516001600281101515610b7157fe5b6020020181905250610bd28161010001518261012001518361014001518c6000602481101515610b9d57fe5b60200201518d6001602481101515610bb157fe5b60200201518d6004600a81101515610bc557fe5b602002015160018b6115d5565b1515610c46576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f646c65207369676d612070726f6f66206661696c65640000000000000000000081525060200191505060405180910390fd5b886000602481101515610c5557fe5b602002015181606001516000600a81101515610c6d57fe5b602002018181525050886001602481101515610c8557fe5b602002015181606001516001600a81101515610c9d57fe5b60200201818152505060108160000181815250505b601881600001511015610d1057888160000151602481101515610cd157fe5b602002015181606001516010836000015160020103600a81101515610cf257fe5b60200201818152505080600001805180919060010181525050610cb2565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663991395ca82606001518a6002600a81101515610d6257fe5b60200201518b6003600a81101515610d7657fe5b60200201516040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600a60200280838360005b83811015610dd1578082015181840152602081019050610db6565b505050509050018381526020018281526020019350505050602060405180830381600087803b158015610e0357600080fd5b505af1158015610e17573d6000803e3d6000fd5b505050506040513d6020811015610e2d57600080fd5b81019080805190602001909291905050501515610eb2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f63742076616c69642070726f6f6620696e76616c69640000000000000000000081525060200191505060405180910390fd5b60008160000181815250505b600881600001511015610f1a57888160000151601c01602481101515610ee057fe5b60200201518161022001518260000151600c81101515610efc57fe5b60200201818152505080600001805180919060010181525050610ebe565b886008602481101515610f2957fe5b60200201518161022001516008600c81101515610f4257fe5b602002018181525050886009602481101515610f5a57fe5b60200201518161022001516009600c81101515610f7357fe5b602002018181525050886012602481101515610f8b57fe5b6020020151816102200151600a600c81101515610fa457fe5b602002018181525050886013602481101515610fbc57fe5b6020020151816102200151600b600c81101515610fd557fe5b60200201818152505060008160000181815250505b60058160000151101561104657876005826000015101600a8110151561100c57fe5b6020020151816102400151826000015160058110151561102857fe5b60200201818152505080600001805180919060010181525050610fea565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e5d34a638261022001518361024001518a8a6040518563ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018085600c60200280838360005b838110156110e85780820151818401526020810190506110cd565b5050505090500184600560200280838360005b838110156111165780820151818401526020810190506110fb565b5050505090500183600e60200280838360005b83811015611144578082015181840152602081019050611129565b5050505090500182600e60200280838360005b83811015611172578082015181840152602081019050611157565b5050505090500194505050505060206040518083038186803b15801561119757600080fd5b505afa1580156111ab573d6000803e3d6000fd5b505050506040513d60208110156111c157600080fd5b81019080805190602001909291905050501515611246576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f616767726174652072616e67652070726f6f6620696e76616c6964000000000081525060200191505060405180910390fd5b6001925050509695505050505050565b6000611260611e51565b83600060048110151561126f57fe5b60200201518160000151600001818152505083600160048110151561129057fe5b6020020151816000015160200181815250508360026004811015156112b157fe5b6020020151816020015160000181815250508360036004811015156112d257fe5b60200201518160200151602001818152505061132c888260408051908101604052808b600060028110151561130357fe5b602002015181526020018b600160028110151561131c57fe5b6020020151815250898988611a0c565b15156113a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f646c65207369676d6120766572696679206661696c656400000000000000000081525060200191505060405180910390fd5b60019150509695505050505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604060020a81565b604081565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060000154908060010154905082565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028060000154908060010154905082565b6114ab611fdd565b600082600001511480156114c3575060008260200151145b156114e6576040805190810160405280600081526020016000815250905061152e565b60007f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd479050604080519081016040528084600001518152602001846020015183038152509150505b919050565b61153b611fdd565b611543611ff7565b836000015181600060048110151561155757fe5b602002018181525050836020015181600160048110151561157457fe5b602002018181525050826000015181600260048110151561159157fe5b60200201818152505082602001518160036004811015156115ae57fe5b6020020181815250506040826080836006600019fa15156115ce57600080fd5b5092915050565b60006115df611fdd565b6116026115ef8b602001516114a3565b8a6020015161153390919063ffffffff16565b905061160c611fdd565b61162f61161c8c600001516114a3565b8b6000015161153390919063ffffffff16565b905061163961201a565b89600060028110151561164857fe5b602002015160000151816000600c8110151561166057fe5b60200201818152505089600060028110151561167857fe5b602002015160200151816001600c8110151561169057fe5b6020020181815250508960016002811015156116a857fe5b602002015160000151816002600c811015156116c057fe5b6020020181815250508960016002811015156116d857fe5b602002015160200151816003600c811015156116f057fe5b6020020181815250508260000151816004600c8110151561170d57fe5b6020020181815250508260200151816005600c8110151561172a57fe5b6020020181815250508160000151816006600c8110151561174757fe5b6020020181815250508160200151816007600c8110151561176457fe5b6020020181815250506000800154816008600c8110151561178157fe5b602002018181525050600060010154816009600c8110151561179f57fe5b6020020181815250508881600a600c811015156117b857fe5b6020020181815250508781600b600c811015156117d157fe5b60200201818152505060008614156118e757600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639751cb1382896040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600c60200280838360005b8381101561187957808201518184015260208101905061185e565b505050509050018281526020019250505060206040518083038186803b1580156118a257600080fd5b505afa1580156118b6573d6000803e3d6000fd5b505050506040513d60208110156118cc57600080fd5b81019080805190602001909291905050509350505050611a00565b60018614156119fc57600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630581c0408289886040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600c60200280838360005b8381101561198757808201518184015260208101905061196c565b50505050905001838152602001828152602001935050505060206040518083038186803b1580156119b757600080fd5b505afa1580156119cb573d6000803e3d6000fd5b505050506040513d60208110156119e157600080fd5b81019080805190602001909291905050509350505050611a00565b5050505b98975050505050505050565b6000611a16611fdd565b611a67611a54611a4f8a600260408051908101604052908160008201548152602001600182015481525050611da490919063ffffffff16565b6114a3565b886020015161153390919063ffffffff16565b9050611a71611e7f565b856000600481101515611a8057fe5b60200201518161016001516000600c81101515611a9957fe5b602002018181525050856001600481101515611ab157fe5b60200201518161016001516001600c81101515611aca57fe5b602002018181525050856002600481101515611ae257fe5b60200201518161016001516002600c81101515611afb57fe5b602002018181525050856003600481101515611b1357fe5b60200201518161016001516003600c81101515611b2c57fe5b60200201818152505081600001518161016001516004600c81101515611b4e57fe5b60200201818152505081602001518161016001516005600c81101515611b7057fe5b6020020181815250508760000151600001518161016001516006600c81101515611b9657fe5b6020020181815250508760000151602001518161016001516007600c81101515611bbc57fe5b60200201818152505060008001548161016001516008600c81101515611bde57fe5b6020020181815250506000600101548161016001516009600c81101515611c0157fe5b6020020181815250508660000151816101600151600a600c81101515611c2357fe5b6020020181815250508660200151816101600151600b600c81101515611c4557fe5b602002018181525050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e26d0b882610160015187876040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600c60200280838360005b83811015611cea578082015181840152602081019050611ccf565b5050505090500183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015611d36578082015181840152602081019050611d1b565b5050505090500194505050505060206040518083038186803b158015611d5b57600080fd5b505afa158015611d6f573d6000803e3d6000fd5b505050506040513d6020811015611d8557600080fd5b8101908080519060200190929190505050925050509695505050505050565b611dac611fdd565b6001821415611dbd57829050611e4b565b6002821415611dd757611dd08384611533565b9050611e4b565b611ddf61203e565b8360000151816000600381101515611df357fe5b6020020181815250508360200151816001600381101515611e1057fe5b60200201818152505082816002600381101515611e2957fe5b6020020181815250506040826060836007600019fa1515611e4957600080fd5b505b92915050565b60a060405190810160405280611e65612061565b8152602001611e72612061565b8152602001600081525090565b6117206040519081016040528060008152602001611e9b61207b565b8152602001611ea861209f565b8152602001611eb56120c3565b8152602001611ec26120e7565b8152602001611ecf6120e7565b8152602001611edc612115565b8152602001611ee9612115565b8152602001611ef66120e7565b8152602001611f036120e7565b8152602001611f10612138565b8152602001611f1d612166565b8152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001611f55612115565b8152602001611f626120c3565b8152602001611f6f612166565b8152602001611f7c61218a565b8152602001611f89612166565b8152602001611f96612166565b8152602001611fa3612166565b8152602001611fb0612166565b8152602001611fbd612115565b8152602001611fca612115565b8152602001611fd7612115565b81525090565b604080519081016040528060008152602001600081525090565b608060405190810160405280600490602082028038833980820191505090505090565b61018060405190810160405280600c90602082028038833980820191505090505090565b606060405190810160405280600390602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b61028060405190810160405280601490602082028038833980820191505090505090565b61020060405190810160405280601090602082028038833980820191505090505090565b61014060405190810160405280600a90602082028038833980820191505090505090565b60a0604051908101604052806120fb612061565b8152602001612108612061565b8152602001600081525090565b608060405190810160405280600490602082028038833980820191505090505090565b6080604051908101604052806002905b612150612061565b8152602001906001900390816121485790505090565b61018060405190810160405280600c90602082028038833980820191505090505090565b60a06040519081016040528060059060208202803883398082019150509050509056fea165627a7a723058207f930404b1df81445c934737463b5da42cfd50b3d02dd50d865b3f346944b8740029`

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

// VerifyAggTransfer is a paid mutator transaction binding the contract method 0x23757b64.
//
// Solidity: function verifyAggTransfer(uint256[36] points, uint256[10] scalar, uint256[14] l, uint256[14] r, uint256[4] ub, uint256 nonce) returns(bool)
func (_Pgcverifier *PgcverifierTransactor) VerifyAggTransfer(opts *bind.TransactOpts, points [36]*big.Int, scalar [10]*big.Int, l [14]*big.Int, r [14]*big.Int, ub [4]*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Pgcverifier.contract.Transact(opts, "verifyAggTransfer", points, scalar, l, r, ub, nonce)
}

// VerifyAggTransfer is a paid mutator transaction binding the contract method 0x23757b64.
//
// Solidity: function verifyAggTransfer(uint256[36] points, uint256[10] scalar, uint256[14] l, uint256[14] r, uint256[4] ub, uint256 nonce) returns(bool)
func (_Pgcverifier *PgcverifierSession) VerifyAggTransfer(points [36]*big.Int, scalar [10]*big.Int, l [14]*big.Int, r [14]*big.Int, ub [4]*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Pgcverifier.Contract.VerifyAggTransfer(&_Pgcverifier.TransactOpts, points, scalar, l, r, ub, nonce)
}

// VerifyAggTransfer is a paid mutator transaction binding the contract method 0x23757b64.
//
// Solidity: function verifyAggTransfer(uint256[36] points, uint256[10] scalar, uint256[14] l, uint256[14] r, uint256[4] ub, uint256 nonce) returns(bool)
func (_Pgcverifier *PgcverifierTransactorSession) VerifyAggTransfer(points [36]*big.Int, scalar [10]*big.Int, l [14]*big.Int, r [14]*big.Int, ub [4]*big.Int, nonce *big.Int) (*types.Transaction, error) {
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
