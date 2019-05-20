const DLESigmaVerifier = artifacts.require("DLESigmaVerifier.sol");
const {BN} = require("openzeppelin-test-helpers");

import {ReadJsonFile} from "./file.js";

contract("dle proof verifier", function(accounts) {
  before(async function(){
    this.ContractIns = await DLESigmaVerifier.new();
  });

  it("verify dle proof", async function(){
    let path = "../proofs/dleSigmaProof";
    let proofTests = ReadJsonFile(path);

    for (var i = 0; i < proofTests.length; i++) {
      const proofTest = proofTests[i];
      const proof = proofTest.proof;
      let points = [];
      points.push(proof.A1.X);
      points.push(proof.A1.Y);
      points.push(proof.A2.X);
      points.push(proof.A2.Y);
      points.push(proofTest.g1.X);
      points.push(proofTest.g1.Y);
      points.push(proofTest.h1.X);
      points.push(proofTest.h1.Y);
      points.push(proofTest.g2.X);
      points.push(proofTest.g2.Y);
      points.push(proofTest.h2.X);
      points.push(proofTest.h2.Y);
      let res = await this.ContractIns.verifyDLESigmaProof(points, proof.Z);
      assert.equal(res, proofTest.expect, "dlesigma proof verify failed");

      // log out gas cost.
      let gasCost = await this.ContractIns.verifyDLESigmaProof.estimateGas(points, proof.Z);
      console.log("when verify result is " + res + ", gas cost " + gasCost);
    }
  });
});
