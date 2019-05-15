const SigmaVerifier = artifacts.require("SigmaVerifier");
const BN128 = artifacts.require("BN128");
const PublicParams = artifacts.require("PublicParams");

module.exports = function(deployer) {
  deployer.deploy(BN128);
  deployer.link(BN128, PublicParams);
  deployer.deploy(PublicParams).then(function(){
    deployer.link(BN128, SigmaVerifier);
    deployer.deploy(SigmaVerifier, PublicParams.address);
  });
};
