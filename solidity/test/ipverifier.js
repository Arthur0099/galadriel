const IPVerifier = artifacts.require("IPVerifier");
const PublicParams = artifacts.require("PublicParams");

import {ReadJsonFile} from "./file.js";

contract("inner product", function(account){
  before(async function(){
    this.PublicParamsIns = await PublicParams.new();
    this.IPVerifierIns = await IPVerifier.new(this.PublicParamsIns.address);
  });

  it("verify inner product proof", async function(){
    let path = "../proofs/ipProof";
    let proofs = ReadJsonFile(path);

    // check proof.
    let p = [];
    p.push(proofs.p.X);
    p.push(proofs.p.Y);

    let l = [];
    let r = [];
    for (let i = 0; i < proofs.proof.l.length; i++) {
      l.push(proofs.proof.l[i].X);
      l.push(proofs.proof.l[i].Y);
      r.push(proofs.proof.r[i].X);
      r.push(proofs.proof.r[i].Y);
    }

    let scalar = [];
    scalar.push(proofs.proof.a);
    scalar.push(proofs.proof.b);
    scalar.push(proofs.c);

    let res = await this.IPVerifierIns.verifyIPProof(p, l, r, scalar);
    assert.equal(res, true, "failed");
    console.log("res " + res);
    let gasCost = await this.IPVerifierIns.verifyIPProof.estimateGas(p, l, r, scalar);
    console.log("when verify is " + res + ", gas cost " + gasCost);
  });
});
