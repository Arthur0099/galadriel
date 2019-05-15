const {BN} = require("openzeppelin-test-helpers");
const SigmaVerifier = artifacts.require("SigmaVerifier");
const PublicParams = artifacts.require("PublicParams");

import {ReadJsonFile} from "./file.js";

contract("sigma proof", function(account){
  before(async function(){
    this.PublicParamsIns = await PublicParams.new();
    this.SigmaVerifierIns = await SigmaVerifier.new(this.PublicParamsIns.address);
  });

  it("verify sigma proof", async function() {
    let path = "../proofs/sigmaProofs";
    let proofTestes = ReadJsonFile(path);

    for (var i = 0; i < proofTestes.length; i++) {
      let proofTest = proofTestes[i];
      let proof = proofTest.proof;
      let points = [];
      points.push(proof.pk1.X);
      points.push(proof.pk1.Y);
      points.push(proof.X1.X);
      points.push(proof.X1.Y);
      points.push(proof.Y1.X);
      points.push(proof.Y1.Y);
      points.push(proof.pk2.X);
      points.push(proof.pk2.Y);
      points.push(proof.X2.X);
      points.push(proof.X2.Y);
      points.push(proof.Y2.X);
      points.push(proof.Y2.Y);
      points.push(proof.A1.X);
      points.push(proof.A1.Y);
      points.push(proof.A2.X);
      points.push(proof.A2.Y);
      points.push(proof.B1.X);
      points.push(proof.B1.Y);
      points.push(proof.B2.X);
      points.push(proof.B2.Y);

      let res = await this.SigmaVerifierIns.verifySigmaProof(points, proof.z1, proof.z2, proof.z3);
      assert.equal(res, proofTest.expect, "sigma proof verify failed");

      // log out gas cost.
      let gasCost = await this.SigmaVerifierIns.verifySigmaProof.estimateGas(points, proof.z1, proof.z2, proof.z3);
      console.log("when verify result is " + res + ", gas cost " + gasCost);
    }
  });
});
