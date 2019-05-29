package contracts;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.Callable;
import org.web3j.abi.FunctionEncoder;
import org.web3j.abi.TypeReference;
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
public class SigmaVerifier extends Contract {
    private static final String BINARY = "0x608060405234801561001057600080fd5b50604051602080610d1c8339810180604052602081101561003057600080fd5b8101908080519060200190929190505050600081905061004e61021a565b8173ffffffffffffffffffffffffffffffffffffffff166304c09ce96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b1580156100af57600080fd5b505afa1580156100c3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525060408110156100e857600080fd5b810190809190505090506100fa61021a565b8273ffffffffffffffffffffffffffffffffffffffff166382529fdb6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401604080518083038186803b15801561015b57600080fd5b505afa15801561016f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250604081101561019457600080fd5b810190809190505090508160006002811015156101ad57fe5b602002015160008001819055508160016002811015156101c957fe5b60200201516000600101819055508060006002811015156101e657fe5b602002015160026000018190555080600160028110151561020357fe5b60200201516002600101819055505050505061023c565b6040805190810160405280600290602082028038833980820191505090505090565b610ad18061024b6000396000f3fe608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680636e26cc311461005c578063b8c9d36514610106578063e2179b8e14610138575b600080fd5b34801561006857600080fd5b506100ec60048036036102e081101561008057600080fd5b810190808061028001906014806020026040519081016040528092919082601460200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190803590602001909291908035906020019092919050505061016a565b604051808215151515815260200191505060405180910390f35b34801561011257600080fd5b5061011b6105f5565b604051808381526020018281526020019250505060405180910390f35b34801561014457600080fd5b5061014d610607565b604051808381526020018281526020019250505060405180910390f35b60008061021586600c60148110151561017f57fe5b602002015187600d60148110151561019357fe5b602002015188600e6014811015156101a757fe5b602002015189600f6014811015156101bb57fe5b60200201518a60106014811015156101cf57fe5b60200201518b60116014811015156101e357fe5b60200201518c60126014811015156101f757fe5b60200201518d601360148110151561020b57fe5b6020020151610619565b905061021f6109cf565b604080519081016040528088600060148110151561023957fe5b6020020151815260200188600160148110151561025257fe5b602002015181525081600060038110151561026957fe5b6020020181905250604080519081016040528088600c60148110151561028b57fe5b6020020151815260200188600d6014811015156102a457fe5b60200201518152508160016003811015156102bb57fe5b602002018190525060408051908101604052808860026014811015156102dd57fe5b602002015181526020018860036014811015156102f657fe5b602002015181525081600260038110151561030d57fe5b6020020181905250610320818784610694565b1515610331576000925050506105ed565b6103396109cf565b604080519081016040528089600660148110151561035357fe5b6020020151815260200189600760148110151561036c57fe5b602002015181525081600060038110151561038357fe5b6020020181905250604080519081016040528089600e6014811015156103a557fe5b6020020151815260200189600f6014811015156103be57fe5b60200201518152508160016003811015156103d557fe5b602002018190525060408051908101604052808960086014811015156103f757fe5b6020020151815260200189600960148110151561041057fe5b602002015181525081600260038110151561042757fe5b602002018190525061043a818785610694565b151561044c57600093505050506105ed565b6104546109fd565b60408051908101604052808a601060148110151561046e57fe5b602002015181526020018a601160148110151561048757fe5b602002015181525081600060028110151561049e57fe5b602002018190525060408051908101604052808a60046014811015156104c057fe5b602002015181526020018a60056014811015156104d957fe5b60200201518152508160016002811015156104f057fe5b602002018190525061050481898887610744565b15156105175760009450505050506105ed565b61051f6109fd565b60408051908101604052808b601260148110151561053957fe5b602002015181526020018b601360148110151561055257fe5b602002015181525081600060028110151561056957fe5b602002018190525060408051908101604052808b600a60148110151561058b57fe5b602002015181526020018b600b6014811015156105a457fe5b60200201518152508160016002811015156105bb57fe5b60200201819052506105cf81898988610744565b15156105e3576000955050505050506105ed565b6001955050505050505b949350505050565b60028060000154908060010154905082565b60008060000154908060010154905082565b6000610686898989898989898960405160200180898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001985050505050505050506040516020818303038152906040528051906020012060019004610846565b905098975050505050505050565b600061069e610a2b565b6106e98560016003811015156106b057fe5b60200201516106db858860026003811015156106c857fe5b602002015161088090919063ffffffff16565b61092d90919063ffffffff16565b90506106f3610a2b565b6107198587600060038110151561070657fe5b602002015161088090919063ffffffff16565b905080600001518260000151148015610739575080602001518260200151145b925050509392505050565b600061074e610a2b565b6107c56107848560006040805190810160405290816000820154815260200160018201548152505061088090919063ffffffff16565b6107b78760026040805190810160405290816000820154815260200160018201548152505061088090919063ffffffff16565b61092d90919063ffffffff16565b90506107cf610a2b565b61081a6107f8858960016002811015156107e557fe5b602002015161088090919063ffffffff16565b88600060028110151561080757fe5b602002015161092d90919063ffffffff16565b90508060000151826000015114801561083a575080602001518260200151145b92505050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381151561087757fe5b06915050919050565b610888610a2b565b600182141561089957829050610927565b60028214156108b3576108ac838461092d565b9050610927565b6108bb610a45565b83600001518160006003811015156108cf57fe5b60200201818152505083602001518160016003811015156108ec57fe5b6020020181815250508281600260038110151561090557fe5b6020020181815250506040826060836007600019fa151561092557600080fd5b505b92915050565b610935610a2b565b61093d610a68565b836000015181600060048110151561095157fe5b602002018181525050836020015181600160048110151561096e57fe5b602002018181525050826000015181600260048110151561098b57fe5b60200201818152505082602001518160036004811015156109a857fe5b6020020181815250506040826080836006600019fa15156109c857600080fd5b5092915050565b60c0604051908101604052806003905b6109e7610a8b565b8152602001906001900390816109df5790505090565b6080604051908101604052806002905b610a15610a8b565b815260200190600190039081610a0d5790505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b608060405190810160405280600490602082028038833980820191505090505090565b60408051908101604052806000815260200160008152509056fea165627a7a72305820a90b73380ce442a40e7021c99a7ebedcc8566c806e1888e19852a9c56b849c490029";

    public static final String FUNC_H = "h";

    public static final String FUNC_G = "g";

    public static final String FUNC_VERIFYSIGMAPROOF = "verifySigmaProof";

    @Deprecated
    protected SigmaVerifier(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected SigmaVerifier(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected SigmaVerifier(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected SigmaVerifier(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
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

    public static RemoteCall<SigmaVerifier> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider, String params) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params)));
        return deployRemoteCall(SigmaVerifier.class, web3j, credentials, contractGasProvider, BINARY, encodedConstructor);
    }

    public static RemoteCall<SigmaVerifier> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider, String params) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params)));
        return deployRemoteCall(SigmaVerifier.class, web3j, transactionManager, contractGasProvider, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<SigmaVerifier> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String params) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params)));
        return deployRemoteCall(SigmaVerifier.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<SigmaVerifier> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String params) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(params)));
        return deployRemoteCall(SigmaVerifier.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public RemoteCall<Boolean> verifySigmaProof(List<BigInteger> points, BigInteger z1, BigInteger z2, BigInteger z3) {
        final Function function = new Function(FUNC_VERIFYSIGMAPROOF, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.StaticArray20<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(points, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.Uint256(z1), 
                new org.web3j.abi.datatypes.generated.Uint256(z2), 
                new org.web3j.abi.datatypes.generated.Uint256(z3)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    @Deprecated
    public static SigmaVerifier load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new SigmaVerifier(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static SigmaVerifier load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new SigmaVerifier(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static SigmaVerifier load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new SigmaVerifier(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static SigmaVerifier load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new SigmaVerifier(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
