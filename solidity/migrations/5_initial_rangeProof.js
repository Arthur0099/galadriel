const BN128 = artifacts.require("BN128");
const PublicParams = artifacts.require("PublicParams");
const IPVerifier = artifacts.require("IPVerifier");
const RangeProofVerifier = artifacts.require("RangeProofVerifier");

module.exports = function(deployer) {
  deployer.deploy(BN128);
  deployer.link(BN128, PublicParams);

  deployer.deploy(PublicParams).then(function(){
    deployer.link(BN128, IPVerifier);
    deployer.deploy(IPVerifier).then(function(){
      deployer.link(BN128, RangeProofVerifier);
      deployer.deploy(RangeProofVerifier, PublicParams.address, IPVerifier.address).then(function(){
        console.log("over");
      });
    });
  });
};
