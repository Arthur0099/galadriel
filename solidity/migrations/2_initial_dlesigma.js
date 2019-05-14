const DLESigmaVerifier = artifacts.require("DLESigmaVerifier");
const Secp256k1 = artifacts.require("Secp256k1");

module.exports = function(deployer) {
  deployer.deploy(Secp256k1);
  deployer.link(Secp256k1, DLESigmaVerifier);
  deployer.deploy(DLESigmaVerifier);
};
