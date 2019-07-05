package contracts;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.concurrent.Callable;
import org.web3j.abi.EventEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Bool;
import org.web3j.abi.datatypes.Event;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.Utf8String;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameter;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.methods.request.EthFilter;
import org.web3j.protocol.core.methods.response.Log;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tuples.generated.Tuple3;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.ContractGasProvider;
import rx.Observable;
import rx.functions.Func1;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 3.6.0.
 */
public class TokenConverter extends Contract {
    private static final String BINARY = "0x60806040523480156200001157600080fd5b506200002c3362000160640100000000026401000000009004565b6200004633620001ca640100000000026401000000009004565b6200005062000447565b6040805190810160405280600581526020017f657468657200000000000000000000000000000000000000000000000000000081525081600001819052506001816020019015159081151581525050600060029050600081600a0a905080670de0b6b3a7640000811515620000c157fe5b0483604001818152505082600260008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190620001299291906200046b565b5060208201518160010160006101000a81548160ff021916908315150217905550604082015181600201559050505050506200051a565b620001848160006200023464010000000002620016a7179091906401000000009004565b8073ffffffffffffffffffffffffffffffffffffffff167f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129960405160405180910390a250565b620001ee8160016200023464010000000002620016a7179091906401000000009004565b8073ffffffffffffffffffffffffffffffffffffffff167fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f60405160405180910390a250565b6200024f828262000323640100000000026401000000009004565b151515620002c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515620003f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001807f526f6c65733a206163636f756e7420697320746865207a65726f20616464726581526020017f737300000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60606040519081016040528060608152602001600015158152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004ae57805160ff1916838001178555620004df565b82800160010185558215620004df579182015b82811115620004de578251825591602001919060010190620004c1565b5b509050620004ee9190620004f2565b5090565b6200051791905b8082111562000513576000816000905550600101620004f9565b5090565b90565b61197b806200052a6000396000f3fe6080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806310154bad146100b45780631d0b347a146101055780631f69565f1461020f578063219e62d4146102eb578063291d95491461035a5780633af32abf146103ab5780634c5a628c146104145780635c95c2d11461042b5780637362d9c81461049a578063bb5f747b146104eb578063d6cd947314610554575b600080fd5b3480156100c057600080fd5b50610103600480360360208110156100d757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061056b565b005b34801561011157600080fd5b506101f56004803603606081101561012857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460018302840111640100000000831117156101a357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061061a565b604051808215151515815260200191505060405180910390f35b34801561021b57600080fd5b5061025e6004803603602081101561023257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061098f565b604051808415151515815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156102ae578082015181840152602081019050610293565b50505050905090810190601f1680156102db5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156102f757600080fd5b506103446004803603604081101561030e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b13565b6040518082815260200191505060405180910390f35b34801561036657600080fd5b506103a96004803603602081101561037d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e62565b005b3480156103b757600080fd5b506103fa600480360360208110156103ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f11565b604051808215151515815260200191505060405180910390f35b34801561042057600080fd5b50610429610f2e565b005b34801561043757600080fd5b506104846004803603604081101561044e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f39565b6040518082815260200191505060405180910390f35b3480156104a657600080fd5b506104e9600480360360208110156104bd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111e5565b005b3480156104f757600080fd5b5061053a6004803603602081101561050e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611294565b604051808215151515815260200191505060405180910390f35b34801561056057600080fd5b506105696112b1565b005b61057433611294565b151561060e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260408152602001807f57686974656c69737441646d696e526f6c653a2063616c6c657220646f65732081526020017f6e6f742068617665207468652057686974656c69737441646d696e20726f6c6581525060400191505060405180910390fd5b610617816112bc565b50565b600061062533610f11565b15156106bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a8152602001807f57686974656c6973746564526f6c653a2063616c6c657220646f6573206e6f7481526020017f2068617665207468652057686974656c697374656420726f6c6500000000000081525060400191505060405180910390fd5b60008473ffffffffffffffffffffffffffffffffffffffff161415151561074e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f696e76616c696420746f6b656e2061646472657373000000000000000000000081525060200191505060405180910390fd5b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16151515610813576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f746f6b656e20616c72656164792061646465640000000000000000000000000081525060200191505060405180910390fd5b60008311151561088b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f696e76616c696420707265636973696f6e00000000000000000000000000000081525060200191505060405180910390fd5b6001600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff02191690831515021790555082600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018190555081600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000019080519060200190610983929190611886565b50600190509392505050565b600060606000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020154818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610aff5780601f10610ad457610100808354040283529160200191610aff565b820191906000526020600020905b815481529060010190602001808311610ae257829003601f168201915b505050505091509250925092509193909250565b6000610b1d611906565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060606040519081016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c015780601f10610bd657610100808354040283529160200191610c01565b820191906000526020600020905b815481529060010190602001808311610be457829003601f168201915b505050505081526020016001820160009054906101000a900460ff161515151581526020016002820154815250509050600083111515610ca9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f616d6f756e742063616e2774206265207a65726f00000000000000000000000081525060200191505060405180910390fd5b80602001511515610d22576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f746f6b656e206e6f7420737570706f72742063757272656e746c79000000000081525060200191505060405180910390fd5b6001816040015110151515610d9f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f746f6b656e277320707265636973696f6e206e6f74207365742072696768740081525060200191505060405180910390fd5b82610dcd8260400151610dbf84604001518761131690919063ffffffff16565b6113a990919063ffffffff16565b141515610e42576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f696e76616c696420616d6f756e7420707265636973696f6e000000000000000081525060200191505060405180910390fd5b610e5981604001518461131690919063ffffffff16565b91505092915050565b610e6b33611294565b1515610f05576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260408152602001807f57686974656c69737441646d696e526f6c653a2063616c6c657220646f65732081526020017f6e6f742068617665207468652057686974656c69737441646d696e20726f6c6581525060400191505060405180910390fd5b610f0e81611476565b50565b6000610f278260016114d090919063ffffffff16565b9050919050565b610f37336115f3565b565b6000610f43611906565b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060606040519081016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110275780601f10610ffc57610100808354040283529160200191611027565b820191906000526020600020905b81548152906001019060200180831161100a57829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815260200160028201548152505090506000831115156110cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f70676320616d6f756e742063616e2774206265207a65726f000000000000000081525060200191505060405180910390fd5b80602001511515611148576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f746f6b656e206e6f7420737570706f72742063757272656e746c79000000000081525060200191505060405180910390fd5b60018160400151101515156111c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f746f6b656e277320707265636973696f6e206e6f74207365742072696768740081525060200191505060405180910390fd5b6111dc8160400151846113a990919063ffffffff16565b91505092915050565b6111ee33611294565b1515611288576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260408152602001807f57686974656c69737441646d696e526f6c653a2063616c6c657220646f65732081526020017f6e6f742068617665207468652057686974656c69737441646d696e20726f6c6581525060400191505060405180910390fd5b6112918161164d565b50565b60006112aa8260006114d090919063ffffffff16565b9050919050565b6112ba33611476565b565b6112d08160016116a790919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f60405160405180910390a250565b6000808211151561138f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525060200191505060405180910390fd5b6000828481151561139c57fe5b0490508091505092915050565b6000808314156113bc5760009050611470565b600082840290508284828115156113cf57fe5b0414151561146b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f81526020017f770000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b809150505b92915050565b61148a81600161178490919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b660405160405180910390a250565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561159c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001807f526f6c65733a206163636f756e7420697320746865207a65726f20616464726581526020017f737300000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b61160781600061178490919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f0a8eb35e5ca14b3d6f28e4abf2f128dbab231a58b56e89beb5d636115001e16560405160405180910390a250565b6116618160006116a790919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f22380c05984257a1cb900161c713dd71d39e74820f1aea43bd3f1bdd2096129960405160405180910390a250565b6116b182826114d0565b151515611726576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b61178e82826114d0565b1515611828576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001807f526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c81526020017f650000000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106118c757805160ff19168380011785556118f5565b828001600101855582156118f5579182015b828111156118f45782518255916020019190600101906118d9565b5b509050611902919061192a565b5090565b60606040519081016040528060608152602001600015158152602001600081525090565b61194c91905b80821115611948576000816000905550600101611930565b5090565b9056fea165627a7a723058206733e2df2759d49632d2c465ce8d00ddfb5e0392a6638a53da1a67f82913d5e20029";

    public static final String FUNC_ADDWHITELISTED = "addWhitelisted";

    public static final String FUNC_REMOVEWHITELISTED = "removeWhitelisted";

    public static final String FUNC_ISWHITELISTED = "isWhitelisted";

    public static final String FUNC_RENOUNCEWHITELISTADMIN = "renounceWhitelistAdmin";

    public static final String FUNC_ADDWHITELISTADMIN = "addWhitelistAdmin";

    public static final String FUNC_ISWHITELISTADMIN = "isWhitelistAdmin";

    public static final String FUNC_RENOUNCEWHITELISTED = "renounceWhitelisted";

    public static final String FUNC_ADDTOKEN = "addToken";

    public static final String FUNC_CONVERTTOPGC = "convertToPGC";

    public static final String FUNC_CONVERTTOTOKEN = "convertToToken";

    public static final String FUNC_GETTOKENINFO = "getTokenInfo";

    public static final Event WHITELISTEDADDED_EVENT = new Event("WhitelistedAdded", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}));
    ;

    public static final Event WHITELISTEDREMOVED_EVENT = new Event("WhitelistedRemoved", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}));
    ;

    public static final Event WHITELISTADMINADDED_EVENT = new Event("WhitelistAdminAdded", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}));
    ;

    public static final Event WHITELISTADMINREMOVED_EVENT = new Event("WhitelistAdminRemoved", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}));
    ;

    @Deprecated
    protected TokenConverter(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected TokenConverter(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected TokenConverter(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected TokenConverter(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<TransactionReceipt> addWhitelisted(String account) {
        final Function function = new Function(
                FUNC_ADDWHITELISTED, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(account)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> removeWhitelisted(String account) {
        final Function function = new Function(
                FUNC_REMOVEWHITELISTED, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(account)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<Boolean> isWhitelisted(String account) {
        final Function function = new Function(FUNC_ISWHITELISTED, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(account)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteCall<TransactionReceipt> renounceWhitelistAdmin() {
        final Function function = new Function(
                FUNC_RENOUNCEWHITELISTADMIN, 
                Arrays.<Type>asList(), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> addWhitelistAdmin(String account) {
        final Function function = new Function(
                FUNC_ADDWHITELISTADMIN, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(account)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<Boolean> isWhitelistAdmin(String account) {
        final Function function = new Function(FUNC_ISWHITELISTADMIN, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(account)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteCall<TransactionReceipt> renounceWhitelisted() {
        final Function function = new Function(
                FUNC_RENOUNCEWHITELISTED, 
                Arrays.<Type>asList(), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public static RemoteCall<TokenConverter> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(TokenConverter.class, web3j, credentials, contractGasProvider, BINARY, "");
    }

    public static RemoteCall<TokenConverter> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(TokenConverter.class, web3j, transactionManager, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<TokenConverter> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(TokenConverter.class, web3j, credentials, gasPrice, gasLimit, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<TokenConverter> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(TokenConverter.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, "");
    }

    public List<WhitelistedAddedEventResponse> getWhitelistedAddedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(WHITELISTEDADDED_EVENT, transactionReceipt);
        ArrayList<WhitelistedAddedEventResponse> responses = new ArrayList<WhitelistedAddedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            WhitelistedAddedEventResponse typedResponse = new WhitelistedAddedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.account = (String) eventValues.getIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<WhitelistedAddedEventResponse> whitelistedAddedEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, WhitelistedAddedEventResponse>() {
            @Override
            public WhitelistedAddedEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(WHITELISTEDADDED_EVENT, log);
                WhitelistedAddedEventResponse typedResponse = new WhitelistedAddedEventResponse();
                typedResponse.log = log;
                typedResponse.account = (String) eventValues.getIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<WhitelistedAddedEventResponse> whitelistedAddedEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(WHITELISTEDADDED_EVENT));
        return whitelistedAddedEventObservable(filter);
    }

    public List<WhitelistedRemovedEventResponse> getWhitelistedRemovedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(WHITELISTEDREMOVED_EVENT, transactionReceipt);
        ArrayList<WhitelistedRemovedEventResponse> responses = new ArrayList<WhitelistedRemovedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            WhitelistedRemovedEventResponse typedResponse = new WhitelistedRemovedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.account = (String) eventValues.getIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<WhitelistedRemovedEventResponse> whitelistedRemovedEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, WhitelistedRemovedEventResponse>() {
            @Override
            public WhitelistedRemovedEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(WHITELISTEDREMOVED_EVENT, log);
                WhitelistedRemovedEventResponse typedResponse = new WhitelistedRemovedEventResponse();
                typedResponse.log = log;
                typedResponse.account = (String) eventValues.getIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<WhitelistedRemovedEventResponse> whitelistedRemovedEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(WHITELISTEDREMOVED_EVENT));
        return whitelistedRemovedEventObservable(filter);
    }

    public List<WhitelistAdminAddedEventResponse> getWhitelistAdminAddedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(WHITELISTADMINADDED_EVENT, transactionReceipt);
        ArrayList<WhitelistAdminAddedEventResponse> responses = new ArrayList<WhitelistAdminAddedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            WhitelistAdminAddedEventResponse typedResponse = new WhitelistAdminAddedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.account = (String) eventValues.getIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<WhitelistAdminAddedEventResponse> whitelistAdminAddedEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, WhitelistAdminAddedEventResponse>() {
            @Override
            public WhitelistAdminAddedEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(WHITELISTADMINADDED_EVENT, log);
                WhitelistAdminAddedEventResponse typedResponse = new WhitelistAdminAddedEventResponse();
                typedResponse.log = log;
                typedResponse.account = (String) eventValues.getIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<WhitelistAdminAddedEventResponse> whitelistAdminAddedEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(WHITELISTADMINADDED_EVENT));
        return whitelistAdminAddedEventObservable(filter);
    }

    public List<WhitelistAdminRemovedEventResponse> getWhitelistAdminRemovedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(WHITELISTADMINREMOVED_EVENT, transactionReceipt);
        ArrayList<WhitelistAdminRemovedEventResponse> responses = new ArrayList<WhitelistAdminRemovedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            WhitelistAdminRemovedEventResponse typedResponse = new WhitelistAdminRemovedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.account = (String) eventValues.getIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Observable<WhitelistAdminRemovedEventResponse> whitelistAdminRemovedEventObservable(EthFilter filter) {
        return web3j.ethLogObservable(filter).map(new Func1<Log, WhitelistAdminRemovedEventResponse>() {
            @Override
            public WhitelistAdminRemovedEventResponse call(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(WHITELISTADMINREMOVED_EVENT, log);
                WhitelistAdminRemovedEventResponse typedResponse = new WhitelistAdminRemovedEventResponse();
                typedResponse.log = log;
                typedResponse.account = (String) eventValues.getIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Observable<WhitelistAdminRemovedEventResponse> whitelistAdminRemovedEventObservable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(WHITELISTADMINREMOVED_EVENT));
        return whitelistAdminRemovedEventObservable(filter);
    }

    public RemoteCall<TransactionReceipt> addToken(String token, BigInteger precision, String name) {
        final Function function = new Function(
                FUNC_ADDTOKEN, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(token), 
                new org.web3j.abi.datatypes.generated.Uint256(precision), 
                new org.web3j.abi.datatypes.Utf8String(name)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<BigInteger> convertToPGC(String tokenAddr, BigInteger tokenAmount) {
        final Function function = new Function(FUNC_CONVERTTOPGC, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(tokenAddr), 
                new org.web3j.abi.datatypes.generated.Uint256(tokenAmount)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<BigInteger> convertToToken(String tokenAddr, BigInteger pgcAmount) {
        final Function function = new Function(FUNC_CONVERTTOTOKEN, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(tokenAddr), 
                new org.web3j.abi.datatypes.generated.Uint256(pgcAmount)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<Tuple3<Boolean, String, BigInteger>> getTokenInfo(String tokenAddr) {
        final Function function = new Function(FUNC_GETTOKENINFO, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(tokenAddr)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}, new TypeReference<Utf8String>() {}, new TypeReference<Uint256>() {}));
        return new RemoteCall<Tuple3<Boolean, String, BigInteger>>(
                new Callable<Tuple3<Boolean, String, BigInteger>>() {
                    @Override
                    public Tuple3<Boolean, String, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple3<Boolean, String, BigInteger>(
                                (Boolean) results.get(0).getValue(), 
                                (String) results.get(1).getValue(), 
                                (BigInteger) results.get(2).getValue());
                    }
                });
    }

    @Deprecated
    public static TokenConverter load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new TokenConverter(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static TokenConverter load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new TokenConverter(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static TokenConverter load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new TokenConverter(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static TokenConverter load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new TokenConverter(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static class WhitelistedAddedEventResponse {
        public Log log;

        public String account;
    }

    public static class WhitelistedRemovedEventResponse {
        public Log log;

        public String account;
    }

    public static class WhitelistAdminAddedEventResponse {
        public Log log;

        public String account;
    }

    public static class WhitelistAdminRemovedEventResponse {
        public Log log;

        public String account;
    }
}
