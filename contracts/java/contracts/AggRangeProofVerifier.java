package contracts;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.Callable;
import org.web3j.abi.FunctionEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Bool;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.tuples.generated.Tuple2;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.ContractGasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 3.6.0.
 */
public class AggRangeProofVerifier extends Contract {
    private static final String BINARY = "0x60806040523480156200001157600080fd5b50604051620025fe380380620025fe833981810160405260208110156200003757600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da8972246040518163ffffffff1660e01b815260040160206040518083038186803b158015620000f057600080fd5b505afa15801562000105573d6000803e3d6000fd5b505050506040513d60208110156200011c57600080fd5b8101908080519060200190929190505050604014620001a3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6269732073697a65206e6f7420657175616c000000000000000000000000000081525060200191505060405180910390fd5b620001ad620006d0565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff1660e01b8152600401604080518083038186803b1580156200021457600080fd5b505afa15801562000229573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200024f57600080fd5b8101908091905050905062000263620006d0565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff1660e01b8152600401604080518083038186803b158015620002ca57600080fd5b505afa158015620002df573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156200030557600080fd5b8101908091905050905062000319620006d0565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324d6147d6040518163ffffffff1660e01b8152600401604080518083038186803b1580156200038057600080fd5b505afa15801562000395573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506040811015620003bb57600080fd5b81019080919050509050620003cf620006f2565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634db118756040518163ffffffff1660e01b81526004016120006040518083038186803b1580156200043857600080fd5b505afa1580156200044d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506120008110156200047457600080fd5b8101908091905050905062000488620006f2565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663da2b99ce6040518163ffffffff1660e01b81526004016120006040518083038186803b158015620004f157600080fd5b505afa15801562000506573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506120008110156200052d57600080fd5b81019080919050509050846000600281106200054557fe5b6020020151600160000181905550846001600281106200056157fe5b60200201516001800181905550836000600281106200057c57fe5b6020020151600360000181905550836001600281106200059857fe5b602002015160036001018190555082600060028110620005b457fe5b602002015160056000018190555082600160028110620005d057fe5b602002015160056001018190555060008090505b6002604002811015620006c357828160020261010081106200060257fe5b6020020151600782608081106200061557fe5b6002020160000181905550826001826002020161010081106200063457fe5b6020020151600782608081106200064757fe5b6002020160010181905550818160020261010081106200066357fe5b602002015161010782608081106200067757fe5b6002020160000181905550816001826002020161010081106200069657fe5b60200201516101078260808110620006aa57fe5b60020201600101819055508080600101915050620005e4565b5050505050505062000716565b6040518060400160405280600290602082028038833980820191505090505090565b60405180612000016040528061010090602082028038833980820191505090505090565b611ed880620007266000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063b8c9d3651161005b578063b8c9d36514610139578063cff0ab961461015e578063e2179b8e146101a8578063e5d34a63146101cd5761007d565b80631213e20114610082578063654474ee146100a7578063b7479f5f146100f0575b600080fd5b61008a610311565b604051808381526020018281526020019250505060405180910390f35b6100d3600480360360208110156100bd57600080fd5b8101908080359060200190929190505050610323565b604051808381526020018281526020019250505060405180910390f35b61011c6004803603602081101561010657600080fd5b810190808035906020019092919050505061034b565b604051808381526020018281526020019250505060405180910390f35b610141610372565b604051808381526020018281526020019250505060405180910390f35b610166610384565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101b06103a9565b604051808381526020018281526020019250505060405180910390f35b6102f760048036036105a08110156101e457600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f82011690508083019250505050505091929192908060a001906005806020026040519081016040528092919082600560200280828437600081840152601f19601f8201169050808301925050505050509192919290806101c00190600e806020026040519081016040528092919082600e60200280828437600081840152601f19601f8201169050808301925050505050509192919290806101c00190600e806020026040519081016040528092919082600e60200280828437600081840152601f19601f82011690508083019250505050505091929192905050506103bb565b604051808215151515815260200191505060405180910390f35b60058060000154908060010154905082565b610107816080811061033157fe5b600202016000915090508060000154908060010154905082565b6007816080811061035857fe5b600202016000915090508060000154908060010154905082565b60038060000154908060010154905082565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60018060000154908060010154905082565b60006103c5611af4565b6040518060400160405280876000600c81106103dd57fe5b60200201518152602001876001600c81106103f457fe5b602002015181525081600001819052506040518060400160405280876002600c811061041c57fe5b60200201518152602001876003600c811061043357fe5b602002015181525081602001819052506040518060400160405280876004600c811061045b57fe5b60200201518152602001876005600c811061047257fe5b602002015181525081604001819052506040518060400160405280876006600c811061049a57fe5b60200201518152602001876007600c81106104b157fe5b60200201518152508160600181905250846000600581106104ce57fe5b6020020151816080018181525050846001600581106104e957fe5b60200201518160a00181815250508460026005811061050457fe5b60200201518160c001818152505061051a611b57565b6040518060400160405280886008600c811061053257fe5b60200201518152602001886009600c811061054957fe5b60200201518152508160006002811061055e57fe5b6020020181905250604051806040016040528088600a600c811061057e57fe5b6020020151815260200188600b600c811061059557fe5b6020020151815250816001600281106105aa57fe5b6020020181905250848260e0015160000181905250838260e0015160200181905250856003600581106105d957fe5b60200201518260e001516040018181525050856004600581106105f857fe5b60200201518260e0015160600181815250506106148183610620565b92505050949350505050565b600061062a611b84565b610632611306565b905061063c611b84565b610644611399565b905061064e611bb2565b61067a85600001516000015186600001516020015187602001516000015188602001516020015161142e565b8160a00181815250506106af85602001516000015186602001516020015187600001516000015188600001516020015161142e565b816101000181815250506106c68160a00151611484565b8160c001819052506106e36106de8260a00151611524565b611484565b8160e001819052506106f98161010001516115ee565b8161016001818152505061072861072382610100015183610100015161161a90919063ffffffff16565b611654565b8161012001818152505061075761075282610100015183610120015161161a90919063ffffffff16565b611654565b8161014001818152505061076b600261168c565b8161018001819052506107a085604001516000015186604001516020015187606001516000015188606001516020015161142e565b8160600181815250506107c48160600151826060015161161a90919063ffffffff16565b8160800181815250506108886107eb8260800151876060015161172990919063ffffffff16565b61087a6108098460600151896040015161172990919063ffffffff16565b61086c6108358661014001518c60016002811061082257fe5b602002015161172990919063ffffffff16565b61085e8761012001518d60006002811061084b57fe5b602002015161172990919063ffffffff16565b6117ce90919063ffffffff16565b6117ce90919063ffffffff16565b6117ce90919063ffffffff16565b816101c001819052506109026108b96108a5836101800151611866565b83610140015161161a90919063ffffffff16565b6108f46108c98460c001516118c9565b6108e685610120015186610100015161192f90919063ffffffff16565b61161a90919063ffffffff16565b61192f90919063ffffffff16565b8161020001818152505061095f61094b610920836101800151611866565b61093d84610100015185610140015161161a90919063ffffffff16565b61161a90919063ffffffff16565b82610200015161192f90919063ffffffff16565b816102000181815250506109ff6109a38660a0015160036040518060400160405290816000820154815260200160018201548152505061172990919063ffffffff16565b6109f16109c2846102000151896080015161192f90919063ffffffff16565b60016040518060400160405290816000820154815260200160018201548152505061172990919063ffffffff16565b6117ce90919063ffffffff16565b816101e00181905250806101e0015160000151816101c0015160000151141580610a3b5750806101e0015160200151816101c001516020015114155b15610a4c5760009350505050611300565b610a54611cd4565b610a85610a728360600151886020015161172990919063ffffffff16565b87600001516117ce90919063ffffffff16565b905060008090505b6001600601811015610cdd576000610b198860e001516000015183600202600e8110610ab557fe5b60200201518960e001516000015160018560020201600e8110610ad457fe5b60200201518a60e001516020015185600202600e8110610af057fe5b60200201518b60e001516020015160018760020201600e8110610b0f57fe5b602002015161142e565b9050808461022001518360078110610b2d57fe5b602002018181525050610b51610b4c828361161a90919063ffffffff16565b611654565b8461026001518360078110610b6257fe5b602002018181525050610b898461026001518360078110610b7f57fe5b6020020151611524565b8461028001518360078110610b9a57fe5b60200201818152505060405180604001604052808960e001516000015184600202600e8110610bc557fe5b602002015181526020018960e001516000015160018560020201600e8110610be957fe5b60200201518152508460400181905250610c37610c288561026001518460078110610c1057fe5b6020020151866040015161172990919063ffffffff16565b846117ce90919063ffffffff16565b925060405180604001604052808960e001516020015184600202600e8110610c5b57fe5b602002015181526020018960e001516020015160018560020201600e8110610c7f57fe5b60200201518152508460400181905250610ccd610cbe8561028001518460078110610ca657fe5b6020020151866040015161172990919063ffffffff16565b846117ce90919063ffffffff16565b9250508080600101915050610a8d565b5060008090505b6002604002811015611187576000811415610e155760008090505b6001600601811015610da75760008461022001518260078110610d1e57fe5b602002015190506000821415610d4e5780856102c001518460808110610d4057fe5b602002018181525050610d99565b610d7e610d7982876102c001518660808110610d6657fe5b602002015161161a90919063ffffffff16565b611654565b856102c001518460808110610d8f57fe5b6020020181815250505b508080600101915050610cff565b50826102c001518160808110610db957fe5b60200201518361030001518260808110610dcf57fe5b602002018181525050610df6836102c001518260808110610dec57fe5b6020020151611524565b836102c001518260808110610e0757fe5b602002018181525050610f0d565b6000610e25826001600601611974565b9050610e7f610e7a8561026001518360016006010360078110610e4457fe5b6020020151866102c00151610e5b600186036119b9565b860360808110610e6757fe5b602002015161161a90919063ffffffff16565b611654565b846102c001518360808110610e9057fe5b602002018181525050610ef1610eec8561028001518360016006010360078110610eb657fe5b6020020151866103000151610ecd600186036119b9565b860360808110610ed957fe5b602002015161161a90919063ffffffff16565b611654565b8461030001518360808110610f0257fe5b602002018181525050505b826102c001518160808110610f1e57fe5b6020020151836102a001518260808110610f3457fe5b6020020181815250508261030001518160808110610f4e57fe5b6020020151836102e001518260808110610f6457fe5b602002018181525050610fb4836101000151610fa68960e0015160400151866102a001518560808110610f9357fe5b602002015161161a90919063ffffffff16565b611a0190919063ffffffff16565b836102a001518260808110610fc557fe5b602002018181525050610ffe8760e0015160600151846102e001518360808110610feb57fe5b602002015161161a90919063ffffffff16565b836102e00151826080811061100f57fe5b602002018181525050604081101561109457611075611051846101800151836040811061103857fe5b602002015185610120015161161a90919063ffffffff16565b846102e00151836080811061106257fe5b602002015161192f90919063ffffffff16565b836102e00151826080811061108657fe5b60200201818152505061110d565b6110f26110ce846101800151604084816110aa57fe5b06604081106110b557fe5b602002015185610140015161161a90919063ffffffff16565b846102e0015183608081106110df57fe5b602002015161192f90919063ffffffff16565b836102e00151826080811061110357fe5b6020020181815250505b6111608361010001516111528560e00151846080811061112957fe5b6020020151866102e00151856080811061113f57fe5b602002015161161a90919063ffffffff16565b61192f90919063ffffffff16565b836102e00151826080811061117157fe5b6020020181815250508080600101915050610ce4565b5060006111c08760800151604051602001808281526020019150506040516020818303038152906040528051906020012060001c611654565b90506112c76111fc8860c0015160036040518060400160405290816000820154815260200160018201548152505061172990919063ffffffff16565b6112b961127c61124d61123e8c608001516112308e60e00151606001518f60e001516040015161161a90919063ffffffff16565b61192f90919063ffffffff16565b8661161a90919063ffffffff16565b60056040518060400160405290816000820154815260200160018201548152505061172990919063ffffffff16565b6112ab61128e89896102e00151611a3b565b61129d8b8a6102a00151611a3b565b6117ce90919063ffffffff16565b6117ce90919063ffffffff16565b6117ce90919063ffffffff16565b836101e001819052508160000151836101e00151600001511480156112f857508160200151836101e0015160200151145b955050505050505b92915050565b61130e611b84565b611316611b84565b60008090505b6002604002811015611391576007816080811061133557fe5b600202016000015482826080811061134957fe5b602002015160000181815250506007816080811061136357fe5b600202016001015482826080811061137757fe5b60200201516020018181525050808060010191505061131c565b508091505090565b6113a1611b84565b6113a9611b84565b60008090505b60026040028110156114265761010781608081106113c957fe5b60020201600001548282608081106113dd57fe5b6020020151600001818152505061010781608081106113f857fe5b600202016001015482826080811061140c57fe5b6020020151602001818152505080806001019150506113af565b508091505090565b600061147a85858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c611654565b9050949350505050565b61148c611cee565b60018160006080811061149b57fe5b60200201818152505081816001608081106114b257fe5b6020020181815250506000600290505b600260400281101561151e576114fc6114f7848460018503608081106114e457fe5b602002015161161a90919063ffffffff16565b611654565b82826080811061150857fe5b60200201818152505080806001019150506114c2565b50919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905060008390506000811415611563576000925050506115e9565b818111156115785781818161157457fe5b0690505b600080600190506000849050600084905060005b600082146115c15781838161159d57fe5b0490508384820286038384840286038095508196508297508398505050505061158c565b60008512156115de578460000387039750505050505050506115e9565b849750505050505050505b919050565b6000817f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001039050919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808061164857fe5b83850991505092915050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161168357fe5b06915050919050565b611694611d11565b6001816000604081106116a357fe5b60200201818152505081816001604081106116ba57fe5b6020020181815250506000600290505b6040811015611723576117016116fc848460018503604081106116e957fe5b602002015161161a90919063ffffffff16565b611654565b82826040811061170d57fe5b60200201818152505080806001019150506116ca565b50919050565b611731611cd4565b6001821415611742578290506117c8565b600282141561175c5761175583846117ce565b90506117c8565b611764611d34565b83600001518160006003811061177657fe5b60200201818152505083602001518160016003811061179157fe5b60200201818152505082816002600381106117a857fe5b6020020181815250506040826060836007600019fa6117c657600080fd5b505b92915050565b6117d6611cd4565b6117de611d56565b8360000151816000600481106117f057fe5b60200201818152505083602001518160016004811061180b57fe5b60200201818152505082600001518160026004811061182657fe5b60200201818152505082602001518160036004811061184157fe5b6020020181815250506040826080836006600019fa61185f57600080fd5b5092915050565b6000808260006040811061187657fe5b602002015190506000600190505b60408110156118bf576118b084826040811061189c57fe5b602002015183611a0190919063ffffffff16565b91508080600101915050611884565b5080915050919050565b600080826000608081106118d957fe5b602002015190506000600190505b60026040028110156119255761191684826080811061190257fe5b602002015183611a0190919063ffffffff16565b915080806001019150506118e7565b5080915050919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508284101561196757838382030161196b565b8284035b91505092915050565b600080826001901b905060005b81851080156119905750600082115b156119a957600182901c91508080600101915050611981565b8084600101039250505092915050565b60008082905060008114156119d25760019150506119fc565b6000600190505b60008211156119f6576002810290508180600190039250506119d9565b80925050505b919050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190508080611a2f57fe5b83850891505092915050565b611a43611cd4565b611a4b611cd4565b611a8083600060808110611a5b57fe5b602002015185600060808110611a6d57fe5b602002015161172990919063ffffffff16565b90506000600190505b6002604002811015611ae957611ada611acb858360808110611aa757fe5b6020020151878460808110611ab857fe5b602002015161172990919063ffffffff16565b836117ce90919063ffffffff16565b91508080600101915050611a89565b508091505092915050565b604051806101000160405280611b08611d78565b8152602001611b15611d78565b8152602001611b22611d78565b8152602001611b2f611d78565b8152602001600081526020016000815260200160008152602001611b51611d92565b81525090565b60405180604001604052806002905b611b6e611d78565b815260200190600190039081611b665790505090565b6040518061100001604052806080905b611b9c611d78565b815260200190600190039081611b945790505090565b604051806103200160405280611bc6611dc6565b8152602001611bd3611dc6565b8152602001611be0611d78565b8152602001600081526020016000815260200160008152602001611c02611dea565b8152602001611c0f611dea565b815260200160008152602001600081526020016000815260200160008152602001611c38611e0d565b8152602001611c45611e30565b8152602001611c52611d78565b8152602001611c5f611d78565b815260200160008152602001611c73611e5e565b8152602001611c80611e5e565b8152602001611c8d611e5e565b8152602001611c9a611e5e565b8152602001611ca7611dea565b8152602001611cb4611dea565b8152602001611cc1611dea565b8152602001611cce611dea565b81525090565b604051806040016040528060008152602001600081525090565b604051806110000160405280608090602082028038833980820191505090505090565b604051806108000160405280604090602082028038833980820191505090505090565b6040518060600160405280600390602082028038833980820191505090505090565b6040518060800160405280600490602082028038833980820191505090505090565b604051806040016040528060008152602001600081525090565b6040518060800160405280611da5611e80565b8152602001611db2611e80565b815260200160008152602001600081525090565b60405180612000016040528061010090602082028038833980820191505090505090565b604051806110000160405280608090602082028038833980820191505090505090565b604051806108000160405280604090602082028038833980820191505090505090565b6040518061080001604052806040905b611e48611d78565b815260200190600190039081611e405790505090565b6040518060e00160405280600790602082028038833980820191505090505090565b604051806101c00160405280600e9060208202803883398082019150509050509056fea265627a7a7231582008a1e1e88b25cd6111bc40d77b24b4ec7ab09d6d7518af51942b5c6f34f4ec3d64736f6c63430005100032";

    public static final String FUNC_G = "g";

    public static final String FUNC_GVBASE = "gvBase";

    public static final String FUNC_H = "h";

    public static final String FUNC_HVBASE = "hvBase";

    public static final String FUNC_PARAMS = "params";

    public static final String FUNC_UBASE = "uBase";

    public static final String FUNC_AGGVERIFYRANGEPROOF = "aggVerifyRangeProof";

    @Deprecated
    protected AggRangeProofVerifier(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected AggRangeProofVerifier(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected AggRangeProofVerifier(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected AggRangeProofVerifier(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<AggRangeProofVerifier> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider, String params_) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params_)));
        return deployRemoteCall(AggRangeProofVerifier.class, web3j, credentials, contractGasProvider, BINARY, encodedConstructor);
    }

    public static RemoteCall<AggRangeProofVerifier> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider, String params_) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params_)));
        return deployRemoteCall(AggRangeProofVerifier.class, web3j, transactionManager, contractGasProvider, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<AggRangeProofVerifier> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String params_) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params_)));
        return deployRemoteCall(AggRangeProofVerifier.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<AggRangeProofVerifier> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String params_) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params_)));
        return deployRemoteCall(AggRangeProofVerifier.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> g() {
        final Function function = new Function(FUNC_G, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> gvBase(BigInteger param0) {
        final Function function = new Function(FUNC_GVBASE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(param0)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> h() {
        final Function function = new Function(FUNC_H, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> hvBase(BigInteger param0) {
        final Function function = new Function(FUNC_HVBASE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(param0)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteCall<String> params() {
        final Function function = new Function(FUNC_PARAMS, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<Tuple2<BigInteger, BigInteger>> uBase() {
        final Function function = new Function(FUNC_UBASE, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple2<BigInteger, BigInteger>>(
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(), 
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteCall<Boolean> aggVerifyRangeProof(List<BigInteger> points, List<BigInteger> scalar, List<BigInteger> l, List<BigInteger> r) {
        final Function function = new Function(FUNC_AGGVERIFYRANGEPROOF, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.StaticArray12<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(points, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.StaticArray5<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(scalar, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.StaticArray14<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(l, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.StaticArray14<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(r, org.web3j.abi.datatypes.generated.Uint256.class))), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    @Deprecated
    public static AggRangeProofVerifier load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new AggRangeProofVerifier(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static AggRangeProofVerifier load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new AggRangeProofVerifier(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static AggRangeProofVerifier load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new AggRangeProofVerifier(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static AggRangeProofVerifier load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new AggRangeProofVerifier(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
