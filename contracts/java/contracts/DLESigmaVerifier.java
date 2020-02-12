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
    private static final String BINARY = "0x608060405234801561001057600080fd5b50610a8b806100206000396000f3fe608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630581c0401461005c5780636907b66d146100fc5780639751cb13146101a6575b600080fd5b34801561006857600080fd5b506100e260048036036101c081101561008057600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f8201169050808301925050505050509192919290803590602001909291908035906020019092919050505061023c565b604051808215151515815260200191505060405180910390f35b34801561010857600080fd5b5061018c60048036036101e081101561012057600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919080359060200190929190803590602001909291905050506102af565b604051808215151515815260200191505060405180910390f35b3480156101b257600080fd5b5061022260048036036101a08110156101ca57600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190505050610324565b604051808215151515815260200191505060405180910390f35b600080610298856000600c8110151561025157fe5b6020020151866001600c8110151561026557fe5b6020020151876002600c8110151561027957fe5b6020020151886003600c8110151561028d57fe5b602002015187610395565b90506102a58585836103f5565b9150509392505050565b60008061030c866000600c811015156102c457fe5b6020020151876001600c811015156102d857fe5b6020020151886002600c811015156102ec57fe5b6020020151896003600c8110151561030057fe5b6020020151888861065c565b90506103198686836103f5565b915050949350505050565b60008061037f846000600c8110151561033957fe5b6020020151856001600c8110151561034d57fe5b6020020151866002600c8110151561036157fe5b6020020151876003600c8110151561037557fe5b60200201516106c5565b905061038c8484836103f5565b91505092915050565b60006103ea86868686866040516020018086815260200185815260200184815260200183815260200182815260200195505050505050604051602081830303815290604052805190602001206001900461071c565b905095945050505050565b60006103ff6109dc565b846004600c8110151561040e57fe5b602002015181600060068110151561042257fe5b602002018181525050846005600c8110151561043a57fe5b602002015181600160068110151561044e57fe5b602002018181525050846000600c8110151561046657fe5b602002015181600260068110151561047a57fe5b602002018181525050846001600c8110151561049257fe5b60200201518160036006811015156104a657fe5b602002018181525050846006600c811015156104be57fe5b60200201518160046006811015156104d257fe5b602002018181525050846007600c811015156104ea57fe5b60200201518160056006811015156104fe57fe5b602002018181525050610512818585610756565b1515610522576000915050610655565b61052a6109dc565b856008600c8110151561053957fe5b602002015181600060068110151561054d57fe5b602002018181525050856009600c8110151561056557fe5b602002015181600160068110151561057957fe5b602002018181525050856002600c8110151561059157fe5b60200201518160026006811015156105a557fe5b602002018181525050856003600c811015156105bd57fe5b60200201518160036006811015156105d157fe5b60200201818152505085600a600c811015156105e957fe5b60200201518160046006811015156105fd57fe5b60200201818152505085600b600c8110151561061557fe5b602002015181600560068110151561062957fe5b60200201818152505061063d818686610756565b151561064e57600092505050610655565b6001925050505b9392505050565b60006106b9878787878787604051602001808781526020018681526020018581526020018481526020018381526020018281526020019650505050505050604051602081830303815290604052805190602001206001900461071c565b90509695505050505050565b60006107128585858560405160200180858152602001848152602001838152602001828152602001945050505050604051602081830303815290604052805190602001206001900461071c565b9050949350505050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019050808381151561074d57fe5b06915050919050565b60006107606109ff565b604080519081016040528086600060068110151561077a57fe5b6020020151815260200186600160068110151561079357fe5b602002015181525090506107a56109ff565b60408051908101604052808760026006811015156107bf57fe5b602002015181526020018760036006811015156107d857fe5b602002015181525090506107ea6109ff565b604080519081016040528088600460068110151561080457fe5b6020020151815260200188600560068110151561081d57fe5b6020020151815250905061083a868461088d90919063ffffffff16565b925061086182610853878461088d90919063ffffffff16565b61093a90919063ffffffff16565b905080600001518360000151148015610881575080602001518360200151145b93505050509392505050565b6108956109ff565b60018214156108a657829050610934565b60028214156108c0576108b9838461093a565b9050610934565b6108c8610a19565b83600001518160006003811015156108dc57fe5b60200201818152505083602001518160016003811015156108f957fe5b6020020181815250508281600260038110151561091257fe5b6020020181815250506040826060836007600019fa151561093257600080fd5b505b92915050565b6109426109ff565b61094a610a3c565b836000015181600060048110151561095e57fe5b602002018181525050836020015181600160048110151561097b57fe5b602002018181525050826000015181600260048110151561099857fe5b60200201818152505082602001518160036004811015156109b557fe5b6020020181815250506040826080836006600019fa15156109d557600080fd5b5092915050565b60c060405190810160405280600690602082028038833980820191505090505090565b604080519081016040528060008152602001600081525090565b606060405190810160405280600390602082028038833980820191505090505090565b60806040519081016040528060049060208202803883398082019150509050509056fea165627a7a72305820ecab467d0b2c07b416ca832d0eb2414653e1a6cabe89b751bd9fcf2fd965724f0029";

    public static final String FUNC_VERIFYDLESIGMAPROOFWITHNONCE = "verifyDLESigmaProofWithNonce";

    public static final String FUNC_VERIFYDLESIGMAPROOFWITHCUSTOM = "verifyDLESigmaProofWithCustom";

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

    public RemoteCall<Boolean> verifyDLESigmaProofWithNonce(List<BigInteger> points, BigInteger z, BigInteger nonce) {
        final Function function = new Function(FUNC_VERIFYDLESIGMAPROOFWITHNONCE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.StaticArray12<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(points, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.Uint256(z), 
                new org.web3j.abi.datatypes.generated.Uint256(nonce)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteCall<Boolean> verifyDLESigmaProofWithCustom(List<BigInteger> points, BigInteger z, BigInteger nonce, BigInteger addr) {
        final Function function = new Function(FUNC_VERIFYDLESIGMAPROOFWITHCUSTOM, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.StaticArray12<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(points, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.Uint256(z), 
                new org.web3j.abi.datatypes.generated.Uint256(nonce), 
                new org.web3j.abi.datatypes.generated.Uint256(addr)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
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
