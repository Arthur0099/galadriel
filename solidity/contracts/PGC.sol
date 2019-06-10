pragma solidity >= 0.5.0 < 0.6.0;

import "./library/BN128.sol";
import "./DLESigmaVerifier.sol";
import "./PublicParams.sol";
import "./RangeProofVerifier.sol";
import "./SigmaVerifier.sol";

contract PGC {

  using BN128 for BN128.G1Point;
  using BN128 for uint;

  // encrypted ct.(balance of account)
  struct CT {
    BN128.G1Point X;
    BN128.G1Point Y;
    uint nonce;
  }

  //
  struct Board {
    uint[20] sigmaPoints;

    //
    CT ct1;
    CT ct2;
    uint[4] ct1Points;
    uint[4] ct2Points;
    CT tmpUpdatedBalance;
    CT refreshBalance;
    BN128.G1Point[2] dleSigmaPoints;
    uint[12] dleTmpPoints;

    //
    uint[10] rpoints;
    uint[5] rscalar;
    uint[2*n] l;
    uint[2*n] r;
  }

  // pk.x => pk.y => encrypt balance.
  mapping(uint => mapping(uint => CT)) balance;

  // bitSize of balance value.
  uint public constant bitSize = 16;
  uint public constant n = 4;

  // public h point.
  BN128.G1Point public h;

  // public g point.
  BN128.G1Point public g;

  // dle sigma verifier.
  DLESigmaVerifier public dleSigmaVerifier;

  // range proof verifier.
  RangeProofVerifier public rangeProofVerifier;

  // sigma verifier.
  SigmaVerifier public sigmaVerifier;

  // params.
  PublicParams public params;

  // events
  event LogDepositAccount(address indexed proxy, uint tox, uint toy, uint amount, uint time);
  event LogTransfer(address indexed proxy, uint fromx, uint fromy, uint tox, uint toy, uint amountSenderXX, uint amountSenderXY, uint amountSenderYX, uint amountSenderYY, uint amountFromXX, uint amountFromXY, uint amountFromYX, uint amountFromYY, uint time);
  event LogBurn(address indexed proxy, address indexed receiver, uint accountx, uint accounty, uint amount, uint time);

  //
  constructor(address params_, address dleSigmaVerifier_, address rangeProofVerifier_, address sigmaVerifier_) public {
    params = PublicParams(params_);
    dleSigmaVerifier = DLESigmaVerifier(dleSigmaVerifier_);
    rangeProofVerifier = RangeProofVerifier(rangeProofVerifier_);
    sigmaVerifier = SigmaVerifier(sigmaVerifier_);

    uint[2] memory tmpH = params.getH();
    uint[2] memory tmpG = params.getG();

    h.X = tmpH[0];
    h.Y = tmpH[1];
    g.X = tmpG[0];
    g.Y = tmpG[1];
    require(bitSize == params.getBitSize(), "bitsize not equal");
  }

  /*
   * @dev deposit eth to an account.
   */
  function depositAccount(uint[2] memory publicKey) public payable returns(bool) {
    // check eth account.
    require(msg.value >= 1 ether, "eth deposited less than 1 eth");
    require(msg.value/1 ether*1 ether == msg.value, "eth amount not an integer");

    // add amount to user account.
    uint amount = msg.value/1 ether;

    CT storage userBalance = balance[publicKey[0]][publicKey[1]];
    CT memory deposited = encrypt(amount, BN128.G1Point(publicKey[0], publicKey[1]));
    userBalance.X = userBalance.X.add(deposited.X);
    userBalance.Y = userBalance.Y.add(deposited.Y);

    emit LogDepositAccount(msg.sender, publicKey[0], publicKey[1], msg.value, now);

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

   * range proof 1 to prove v in pk1's ct(points[2-3]) in range [0, 2^bitSize-1]
   * rpoints[0-1]: range proof 1 A.
   * rpoints[2-3]: range proof 1 S.
   * rpoints[4-5]: range proof 1 T1.
   * rpoints[6-7]: range proof 1 T2.

   * range proof 2 to prove v in refreshed balance (points[10-11]) in range [0, 2^bitSize-1]
   * rpoints[8-9]: range proof 2 A.
   * rpoints[10-11]: range proof 2 S.
   * rpoints[12-13]: range proof 2 T1.
   * rpoints[14-15]: range proof 2 T2.
   * l[0-2*n-1]: range proof 1 l.x, l.y.
   * r[0-2*n-1]: range proof 1 r.x, r.y.
   * l[2*n-4*n-1]: range proof 2 l.x, l.y.
   * r[2*n-4*n-1]: range proof 2 r.x, r.y.
   */
  function transfer(uint[28] memory points, uint[14] memory scalar, uint[16] memory rpoints, uint[4*n] memory l, uint[4*n] memory r, uint nonce, uint[2] memory sig) public returns(bool) {
    // check for sig and nonce.
    require(verifyTransferSig(points, scalar, rpoints, l, r, nonce, sig), "verify sig failed for transfertx");
    CT storage userBalance = balance[points[0]][points[1]];
    require(nonce == userBalance.nonce, "invalid nonce");
    Board memory b;
    // check v in ct1 == c in ct2.
    for (uint i = 0; i < 20; i++) {
      b.sigmaPoints[i] = points[i];
    }
    require(sigmaVerifier.verifySigmaProof(b.sigmaPoints, scalar[0], scalar[1], scalar[2]), "sigma verify failed");

    // check balance updated is same with refreshed balance.
    b.ct1.X = BN128.G1Point(points[2], points[3]);
    b.ct1.Y = BN128.G1Point(points[4], points[5]);
    for (uint i = 0; i < 4; i++) {
      b.ct1Points[i] = points[2+i];
    }
    // get tmp balance = alice'balnce - transfer'balance
    b.tmpUpdatedBalance.X = userBalance.X.add(b.ct1.X.neg());
    b.tmpUpdatedBalance.Y = userBalance.Y.add(b.ct1.Y.neg());
    b.refreshBalance.X = BN128.G1Point(points[20], points[21]);
    b.refreshBalance.Y = BN128.G1Point(points[22], points[23]);
    b.dleSigmaPoints[0] = BN128.G1Point(points[24], points[25]);
    b.dleSigmaPoints[1] = BN128.G1Point(points[26], points[27]);
    // only this failed.
    require(verifyDLESigmaProof(b.tmpUpdatedBalance, b.refreshBalance, b.dleSigmaPoints, BN128.G1Point(points[0], points[1]), scalar[3]), "dle sigma proof failed");

    // check range proof 1.
    for (uint i = 0; i < 8; i++) {
      b.rpoints[i] = rpoints[i];
    }
    // set ct1.Y (commitment).
    b.rpoints[8] = points[4];
    b.rpoints[9] = points[5];
    for (uint i = 0; i < 2*n; i++) {
      b.l[i] = l[i];
      b.r[i] = r[i];
    }
    for (uint i = 0; i < 5; i++) {
      b.rscalar[i] = scalar[4+i];
    }
    require(rangeProofVerifier.verifyRangeProof(b.rpoints, b.rscalar, b.l, b.r), "range proof 1 failed");

    // check range proof 2.
    for (uint i = 0; i < 8; i++) {
      b.rpoints[i] = rpoints[8+i];
    }
    b.rpoints[8] = points[22];
    b.rpoints[9] = points[23];
    for (uint i = 0; i < 2*n; i++) {
      b.l[i] = l[i+2*n];
      b.r[i] = r[i+2*n];
    }
    for (uint i = 0; i < 5; i++) {
      b.rscalar[i] = scalar[9+i];
    }
    require(rangeProofVerifier.verifyRangeProof(b.rpoints, b.rscalar, b.l, b.r), "range proof 2 verify failed");

    // update sender's balance.
    userBalance.X = b.tmpUpdatedBalance.X;
    userBalance.Y = b.tmpUpdatedBalance.Y;
    // update sender's nonce.
    userBalance.nonce = nonce + 1;

    // update receiver's balance.
    CT storage receiverBalance = balance[points[6]][points [7]];
    b.ct2.X = BN128.G1Point(points[8], points[9]);
    b.ct2.Y = BN128.G1Point(points[10], points[11]);
    for (uint i = 0; i < 4; i++) {
      b.ct2Points[i] = points[8+i];
    }
    receiverBalance.X = receiverBalance.X.add(b.ct2.X);
    receiverBalance.Y = receiverBalance.Y.add(b.ct2.Y);

    emit LogTransfer(msg.sender, points[0], points[1], points[6], points[7], b.ct1.X.X, b.ct1.X.Y, b.ct1.Y.X, b.ct1.Y.Y, b.ct2.X.X, b.ct2.X.Y, b.ct2.Y.X, b.ct2.Y.Y, now);

    return true;
  }

  function getUserBalance(uint x, uint y) public view returns (uint[5] memory ct) {
    CT memory userBalance = balance[x][y];
    ct[0] = userBalance.X.X;
    ct[1] = userBalance.X.Y;
    ct[2] = userBalance.Y.X;
    ct[3] = userBalance.Y.Y;
    ct[4] = userBalance.nonce;
    return ct;
  }

  function verifyDLESigmaProof(CT memory ori, CT memory refresh, BN128.G1Point[2] memory p, BN128.G1Point memory pk, uint z) internal view returns(bool) {
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
    points[10] = pk.X;
    points[11] = pk.Y;
    return dleSigmaVerifier.verifyDLESigmaProof(points, z);
  }


  /*
   * @dev burn withdraw all eth back to eth account.
   * proof[0-1]: A1 point.
   * proof[2-3]: A2 point.
   *
   */
  function burn(address payable receiver, uint amount, uint[2] memory publicKey, uint[4] memory proof, uint z, uint nonce, uint[2] memory sig) public returns(bool) {
    // check for sig.
    require(verifyBurnSig(uint(receiver), amount, publicKey, proof, z, nonce, sig), "invalid sig for burntx");
    // do nothing
    require(amount >= 1, "invalid amount");
    // compute y' = Y - g*m.
    CT storage userBalance = balance[publicKey[0]][publicKey[1]];
    require(nonce == userBalance.nonce, "invalid nonce");
    // todo: check not zero.

    BN128.G1Point memory y = userBalance.Y.add(g.mul(amount).neg());
    Board memory board;
    board.dleTmpPoints[0] = proof[0];
    board.dleTmpPoints[1] = proof[1];
    board.dleTmpPoints[2] = proof[2];
    board.dleTmpPoints[3] = proof[3];
    board.dleTmpPoints[4] = y.X;
    board.dleTmpPoints[5] = y.Y;
    board.dleTmpPoints[6] = userBalance.X.X;
    board.dleTmpPoints[7] = userBalance.X.Y;
    board.dleTmpPoints[8] = h.X;
    board.dleTmpPoints[9] = h.Y;
    board.dleTmpPoints[10] = publicKey[0];
    board.dleTmpPoints[11] = publicKey[1];

    // revert when error.
    require(dleSigmaVerifier.verifyDLESigmaProof(board.dleTmpPoints, z), "dle sigma verify failed");

    // update user encrypted balance.
    board.tmpUpdatedBalance = encrypt(0, BN128.G1Point(publicKey[0], publicKey[1]));
    userBalance.X = board.tmpUpdatedBalance.X;
    userBalance.Y = board.tmpUpdatedBalance.Y;
    // update user nonce.
    userBalance.nonce = nonce + 1;

    // transfer eth back to user.
    receiver.transfer(amount*1 ether);

    emit LogBurn(msg.sender, receiver, publicKey[0], publicKey[1], amount*1 ether, now);
  }

  /*
   * @dev encrypt for despoit eth account. the randomness is set to 0.
   * @dev it's ok since the eth amount is public know to everyone.
   * @dev but the ctx's randomness is not know.
   */
  function encrypt(uint amount, BN128.G1Point memory pk) internal view returns(CT memory) {
    require(amount < 2**bitSize, "amount out of range");

    // no sense.
    uint r = 0;

    CT memory ct;
    // X = pk * r.
    ct.X = pk.mul(r);
    // Y = g*m + h*r.
    ct.Y = g.mul(amount).add(h.mul(r));

    return ct;
  }

  /*
   *
   */
  function verifyTransferSig(uint[28] memory points, uint[14] memory scalar, uint[16] memory rpoints, uint[4*n] memory l, uint[4*n] memory r, uint nonce, uint[2] memory sig) internal returns(bool) {
    uint hash = uint(keccak256(abi.encodePacked(points, scalar, rpoints, l, r, nonce))).mod();
    return verifySig(hash, points[0], points[1], sig[0], sig[1]);
  }

  /**
   *
   */
  function verifyBurnSig(uint receiver, uint amount, uint[2] memory publicKey, uint[4] memory proof, uint z, uint nonce, uint[2] memory sig) public returns(bool) {
    uint hash = uint(keccak256(abi.encodePacked(receiver, amount, publicKey, proof, z, nonce))).mod();
    return verifySig(hash, publicKey[0], publicKey[1], sig[0], sig[1]);
  }

  /*
   *
   */
  function verifySig(uint hash, uint pkx, uint pky, uint r, uint s) internal returns(bool) {
    s = s.inv();
    uint u1 = hash.mul(s);
    uint u2 = r.mul(s);
    BN128.G1Point memory res = h.mul(u1).add(BN128.G1Point(pkx, pky).mul(u2));
    if (res.X != r) {
      return false;
    }
    return true;
  }
}
