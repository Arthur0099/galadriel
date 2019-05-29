const PGC = artifacts.require("PGC");
const DLESigmaVerifier = artifacts.require("DLESigmaVerifier");
const PublicParams = artifacts.require("PublicParams");
const IPVerifier = artifacts.require("IPVerifier");
const RangeProofVerifier = artifacts.require("RangeProofVerifier");
const SigmaVerifier = artifacts.require("SigmaVerifier");
const {ether, balance, BN, constants, expectEvent, expectRevert} = require('openzeppelin-test-helpers');
const { expect } = require('chai');

import {ReadJsonFile} from "./file.js";

contract("pgc", function(account){
  before(async function(){
    this.PublicParamsIns = await PublicParams.new();
    this.DLESigmaVerifierIns = await DLESigmaVerifier.new();
    this.IPVerifierIns = await IPVerifier.new(this.PublicParamsIns.address);
    this.RangeProofVerifierIns = await RangeProofVerifier.new(this.PublicParamsIns.address, this.IPVerifierIns.address);
    this.SigmaVerifierIns = await SigmaVerifier.new(this.PublicParamsIns.address);
    this.PGCIns = await PGC.new(this.PublicParamsIns.address, this.DLESigmaVerifierIns.address, this.RangeProofVerifierIns.address, this.SigmaVerifierIns.address);
  });

  //
  it("depoist and burn", async function(){
    let path = "../proofs/depositBurn";
    let test = ReadJsonFile(path);

    let sender = account[0];
    assert.equal(2, test.deposit.length, "deposit length invalid");
    assert.equal(2, test.burn.length, "burn length invalid");

    // deposit alice.
    let accountAlice = [];
    let alice = test.deposit[0];
    accountAlice.push(alice.account.X);
    accountAlice.push(alice.account.Y);
    let aliceBalance = ether(alice.balance);
    await this.PGCIns.depositAccount(accountAlice, {from:sender, value: aliceBalance});

    // deposit bob.
    let accountBob = [];
    let bob = test.deposit[1];
    accountBob.push(bob.account.X);
    accountBob.push(bob.account.Y);
    let bobBalance = ether(bob.balance);
    await this.PGCIns.depositAccount(accountBob, {from: sender, value: bobBalance});

    // burn.
    let aliceReceiver = account[1];
    let proofAlice = [];
    let aliceBurn = test.burn[0];
    proofAlice.push(aliceBurn.proof.A1.X);
    proofAlice.push(aliceBurn.proof.A1.Y);
    proofAlice.push(aliceBurn.proof.A2.X);
    proofAlice.push(aliceBurn.proof.A2.Y);
    let aliceBalanceBefore = await balance.current(aliceReceiver);
    const burnAliceTx = await this.PGCIns.burn(aliceReceiver, alice.balance, accountAlice, proofAlice, aliceBurn.proof.Z);
    let aliceBalanceAfter = await balance.current(aliceReceiver);
    assert.equal(aliceBalanceAfter.sub(aliceBalanceBefore).eq(aliceBalance), true, "burn alice tx amount not correct");
    // make sure alice have no eth any more.
    await this.PGCIns.burn(aliceReceiver, alice.balance, accountAlice, proofAlice, aliceBurn.proof.Z);
    let aliceBalanceFinal = await balance.current(aliceReceiver);
    assert.equal(aliceBalanceFinal.sub(aliceBalanceAfter).eq(new BN(0)), true, "burn alice failed");

    // burn bob.
    let bobReceiver = account[2];
    let proofBob = [];
    let bobBurn = test.burn[1];
    proofBob.push(bobBurn.proof.A1.X);
    proofBob.push(bobBurn.proof.A1.Y);
    proofBob.push(bobBurn.proof.A2.X);
    proofBob.push(bobBurn.proof.A2.Y);
    let bobBalanceBefore = await balance.current(bobReceiver);
    await this.PGCIns.burn(bobReceiver, bob.balance, accountBob, proofBob, bobBurn.proof.Z);
    let bobBalanceAfter = await balance.current(bobReceiver);
    assert.equal(bobBalanceAfter.sub(bobBalanceBefore).eq(bobBalance), true, "burn bob tx amount not correct");
    // make sure bob have no eth any more.
    await this.PGCIns.burn(bobReceiver, bob.balance, accountBob, proofBob, bobBurn.proof.Z);
    let bobBalanceFinal = await balance.current(bobReceiver);
    assert.equal(bobBalanceFinal.sub(bobBalanceAfter).eq(new BN(0)), true, "burn bob failed");
  });
});
