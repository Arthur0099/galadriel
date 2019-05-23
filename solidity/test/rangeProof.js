const IPVerifier= artifacts.require("IPVerifier");
const PublicParams = artifacts.require("PublicParams");
const RangeProofVerifier = artifacts.require("RangeProofVerifier");

import {ReadJsonFile} from "./file.js";

contract("range proof", function(account){
  before(async function(){
    this.PublicParamsIns = await PublicParams.new();
    this.IPVerifierIns = await IPVerifier.new(this.PublicParamsIns.address);
    this.RangeProofVerifierIns = await RangeProofVerifier.new(this.PublicParamsIns.address, this.IPVerifierIns.address);
  });

  it("verify range proof", async function(){
    let path = "../proofs/rangeProof";
    let proofTest = ReadJsonFile(path);

    let rangeProof = proofTest.Proof;

    let points = [];
    points.push(rangeProof.A.X);
    points.push(rangeProof.A.Y);
    points.push(rangeProof.S.X);
    points.push(rangeProof.S.Y);
    points.push(rangeProof.T1.X);
    points.push(rangeProof.T1.Y);
    points.push(rangeProof.T2.X);
    points.push(rangeProof.T2.Y);
    points.push(proofTest.P.X);
    points.push(proofTest.P.Y);

    let scalar = [];
    scalar.push(rangeProof.t);
    scalar.push(rangeProof.tx);
    scalar.push(rangeProof.u);
    scalar.push(rangeProof.ipProof.a);
    scalar.push(rangeProof.ipProof.b);

    let l = [];
    for (let i = 0; i < rangeProof.ipProof.l.length; i++) {
      l.push(rangeProof.ipProof.l[i].X);
      l.push(rangeProof.ipProof.l[i].Y);
    }

    let r = [];
    for (let i = 0; i < rangeProof.ipProof.r.length; i++) {
      r.push(rangeProof.ipProof.r[i].X);
      r.push(rangeProof.ipProof.r[i].Y);
    }

    let res = await this.RangeProofVerifierIns.verifyRangeProof(points, scalar, l, r);
    assert.equal(res, true, "failed");
    let gasCost = await this.RangeProofVerifierIns.verifyRangeProof.estimateGas(points, scalar, l, r);
    console.log("when veriry is " + res + ", gas cost " + gasCost);
    //assert.equal(false, true, "t");
  });
});
