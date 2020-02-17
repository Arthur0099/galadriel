pragma solidity >= 0.5.0 < 0.6.0;

import "./library/BN128.sol";
import "./DLESigmaVerifier.sol";
import "./PublicParams.sol";
import "./RangeProofVerifier.sol";
import "./SigmaVerifier.sol";
import "./AggRangeProofVerifier.sol";

// response for verifying all proofs.
contract PGCVerifier {
  using BN128 for BN128.G1Point;
  using BN128 for uint;

  // bitSize of balance value.
  uint public constant bitSize = 32;
  uint public constant n = 5;
  uint public constant maxNumber = 2**bitSize;

  // public h point.
  BN128.G1Point public h;

  // public g point.
  BN128.G1Point public g;

  // dle sigma verifier.
  DLESigmaVerifier public dleSigmaVerifier;

  // range proof verifier.
  RangeProofVerifier public rangeProofVerifier;

  // agg range proof verifier.
  AggRangeProofVerifier public aggRangeProofVerifier;

  // sigma verifier.
  SigmaVerifier public sigmaVerifier;

  // params.
  PublicParams public params;

  // encrypted ct.(balance of account)
  struct CT {
    BN128.G1Point X;
    BN128.G1Point Y;
    uint nonce;
  }

  struct Board {
    // for tmp calculation.
    uint i;
    uint[20] sigmaPoints;
    uint[16] ptePoints;
    uint[10] ctValidPoints;
    CT ct1;
    CT ct2;
    uint[4] ct1Points;
    uint[4] ct2Points;
    CT tmpUpdatedBalance;
    CT refreshBalance;
    BN128.G1Point[2] dleSigmaPoints;
    uint[12] dleTmpPoints;
    address token;
    uint tokenAmount;
    uint pgcAmount;

    // to call verify contract.
    uint[4] proof;
    uint[10] rpoints;
    uint[12] aggRpoints;
    uint[5] rscalar;
    uint[2*n] l;
    uint[2*n] r;
    uint[12] ll;
    uint[12] rr;

    // for log event.
    uint[4] fromto;
    uint[4] logCt1;
    uint[4] logCt2;
  }

  constructor(address params_, address dleSigmaVerifier_, address rangeProofVerifier_, address aggVerifier_, address sigmaVerifier_) public {
    params = PublicParams(params_);
    dleSigmaVerifier = DLESigmaVerifier(dleSigmaVerifier_);
    rangeProofVerifier = RangeProofVerifier(rangeProofVerifier_);
    sigmaVerifier = SigmaVerifier(sigmaVerifier_);
    aggRangeProofVerifier = AggRangeProofVerifier(aggVerifier_);

    uint[2] memory tmpH = params.getH();
    uint[2] memory tmpG = params.getG();

    h.X = tmpH[0];
    h.Y = tmpH[1];
    g.X = tmpG[0];
    g.Y = tmpG[1];
    require(bitSize == params.getBitSize(), "bitsize not equal");
  }

  //
  function verifyAggTransfer(uint[36] memory points, uint[10] memory scalar, uint[12] memory l, uint [12] memory r, uint[4] memory ub, uint nonce) public returns(bool) {
    CT memory userBalance;
    userBalance.X.X = ub[0];
    userBalance.X.Y = ub[1];
    userBalance.Y.X = ub[2];
    userBalance.Y.Y = ub[3];

    Board memory b;
    for (b.i = 0; b.i < 16; b.i++) {
      b.ptePoints[b.i] = points[b.i];
    }
    // verify pte proof.
    // pk1, pk2, ct enc. pte proof.
    // // 7 mul, 4 add.
    require(sigmaVerifier.verifyPTEProof(b.ptePoints, scalar[0], scalar[1]), "pte equal proof invalid");

    b.ct1.X = BN128.G1Point(points[4], points[5]);
    b.ct1.Y = BN128.G1Point(points[8], points[9]);
    // verify dle proof.
    // 2 add.
    // get tmp balance = alice'balnce - transfer'balance
    b.tmpUpdatedBalance.X = userBalance.X.add(b.ct1.X.neg());
    b.tmpUpdatedBalance.Y = userBalance.Y.add(b.ct1.Y.neg());
    b.refreshBalance.X = BN128.G1Point(points[16], points[17]);
    b.refreshBalance.Y = BN128.G1Point(points[18], points[19]);
    b.dleSigmaPoints[0] = BN128.G1Point(points[24], points[25]);
    b.dleSigmaPoints[1] = BN128.G1Point(points[26], points[27]);
    // only this failed.
    // 4 mul, 4 add.
    require(verifyDLESigmaProof(b.tmpUpdatedBalance, b.refreshBalance, b.dleSigmaPoints, points[0], points[1], scalar[4], 1, nonce), "dle sigma proof failed");

    // verifyCTValidProof.
    b.ctValidPoints[0] = points[0];
    b.ctValidPoints[1] = points[1];
    for (b.i = 16; b.i < 24; b.i++) {
      b.ctValidPoints[2+b.i-16] = points[b.i];
    }
    require(sigmaVerifier.verifyCTValidProof(b.ctValidPoints, scalar[2], scalar[3]), "ct valid proof invalid");

    for (b.i = 0; b.i < 8; b.i++) {
      b.aggRpoints[b.i] = points[28+b.i];
    }
    // set commitments.
    // transfer y.
    b.aggRpoints[8] = points[8];
    b.aggRpoints[9] = points[9];

    // pk1's refreshed balance.
    b.aggRpoints[10] = points[18];
    b.aggRpoints[11] = points[19];
    for (b.i = 0; b.i < 5; b.i++) {
      b.rscalar[b.i] = scalar[b.i+5];
    }
    // verify agg range proof.
    require(aggRangeProofVerifier.aggVerifyRangeProof(b.aggRpoints, b.rscalar, l, r), "aggrate range proof invalid");

    return true;
  }

  /*
   * @dev transfer from on account to another.
   * sigma proof to prove value in pk1's ct == value pk2's ct
   * dle sigma proof to prove value updated in pk1's ct == value in refreshed pk1's ct

   * points[0-1]: pk1
   * points[2-3]: ct1 X
   * points[4-5]: ct1 Y
   * points[6-7]: pk2
   * points[8-9]: ct2 X
   * points[10-11]: ct2 Y
   * points[12-13]: sigma proof A1.
   * points[14-15]: sigma proof A2.
   * points[16-17]: sigma proof B1.
   * points[18-19]: sigma proof B2.
   * points[20-21]: pk1's refreshed balance ct.X
   * points[22-23]: pk1's refreshed balance ct.Y
   * points[24-25]: dle sigma proof A1.
   * points[26-27]: dle sigma proof A2.

   * scalar[0-2]: sigma proof z1, z2, z3.
   * scalar[3]: dle sigma proof z.

   * scalar[4]: range proof 1 t.
   * scalar[5]: range proof 1 tx.
   * scalar[6]: range proof 1 u.
   * scalar[7]: range proof 1 a.
   * scalar[8]: range proof 1 b.

   * scalar[9]: range proof 2 t.
   * scalar[10]: range proof 2 tx.
   * scalar[11]: range proof 2 u.
   * scalar[12]: range proof 2 a.
   * scalar[13]: range proof 2 b.

   * range proof 1 to prove v in pk1's ct(points[4-5]) in range [0, 2^bitSize-1]
   * rpoints[0-1]: range proof 1 A.
   * rpoints[2-3]: range proof 1 S.
   * rpoints[4-5]: range proof 1 T1.
   * rpoints[6-7]: range proof 1 T2.

   * range proof 2 to prove v in refreshed balance (points[22-23]) in range [0, 2^bitSize-1]
   * rpoints[8-9]: range proof 2 A.
   * rpoints[10-11]: range proof 2 S.
   * rpoints[12-13]: range proof 2 T1.
   * rpoints[14-15]: range proof 2 T2.
   * l[0-2*n-1]: range proof 1 l.x, l.y.
   * r[0-2*n-1]: range proof 1 r.x, r.y.
   * l[2*n-4*n-1]: range proof 2 l.x, l.y.
   * r[2*n-4*n-1]: range proof 2 r.x, r.y.
   */
   // n=5, bitSize=32, 190 mul, 178 add.
  // function verifyTransfer(uint[28] memory points, uint[14] memory scalar, uint[16] memory rpoints, uint[4*n] memory l, uint[4*n] memory r, uint[4] memory ub) public returns(bool) {

  //   CT memory userBalance;
  //   userBalance.X.X = ub[0];
  //   userBalance.X.Y = ub[1];
  //   userBalance.Y.X = ub[2];
  //   userBalance.Y.Y = ub[3];

  //   Board memory b;
  //   // check v in ct1 == c in ct2.
  //   for (b.i = 0; b.i < 20; b.i++) {
  //     b.sigmaPoints[b.i] = points[b.i];
  //   }
  //   // 10 mul, 6 add.
  //   require(sigmaVerifier.verifySigmaProof(b.sigmaPoints, scalar[0], scalar[1], scalar[2]), "sigma verify failed");

  //   // check balance updated is same with refreshed balance.
  //   b.ct1.X = BN128.G1Point(points[2], points[3]);
  //   b.ct1.Y = BN128.G1Point(points[4], points[5]);
  //   for (b.i = 0; b.i < 4; b.i++) {
  //     b.ct1Points[b.i] = points[2+b.i];
  //   }
  //   // 2 add.
  //   // get tmp balance = alice'balnce - transfer'balance
  //   b.tmpUpdatedBalance.X = userBalance.X.add(b.ct1.X.neg());
  //   b.tmpUpdatedBalance.Y = userBalance.Y.add(b.ct1.Y.neg());
  //   b.refreshBalance.X = BN128.G1Point(points[20], points[21]);
  //   b.refreshBalance.Y = BN128.G1Point(points[22], points[23]);
  //   b.dleSigmaPoints[0] = BN128.G1Point(points[24], points[25]);
  //   b.dleSigmaPoints[1] = BN128.G1Point(points[26], points[27]);
  //   // 4 mul, 4 add.
  //   require(verifyDLESigmaProof(b.tmpUpdatedBalance, b.refreshBalance, b.dleSigmaPoints, points[0], points[1], scalar[3], 0, 0), "dle sigma proof failed");

  //   // check range proof 1.
  //   for (b.i = 0; b.i < 8; b.i++) {
  //     b.rpoints[b.i] = rpoints[b.i];
  //   }
  //   // set ct1.Y (commitment).
  //   b.rpoints[8] = points[4];
  //   b.rpoints[9] = points[5];
  //   for (b.i = 0; b.i < 2*n; b.i++) {
  //     b.l[b.i] = l[b.i];
  //     b.r[b.i] = r[b.i];
  //   }
  //   for (b.i = 0; b.i < 5; b.i++) {
  //     b.rscalar[b.i] = scalar[4+b.i];
  //   }
  //   // 2*(bitSize+n)+14 mul, 2*(bitSize+n)+9 add.
  //   require(rangeProofVerifier.optimizedVerifyRangeProof(b.rpoints, b.rscalar, b.l, b.r), "range proof 1 failed");

  //   // check range proof 2.
  //   for (b.i = 0; b.i < 8; b.i++) {
  //     b.rpoints[b.i] = rpoints[8+b.i];
  //   }
  //   b.rpoints[8] = points[22];
  //   b.rpoints[9] = points[23];
  //   for (b.i = 0; b.i < 2*n; b.i++) {
  //     b.l[b.i] = l[b.i+2*n];
  //     b.r[b.i] = r[b.i+2*n];
  //   }
  //   for (b.i = 0; b.i < 5; b.i++) {
  //     b.rscalar[b.i] = scalar[9+b.i];
  //   }
  //   // 2*(bitSize+n)+14 mul, 2*(bitSize+n)+9 add.
  //   require(rangeProofVerifier.optimizedVerifyRangeProof(b.rpoints, b.rscalar, b.l, b.r), "range proof 2 verify failed");

  //   return true;
  // }

  /*
   * @dev burn part amount of a pgc accunt.
   * dle sigma proof 1 to prove amout is same with value in ct and user holds the private key.
   * dle sigma proof 2 to prove user's updated balance is same with refreshBalance.
   * proof 1 to prove amout in range.
   * proof 2 to prove updated balance in range.

   * points[0-1]: pk
   * points[2-3]: ct.X
   * points[4-5]: ct.Y
   * points[6-7]: refreshed balance ct.X
   * points[8-9]: refreshed balance ct.Y
   * points[10-11]: dle sigma proof 1 A1
   * points[12-13]: dle sigma proof 1 A2
   * points[14-15]: dle sigma proof 2 A1
   * points[16-17]: dle sigma proof 2 A2
   *
   * scalar[0]: dle sigma proof 1 z.
   * scalar[1]: dle sigma proof 2 z.
   *
   * scalar[2]: range proof 1 t.
   * scalar[3]: range proof 1 tx.
   * scalar[4]: range proof 1 u.
   * scalar[5]: range proof 1 a.
   * scalar[6]: range proof 1 b.
   *
   * scalar[7]: range proof 2 t.
   * scalar[8]: range proof 2 tx.
   * scalar[9]: range proof 2 u.
   * scalar[10]: range proof 2 a.
   * scalar[11]: range proof 2 b.
   *
   * range proof 1 to prove v in ct(points[4-5]) in range [0, 2^bitSize-1]
   * rpoints[0-1]: range proof 1 A.
   * rpoints[2-3]: range proof 1 S.
   * rpoints[4-5]: range proof 1 T1.
   * rpoints[6-7]: range proof 1 T2.

   * range proof 2 to prove v in refreshed balance (points[8-9]) in range [0, 2^bitSize-1]
   * rpoints[8-9]: range proof 2 A.
   * rpoints[10-11]: range proof 2 S.
   * rpoints[12-13]: range proof 2 T1.
   * rpoints[14-15]: range proof 2 T2.
   * l[0-2*n-1]: range proof 1 l.x, l.y.
   * r[0-2*n-1]: range proof 1 r.x, r.y.
   * l[2*n-4*n-1]: range proof 2 l.x, l.y.
   * r[2*n-4*n-1]: range proof 2 r.x, r.y.
   */
  // function verifyBurnPart(uint amount, uint[18] memory points, uint[12] memory scalar, uint[16] memory rpoints, uint[4*n] memory l, uint[4*n] memory r, uint[4] memory ub) public returns(bool) {
  //   CT memory userBalance;
  //   userBalance.X.X = ub[0];
  //   userBalance.X.Y = ub[1];
  //   userBalance.Y.X = ub[2];
  //   userBalance.Y.Y = ub[3];

  //   Board memory b;
  //   b.ct1.X = BN128.G1Point(points[2], points[3]);
  //   b.ct1.Y = BN128.G1Point(points[4], points[5]);
  //   b.proof[0] = points[10];
  //   b.proof[1] = points[11];
  //   b.proof[2] = points[12];
  //   b.proof[3] = points[13];
  //   // check amount is same with value in ct.
  //   // require(verifyEqualProof(amount, b.ct1, BN128.G1Point(points[0], points[1]), b.proof, scalar[0], 0, 0, 0), "dle sigma proof 1 failed");

  //   // check balance updated is ame with refreshed balance.
  //   b.ct1.X = BN128.G1Point(points[2], points[3]);
  //   b.ct1.Y = BN128.G1Point(points[4], points[5]);
  //   // tmp balance = alice'balance - burn balance.
  //   b.tmpUpdatedBalance.X = userBalance.X.add(b.ct1.X.neg());
  //   b.tmpUpdatedBalance.Y = userBalance.Y.add(b.ct1.Y.neg());
  //   b.refreshBalance.X = BN128.G1Point(points[6], points[7]);
  //   b.refreshBalance.Y = BN128.G1Point(points[8], points[9]);
  //   b.dleSigmaPoints[0] = BN128.G1Point(points[14], points[15]);
  //   b.dleSigmaPoints[1] = BN128.G1Point(points[16], points[17]);
  //   require(verifyDLESigmaProof(b.tmpUpdatedBalance, b.refreshBalance, b.dleSigmaPoints, points[0], points[1], scalar[1], 0, 0), "dle sigma proof 2 failed");

  //   // check range proof 1.
  //   for (b.i = 0; b.i < 8; b.i++) {
  //     b.rpoints[b.i] = rpoints[b.i];
  //   }
  //   // set ct.y
  //   b.rpoints[8] = points[4];
  //   b.rpoints[9] = points[5];
  //   for (b.i = 0; b.i < 2*n; b.i++ ) {
  //     b.l[b.i] = l[b.i];
  //     b.r[b.i] = r[b.i];
  //   }
  //   for (b.i = 0; b.i < 5; b.i++ ) {
  //     b.rscalar[b.i] = scalar[b.i+2];
  //   }
  //   require(rangeProofVerifier.optimizedVerifyRangeProof(b.rpoints, b.rscalar, b.l, b.r), "range proof 1 failed");

  //   // check range proof 2.
  //   for (b.i = 0; b.i < 8; b.i++ ) {
  //     b.rpoints[b.i] = rpoints[8+b.i];
  //   }
  //   b.rpoints[8] = points[8];
  //   b.rpoints[9] = points[9];
  //   for (b.i = 0; b.i < 2*n; b.i++ ) {
  //     b.l[b.i] = l[b.i+2*n];
  //     b.r[b.i] = r[b.i+2*n];
  //   }
  //   for (b.i = 0; b.i < 5; b.i++ ) {
  //     b.rscalar[b.i] = scalar[7+b.i];
  //   }
  //   require(rangeProofVerifier.optimizedVerifyRangeProof(b.rpoints, b.rscalar, b.l, b.r), "range proof 2 verify failed");

  //   return true;
  // }

  function verifyBurn(uint amount, uint[2] memory publicKey, uint[4] memory proof, uint z, uint[4] memory ub, uint[] memory input) public returns(bool) {
    CT memory userBalance;
    userBalance.X.X = ub[0];
    userBalance.X.Y = ub[1];
    userBalance.Y.X = ub[2];
    userBalance.Y.Y = ub[3];

    // Board memory board;
    require(verifyEqualProof(amount, userBalance, BN128.G1Point(publicKey[0], publicKey[1]), proof, z, input), "dle sigma verify failed");

    return true;
  }

  // 4 mul, 4 add.
  function verifyDLESigmaProof(CT memory ori, CT memory refresh, BN128.G1Point[2] memory p, uint pkx, uint pky, uint z, uint t, uint nonce) internal view returns(bool) {
    BN128.G1Point memory g1 = refresh.Y.add(ori.Y.neg());
    BN128.G1Point memory h1 = refresh.X.add(ori.X.neg());
    uint[12] memory points;
    points[0] = p[0].X;
    points[1] = p[0].Y;
    points[2] = p[1].X;
    points[3] = p[1].Y;
    points[4] = g1.X;
    points[5] = g1.Y;
    points[6] = h1.X;
    points[7] = h1.Y;
    points[8] = h.X;
    points[9] = h.Y;
    points[10] = pkx;
    points[11] = pky;
    if (t == 0) {
        return dleSigmaVerifier.verifyDLESigmaProof(points, z);
    }
    if (t == 1) {
        return dleSigmaVerifier.verifyDLESigmaProofWithNonce(points, z, nonce);
    }
  }

   function verifyEqualProof(uint amount, CT memory ct, BN128.G1Point memory pk, uint[4] memory proof, uint z, uint[] memory input) internal view returns(bool) {
    BN128.G1Point memory y = ct.Y.add(g.mul(amount).neg());
    Board memory board;
    board.dleTmpPoints[0] = proof[0];
    board.dleTmpPoints[1] = proof[1];
    board.dleTmpPoints[2] = proof[2];
    board.dleTmpPoints[3] = proof[3];
    board.dleTmpPoints[4] = y.X;
    board.dleTmpPoints[5] = y.Y;
    board.dleTmpPoints[6] = ct.X.X;
    board.dleTmpPoints[7] = ct.X.Y;
    board.dleTmpPoints[8] = h.X;
    board.dleTmpPoints[9] = h.Y;
    board.dleTmpPoints[10] = pk.X;
    board.dleTmpPoints[11] = pk.Y;

    return dleSigmaVerifier.verifyDLESigmaProofWithCustom(board.dleTmpPoints, z, input);
  }
}
