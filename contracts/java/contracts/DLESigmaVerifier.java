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
    private static final String BINARY = "0x608060405234801561001057600080fd5b506109c1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e26d0b81461003b5780639751cb1314610158575b600080fd5b61013e60048036036101c081101561005257600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f820116905080830192505050505050919291929080359060200190929190803590602001906401000000008111156100bb57600080fd5b8201836020820111156100cd57600080fd5b803590602001918460208302840111640100000000831117156100ef57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506101e1565b604051808215151515815260200191505060405180910390f35b6101c760048036036101a081101561016f57600080fd5b81019080806101800190600c806020026040519081016040528092919082600c60200280828437600081840152601f19601f82011690508083019250505050505091929192908035906020019092919050505061037a565b604051808215151515815260200191505060405180910390f35b6000606060046040519080825280602002602001820160405280156102155781602001602082028038833980820191505090505b50905060008090505b600481101561025d578581600c811061023357fe5b602002015182828151811061024457fe5b602002602001018181525050808060010191505061021e565b5060006102bf8260405160200180828051906020019060200280838360005b8381101561029757808201518184015260208101905061027c565b505050509050019150506040516020818303038152906040528051906020012060001c6103e3565b905060006103228560405160200180828051906020019060200280838360005b838110156102fa5780820151818401526020810190506102df565b505050509050019150506040516020818303038152906040528051906020012060001c6103e3565b90506000610360838360405160200180838152602001828152602001925050506040516020818303038152906040528051906020012060001c6103e3565b905061036d88888361041b565b9450505050509392505050565b6000806103cd846000600c811061038d57fe5b6020020151856001600c811061039f57fe5b6020020151866002600c81106103b157fe5b6020020151876003600c81106103c357fe5b602002015161064e565b90506103da84848361041b565b91505092915050565b6000807f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905080838161041257fe5b06915050919050565b600061042561090c565b846004600c811061043257fe5b60200201518160006006811061044457fe5b602002018181525050846005600c811061045a57fe5b60200201518160016006811061046c57fe5b602002018181525050846000600c811061048257fe5b60200201518160026006811061049457fe5b602002018181525050846001600c81106104aa57fe5b6020020151816003600681106104bc57fe5b602002018181525050846006600c81106104d257fe5b6020020151816004600681106104e457fe5b602002018181525050846007600c81106104fa57fe5b60200201518160056006811061050c57fe5b6020020181815250506105208185856106a4565b61052e576000915050610647565b61053661090c565b856008600c811061054357fe5b60200201518160006006811061055557fe5b602002018181525050856009600c811061056b57fe5b60200201518160016006811061057d57fe5b602002018181525050856002600c811061059357fe5b6020020151816002600681106105a557fe5b602002018181525050856003600c81106105bb57fe5b6020020151816003600681106105cd57fe5b60200201818152505085600a600c81106105e357fe5b6020020151816004600681106105f557fe5b60200201818152505085600b600c811061060b57fe5b60200201518160056006811061061d57fe5b6020020181815250506106318186866106a4565b61064057600092505050610647565b6001925050505b9392505050565b600061069a85858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c6103e3565b9050949350505050565b60006106ae61092e565b6040518060400160405280866000600681106106c657fe5b60200201518152602001866001600681106106dd57fe5b602002015181525090506106ef61092e565b60405180604001604052808760026006811061070757fe5b602002015181526020018760036006811061071e57fe5b6020020151815250905061073061092e565b60405180604001604052808860046006811061074857fe5b602002015181526020018860056006811061075f57fe5b6020020151815250905061077c86846107cf90919063ffffffff16565b92506107a38261079587846107cf90919063ffffffff16565b61087490919063ffffffff16565b9050806000015183600001511480156107c3575080602001518360200151145b93505050509392505050565b6107d761092e565b60018214156107e85782905061086e565b6002821415610802576107fb8384610874565b905061086e565b61080a610948565b83600001518160006003811061081c57fe5b60200201818152505083602001518160016003811061083757fe5b602002018181525050828160026003811061084e57fe5b6020020181815250506040826060836007600019fa61086c57600080fd5b505b92915050565b61087c61092e565b61088461096a565b83600001518160006004811061089657fe5b6020020181815250508360200151816001600481106108b157fe5b6020020181815250508260000151816002600481106108cc57fe5b6020020181815250508260200151816003600481106108e757fe5b6020020181815250506040826080836006600019fa61090557600080fd5b5092915050565b6040518060c00160405280600690602082028038833980820191505090505090565b604051806040016040528060008152602001600081525090565b6040518060600160405280600390602082028038833980820191505090505090565b604051806080016040528060049060208202803883398082019150509050509056fea265627a7a7231582010b50328b33cc1ebb1333bf4b609350c77fbe96383c9ff7ceb191082dcb9236f64736f6c63430005100032";

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

    public RemoteCall<Boolean> verifyDLESigmaProofWithCustom(List<BigInteger> points, BigInteger z, List<BigInteger> input) {
        final Function function = new Function(FUNC_VERIFYDLESIGMAPROOFWITHCUSTOM, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.StaticArray12<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(points, org.web3j.abi.datatypes.generated.Uint256.class)), 
                new org.web3j.abi.datatypes.generated.Uint256(z), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Uint256>(
                        org.web3j.abi.Utils.typeMap(input, org.web3j.abi.datatypes.generated.Uint256.class))), 
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
