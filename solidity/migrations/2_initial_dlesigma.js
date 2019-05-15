const DLESigmaVerifier = artifacts.require("DLESigmaVerifier");
const Secp256k1 = artifacts.require("Secp256k1");
const BN128 = artifacts.require("BN128");

module.exports = function(deployer) {
  deployer.deploy(Secp256k1);
  deployer.deploy(BN128);
  deployer.link(Secp256k1, DLESigmaVerifier);
  deployer.link(BN128, DLESigmaVerifier);
  deployer.deploy(DLESigmaVerifier);
};
