const DLESigmaVerifier = artifacts.require("DLESigmaVerifier.sol");
const {BN} = require("openzeppelin-test-helpers");

import {ReadJsonFile} from "./file.js";

contract("dle proof verifier", function(accounts) {
  before(async function(){
    this.ContractIns = await DLESigmaVerifier.new();
  });

  it("verify dle proof", async function(){
    let path = "/home/ubuntu/gopath/src/github.com/pgc/solidity/dlesigmaproof";
    let proof = ReadJsonFile(path);
    let points = [];
    points.push(proof.A1.X);
    points.push(proof.A1.Y);
    points.push(proof.A2.X);
    points.push(proof.A2.Y);
    points.push(proof.g1.X);
    points.push(proof.g1.Y);
    points.push(proof.h1.X);
    points.push(proof.h1.Y);
    points.push(proof.g2.X);
    points.push(proof.g2.Y);
    points.push(proof.h2.X);
    points.push(proof.h2.Y);
    let res = await this.ContractIns.verifyDLESigmaProof(points, proof.Z);
    assert.equal(res, true, "dlesigma proof verify failed");

    // log out gas cost.
    let gasCost = await this.ContractIns.verifyDLESigmaProof.estimateGas(points, proof.Z);
    console.log("gas cost " + gasCost);
  });
});
