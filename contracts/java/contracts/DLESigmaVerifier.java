package contracts;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.List;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Bool;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
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
public class DLESigmaVerifier extends Contract {
    private static final String BINARY = "0x608060405234801561001057600080fd5b50610766806100206000396000f3fe608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680639751cb1314610046575b600080fd5b34801561005257600080fd5b506100c260048036036101a081101561006a57600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291905050506100dc565b604051808215151515815260200191505060405180910390f35b600080610137846000600c811015156100f157fe5b6020020151856001600c8110151561010557fe5b6020020151866002600c8110151561011957fe5b6020020151876003600c8110151561012d57fe5b60200201516103a0565b90506101416106b7565b846004600c8110151561015057fe5b602002015181600060068110151561016457fe5b602002018181525050846005600c8110151561017c57fe5b602002015181600160068110151561019057fe5b602002018181525050846000600c811015156101a857fe5b60200201518160026006811015156101bc57fe5b602002018181525050846001600c811015156101d457fe5b60200201518160036006811015156101e857fe5b602002018181525050846006600c8110151561020057fe5b602002015181600460068110151561021457fe5b602002018181525050846007600c8110151561022c57fe5b602002015181600560068110151561024057fe5b6020020181815250506102548185846103f7565b15156102655760009250505061039a565b61026d6106b7565b856008600c8110151561027c57fe5b602002015181600060068110151561029057fe5b602002018181525050856009600c811015156102a857fe5b60200201518160016006811015156102bc57fe5b602002018181525050856002600c811015156102d457fe5b60200201518160026006811015156102e857fe5b602002018181525050856003600c8110151561030057fe5b602002015181600360068110151561031457fe5b60200201818152505085600a600c8110151561032c57fe5b602002015181600460068110151561034057fe5b60200201818152505085600b600c8110151561035857fe5b602002015181600560068110151561036c57fe5b6020020181815250506103808186856103f7565b1515610392576000935050505061039a565b600193505050505b92915050565b60006103ed8585858560405160200180858152602001848152602001838152602001828152602001945050505050604051602081830303815290604052805190602001206001900461052e565b9050949350505050565b60006104016106da565b604080519081016040528086600060068110151561041b57fe5b6020020151815260200186600160068110151561043457fe5b602002015181525090506104466106da565b604080519081016040528087600260068110151561046057fe5b6020020151815260200187600360068110151561047957fe5b6020020151815250905061048b6106da565b60408051908101604052808860046006811015156104a557fe5b602002015181526020018860056006811015156104be57fe5b602002015181525090506104db868461056890919063ffffffff16565b9250610502826104f4878461056890919063ffffffff16565b61061590919063ffffffff16565b905080600001518360000151148015610522575080602001518360200151145b93505050509392505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381151561055f57fe5b06915050919050565b6105706106da565b60018214156105815782905061060f565b600282141561059b576105948384610615565b905061060f565b6105a36106f4565b83600001518160006003811015156105b757fe5b60200201818152505083602001518160016003811015156105d457fe5b602002018181525050828160026003811015156105ed57fe5b6020020181815250506040826060836007600019fa151561060d57600080fd5b505b92915050565b61061d6106da565b610625610717565b836000015181600060048110151561063957fe5b602002018181525050836020015181600160048110151561065657fe5b602002018181525050826000015181600260048110151561067357fe5b602002018181525050826020015181600360048110151561069057fe5b6020020181815250506040826080836006600019fa15156106b057600080fd5b5092915050565b60c060405190810160405280600690602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a72305820f99f7a39d97559b8e52e5204a6cb2ab7117035a4327da5f39aa3c8b5a77c92200029";

    public static final String FUNC_VERIFYDLESIGMAPROOF = "verifyDLESigmaProof";

    @Deprecated
    protected DLESigmaVerifier(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected DLESigmaVerifier(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected DLESigmaVerifier(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected DLESigmaVerifier(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Boolean> verifyDLESigmaProof(List<BigInteger> points, BigInteger z) {
        final Function function = new Function(FUNC_VERIFYDLESIGMAPROOF, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.StaticArray12<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(points, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.Uint256(z)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public static RemoteCall<DLESigmaVerifier> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(DLESigmaVerifier.class, web3j, credentials, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<DLESigmaVerifier> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(DLESigmaVerifier.class, web3j, credentials, gasPrice, gasLimit, BINARY, "");
    }

    public static RemoteCall<DLESigmaVerifier> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(DLESigmaVerifier.class, web3j, transactionManager, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<DLESigmaVerifier> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(DLESigmaVerifier.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, "");
    }

    @Deprecated
    public static DLESigmaVerifier load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new DLESigmaVerifier(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static DLESigmaVerifier load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new DLESigmaVerifier(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static DLESigmaVerifier load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new DLESigmaVerifier(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static DLESigmaVerifier load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new DLESigmaVerifier(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
