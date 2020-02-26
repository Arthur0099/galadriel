pragma solidity >= 0.5.0 < 0.6.0;

import "./library/BN128.sol";
import "./PublicParams.sol";
import "./TokenConverter.sol";
import "./PGCVerifier.sol";

/**
 * @dev Interface of the ERC20 standard as defined in the EIP. Does not include
 * the optional functions; to access them see `ERC20Detailed`.
 */
interface IERC20 {
    /**
     * @dev Moves `amount` tokens from the caller's account to `recipient`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a `Transfer` event.
     */
    function transfer(address recipient, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `sender` to `recipient` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a `Transfer` event.
     */
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}


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
    // for tmp calculation.
    CT ct1;
    CT ct2;

    CT tmpUpdatedBalance;
    address token;
    uint tokenAmount;
    uint pgcAmount;

    //
    uint[4] userBalance;
    uint localNonce;

    // for log event.
    uint[4] fromto;
    uint[4] logCt1;
    uint[4] logCt2;
    uint[6] aggCT;
  }

  // pk.x => pk.y => uint(address) => ct encrypt current balance.
  mapping(uint => mapping(uint => mapping(uint => CT))) currentBalance;
  // pk.x => pk.y => uint(address) => ct encrypt last balance.
  // balnce not roll over into current balance but can be spent.
  // nonce not set in this ct.
  mapping(uint => mapping(uint => mapping(uint => CT))) lastBalance;

  // pk.x => pk.y => uint(address) => ct encrypt pending balance.
  // balance in current epoch and can't be spent.
  mapping(uint => mapping(uint => mapping(uint => CT))) pendingBalance;

  // pk.x => pk.y => epochLength. number fo block one epoch contains.
  // epoch = block.number / epochLength;
  mapping(uint => mapping(uint => uint)) epochLength;

  // pk.x => pk.y => uint(address) => lastRolloverEpoch;
  // the epoch of last out tx of user in token.
  // currentEpoch = currentblock.number / epoch
  // epoch between currentEpoch and last rollover epoch is last epoch.
  mapping(uint => mapping(uint => mapping(uint => uint))) lastRolloverEpoch;

  // pk.x => pk.y => uint(address) => lastPendingEpoch;
  mapping(uint => mapping(uint => mapping(uint => uint))) lastPendingEpoch;

  // pk.x => pk.y => bool.
  // pending status open or not.
  mapping(uint => mapping(uint => bool)) pendingOpened;

  // pk.x => pk.y => uint.
  mapping(uint => mapping(uint => uint)) pendingFunNonce;

  // bitSize of balance value.
  uint public constant bitSize = 64;
  uint public constant n = 6;
  uint public constant maxNumber = 2**bitSize;
  // 2^step = aggSize
  uint constant step = 1;
  // proof l, r size.
  uint constant lrSize = n+step;

  // public h point.
  BN128.G1Point public h;

  // public g point.
  BN128.G1Point public g;

  // params.
  PublicParams public params;

  // token converter.
  TokenConverter public tokenConverter;

  // proof verifier.
  PGCVerifier public pgcVerifier;

  // events
  event LogDepositAccount(address indexed proxy, address indexed token, uint tox, uint toy, uint amount, uint time);
  event LogAggTransfer(address indexed proxy, address indexed token, uint[4] fromto, uint[6] ct, uint time);
  event LogTransfer(address indexed proxy, address indexed token, uint[4] fromto, uint[4] ct1, uint[4] ct2, uint time);
  event LogBurn(address indexed proxy, address indexed receiver, address indexed token, uint accountx, uint accounty, uint amount, uint time);
  event LogBurnPart(address indexed proxy, address indexed receiver, address indexed token, uint accountx, uint accounty, uint amount, uint time);

  // constructor.
  constructor(address params_, address pgcVerifier_, address tokenConverter_) public {
    pgcVerifier = PGCVerifier(pgcVerifier_);
    tokenConverter = TokenConverter(tokenConverter_);
    params = PublicParams(params_);

    uint[2] memory tmpH = params.getH();
    uint[2] memory tmpG = params.getG();

    h.X = tmpH[0];
    h.Y = tmpH[1];
    g.X = tmpG[0];
    g.Y = tmpG[1];
    require(bitSize == params.getBitSize(), "bitsize not equal");
  }

  // pgc account open the capacity of pending state.
  // the incoming tx will be set to pending state not to current state directly.
  function openPending(uint x, uint y, uint epochLength_, uint nonce, uint[2] memory sig) public {
    require(!pendingOpened[x][y], "pending already opened");
    uint localNonce = pendingFunNonce[x][y];
    require(nonce == localNonce, "invalid nonce");
    // verify sig.
    uint hash = uint(keccak256(abi.encodePacked(x, y, epochLength_, nonce))).mod();
    require(verifySig(hash, x, y, sig[0], sig[1]), "verify open pending sig failed");

    // epoch length can't below 5.
    if (epochLength_ < 5) {
      // set default epoch length value.
      epochLength_ = 50;
    }

    epochLength[x][y] = epochLength_;
    pendingOpened[x][y] = true;
    pendingFunNonce[x][y] = localNonce.add(1);
  }

  // pgc account close the capacity of pending state.
  function closePending(uint x, uint y, uint nonce, uint[2] memory sig) public {
    require(pendingOpened[x][y], "pending already closed");
    uint localNonce = pendingFunNonce[x][y];
    require(nonce == localNonce, "invalid nonce");
    // verify sig.
    uint hash = uint(keccak256(abi.encodePacked(x, y, nonce))).mod();
    require(verifySig(hash, x, y, sig[0], sig[1]), "verify close pending sig failed");

    epochLength[x][y] = 0;
    pendingOpened[x][y] = false;
    pendingFunNonce[x][y] = localNonce.add(1);
  }

  function depositAccountETH(uint[2] memory publicKey) public payable returns(bool) {
    depositAccount(publicKey, address(0), 0);
  }

  /*
   * @dev deposit eth to an account.
   * randomness is always set to zero. a depositer don't need to know the privatekey of pgc amount.
   */
  function depositAccount(uint[2] memory publicKey, address tokenAddr, uint tokenAmount) public payable returns(bool) {
    uint pgcAmount;
    if (uint(tokenAddr) == 0) {
       // use msg.value to calculate amount instead.
      pgcAmount = tokenConverter.convertToPGC(tokenAddr, msg.value);
    } else {
      require(msg.value == 0, "deposit token don't receive eth");

      // transfer token to this contract.
      require(IERC20(tokenAddr).transferFrom(msg.sender, address(this), tokenAmount), "token transfer failed");
      pgcAmount = tokenConverter.convertToPGC(tokenAddr, tokenAmount);
    }
    // encrypt amount and pending it.
    CT memory pb = encrypt(pgcAmount, BN128.G1Point(publicKey[0], publicKey[1]));
    toBalanceOrPending(pb, publicKey[0], publicKey[1], uint(tokenAddr));

    emit LogDepositAccount(msg.sender, tokenAddr, publicKey[0], publicKey[1], pgcAmount, now);

    return true;
  }

  function aggTransferETH(uint[36] memory points, uint[10] memory scalar, uint[2*lrSize] memory l, uint [2*lrSize] memory r) public returns (bool) {
      aggTransfer(points, scalar, 0, l, r);
  }

  /*
   * points[0-1]: pk1
   * points[2-3]: pk2
   * points[4-5]: ct x1
   * points[6-7]: ct x2
   * points[8-9]: ct y
   * points[10-11]: pte A1
   * points[12-13]: pte A2
   * points[14-15]: pte B
   * points[16-17]: pk1's refreshed balance ct.X
   * points[18-19]: pk1's refreshed balance ct.Y
   * points[20-21]: ct valid proof A
   * points[22-23]: ct valid proof B
   * points[24-25]: dle eq proof A1.
   * points[26-27]: dle eq proof A2.
   * points[28-29]: agg range proof A.
   * points[30-31]: agg range proof S.
   * points[32-33]: agg range proof T1.
   * points[34-35]: agg range proof T2.

   * scalar[0]: nonce.
   * scalar[0]: pte proof z1.
   * scalar[1]: pte proof z2.
   * scalar[2]: ct valid z1.
   * scalar[3]: ct valid z2. 
   * scalar[4]: dloge proof Z.
   * scalar[5]: agg range proof t.
   * scalar[6]: agg range proof tx.
   * scalar[7]: agg range proof u.
   * scalar[8]: agg range proof a.
   * scalar[9]: agg range proof b.

   * l[0-11]: agg range proof l.x, l.y.
   * r[0-11]: agg range proof r.x, r.y.
   */
  function aggTransfer(uint[36] memory points, uint[10] memory scalar, uint token, uint[2*lrSize] memory l, uint [2*lrSize] memory r) public returns (bool) {
    Board memory b;

    // nonce used for computing hash. no need to check.
    // 2 add if no pending.
    (b.userBalance, b.localNonce) = getUserBalance(points[0], points[1], address(token));
    // check proof.
    require(pgcVerifier.verifyAggTransfer(points, scalar, l, r, b.userBalance, b.localNonce), "transfer proofs invalid");

    // update sender's balance.
    // 4 add.
    // b.ct1.X = BN128.G1Point(points[4], points[5]);
    // b.ct1.Y = BN128.G1Point(points[8], points[9]);
    // b.tmpUpdatedBalance = getBalanceCanSpentInternal(points[0], points[1], token);
    // b.tmpUpdatedBalance.X = b.tmpUpdatedBalance.X.add(b.ct1.X.neg());
    // b.tmpUpdatedBalance.Y = b.tmpUpdatedBalance.Y.add(b.ct1.Y.neg());
    b.tmpUpdatedBalance.X = BN128.G1Point(points[16], points[17]);
    b.tmpUpdatedBalance.Y = BN128.G1Point(points[18], points[19]);
    b.tmpUpdatedBalance.nonce = b.localNonce;
    rolloverAndUpdate(b.tmpUpdatedBalance, points[0], points[1], token);

    // update receiver balance or make it to pending.
    b.ct2.X = BN128.G1Point(points[6], points[7]);
    b.ct2.Y = BN128.G1Point(points[8], points[9]);
    toBalanceOrPending(b.ct2, points[2], points[3], token);

    // set for event.
    b.fromto[0] = points[0];
    b.fromto[1] = points[1];
    b.fromto[2] = points[2];
    b.fromto[3] = points[3];
    b.aggCT[0] = points[4];
    b.aggCT[1] = points[5];
    b.aggCT[2] = points[6];
    b.aggCT[3] = points[7];
    b.aggCT[4] = points[8];
    b.aggCT[5] = points[9];
    emit LogAggTransfer(msg.sender, address(token), b.fromto, b.aggCT, now);

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
  // function transfer(uint[28] memory points, uint[14] memory scalar, uint[16] memory rpoints, uint[4*n] memory l, uint[4*n] memory r, uint token, uint nonce, uint[2] memory sig) public returns(bool) {
  //   // check for sig and nonce.
  //   // 4 mul, 1 add.
  //   require(verifyTransferSig(points, scalar, rpoints, l, r, token, nonce, sig), "verify sig failed for transfertx");

  //   Board memory b;
  //   // 2 add if no pending.
  //   (b.userBalance, b.localNonce) = getUserBalance(points[0], points[1], address(token));
  //   // check for nonce.
  //   require(nonce == b.localNonce, "invalid nonce");
  //   // n=5, bitSize=32, 190 mul, 178 add.
  //   require(pgcVerifier.verifyTransfer(points, scalar, rpoints, l, r, b.userBalance), "verify proofs for transfer failed");

  //   // update sender's balance.
  //   // 4 add.
  //   b.ct1.X = BN128.G1Point(points[2], points[3]);
  //   b.ct1.Y = BN128.G1Point(points[4], points[5]);
  //   // the refreshed balance is checked before, just use it instead.
  //   // b.tmpUpdatedBalance = getBalanceCanSpentInternal(points[0], points[1], token);
  //   // b.tmpUpdatedBalance.X = b.tmpUpdatedBalance.X.add(b.ct1.X.neg());
  //   // b.tmpUpdatedBalance.Y = b.tmpUpdatedBalance.Y.add(b.ct1.Y.neg());
  //   b.tmpUpdatedBalance.X = BN128.G1Point(points[20], points[21]);
  //   b.tmpUpdatedBalance.Y = BN128.G1Point(points[22], points[23]);
  //   rolloverAndUpdate(b.tmpUpdatedBalance, points[0], points[1], token);

  //   // update receiver balance or make it to pending.
  //   b.ct2.X = BN128.G1Point(points[8], points[9]);
  //   b.ct2.Y = BN128.G1Point(points[10], points[11]);
  //   toBalanceOrPending(b.ct2, points[6], points[7], token);

  //   // set for event.
  //   b.fromto[0] = points[0];
  //   b.fromto[1] = points[1];
  //   b.fromto[2] = points[6];
  //   b.fromto[3] = points[7];
  //   b.logCt1[0] = b.ct1.X.X;
  //   b.logCt1[1] = b.ct1.X.Y;
  //   b.logCt1[2] = b.ct1.Y.X;
  //   b.logCt1[3] = b.ct1.Y.Y;
  //   b.logCt2[0] = b.ct2.X.X;
  //   b.logCt2[1] = b.ct2.X.Y;
  //   b.logCt2[2] = b.ct2.Y.X;
  //   b.logCt2[3] = b.ct2.Y.Y;

  //   emit LogTransfer(msg.sender, address(token), b.fromto, b.logCt1, b.logCt2, now);

  //   return true;
  // }

  function getUserBalance(uint x, uint y, address token) public view returns (uint[4] memory ct, uint nonce) {
    CT memory userBalance = getBalanceCanSpentInternal(x, y, uint(token));
    ct[0] = userBalance.X.X;
    ct[1] = userBalance.X.Y;
    ct[2] = userBalance.Y.X;
    ct[3] = userBalance.Y.Y;
    nonce = userBalance.nonce;
    return (ct, nonce);
  }

  function getPendingFunNonce(uint x, uint y) public view returns(uint) {
    return pendingFunNonce[x][y];
  }

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
  // function burnPart(address payable receiver, uint token, uint amount, uint[18] memory points, uint[12] memory scalar, uint[16] memory rpoints, uint[4*n] memory l, uint[4*n] memory r, uint nonce, uint[2] memory sig) public returns(bool) {
  //   require(amount >= 1, "invalid amount");
  //   // check sig.
  //   require(verifyBurnPartSig(uint(receiver), token, amount, points, scalar, rpoints, l, r, nonce, sig), "verify sig failed for burn part");

  //   Board memory b;
  //   b.token = address(token);
  //   (b.userBalance, b.localNonce) = getUserBalance(points[0], points[1], b.token);
  //   require(nonce == b.localNonce, "invalid nonce");

  //   // check for proofs.
  //   require(pgcVerifier.verifyBurnPart(amount, points, scalar, rpoints, l, r, b.userBalance), "verify burnpart proofs failed");

  //   // calculate balance after transfer.
  //   b.ct1.X = BN128.G1Point(points[2], points[3]);
  //   b.ct1.Y = BN128.G1Point(points[4], points[5]);
  //   b.tmpUpdatedBalance = getBalanceCanSpentInternal(points[0], points[1], token);
  //   b.tmpUpdatedBalance.X = b.tmpUpdatedBalance.X.add(b.ct1.X.neg());
  //   b.tmpUpdatedBalance.Y = b.tmpUpdatedBalance.Y.add(b.ct1.Y.neg());

  //   // update sender's balance.
  //   rolloverAndUpdate(b.tmpUpdatedBalance, points[0], points[1], token);

  //   if (token == 0) {
  //     // transfer eth.
  //     receiver.transfer(tokenConverter.convertToToken(address(0), amount));
  //   } else {
  //     // transfer token back to user.
  //     b.tokenAmount = tokenConverter.convertToToken(b.token, amount);
  //     require(IERC20(b.token).transfer(receiver, b.tokenAmount), "transfer token back to user failed");
  //   }

  //   // emit event.
  //   emit LogBurnPart(msg.sender, receiver, b.token, points[0], points[1], amount, now);
  // }


  function burnETH(address payable receiver, uint amount, uint[2] memory publicKey, uint[4] memory proof, uint z) public returns(bool) {
    burn(receiver, 0, amount, publicKey, proof, z);
  }


  /*
   * @dev burn withdraw all eth back to eth account.
   * proof[0-1]: A1 point.
   * proof[2-3]: A2 point.
   *
   */
  function burn(address payable receiver, uint token, uint amount, uint[2] memory publicKey, uint[4] memory proof, uint z) public returns(bool) {
    // do nothing
    require(amount >= 1, "invalid amount");

    // check proof.
    uint[4] memory userBalance;
    uint nonceSent;
    (userBalance, nonceSent) = getUserBalance(publicKey[0], publicKey[1], address(token));
    // revert when error.

    uint[] memory input = new uint[](3);
    input[0] = nonceSent;
    input[1] = uint(receiver);
    input[2] = token;
    // receiver, nonce, token used for computing hash
    require(pgcVerifier.verifyBurn(amount, publicKey, proof, z, userBalance, input), "dle sigma verify failed");

    // update user encrypted balance.
    // CT memory tmpUpdatedBalance = encrypt(0, BN128.G1Point(publicKey[0], publicKey[1]));
    CT memory tmpUpdatedBalance;
    tmpUpdatedBalance.X = BN128.G1Point(0, 0);
    tmpUpdatedBalance.Y = BN128.G1Point(0, 0);
    tmpUpdatedBalance.nonce = nonceSent;
    rolloverAndUpdate(tmpUpdatedBalance, publicKey[0], publicKey[1], token);

    if (token == 0) {
      // transfer eth back to user.
      receiver.transfer(tokenConverter.convertToToken(address(token), amount));
    } else {
      require(IERC20(address(token)).transfer(receiver, tokenConverter.convertToToken(address(token), amount)), "transfer token back to user failed");
    }

    emit LogBurn(msg.sender, receiver, address(token), publicKey[0], publicKey[1], amount, now);
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
    // cause r is set to 0.
    ct.X = BN128.G1Point(0, 0);
    // Y = g*m + h*r.
    ct.Y = g.mul(amount);

    return ct;
  }

  /*
   *
   */
  function verifyBurnPartSig(uint receiver, uint token, uint amount, uint[18] memory points, uint[12] memory scalar, uint[16] memory rpoints, uint[4*n] memory l, uint[4*n] memory r, uint nonce, uint[2] memory sig) public view returns(bool) {
    uint hash = uint(keccak256(abi.encodePacked(receiver, token, amount, points, scalar, rpoints, l, r, nonce))).mod();

    return verifySig(hash, points[0], points[1], sig[0], sig[1]);
  }

  /*
   * @dev 4 mul, 1 add.
   */
  function verifySig(uint hash, uint pkx, uint pky, uint r, uint s) internal view returns(bool) {
    s = s.inv();
    uint u1 = hash.mul(s);
    uint u2 = r.mul(s);
    BN128.G1Point memory res = h.mul(u1).add(BN128.G1Point(pkx, pky).mul(u2));
    if (res.X != r) {
      return false;
    }
    return true;
  }

  /*
   * @dev return balance can be spent by public key and token address.
   */
  function getBalanceCanSpentInternal(uint x, uint y, uint token) internal view returns(CT memory) {
    // add currentBalance and lastBalance not matter whether pending status
    // open or not since lastBalance will be set to zero when closed except
    // the first out tx after close. but it's right to add the balance since it doesn't be spent.
    CT memory cb = currentBalance[x][y][token];
    CT memory lb = lastBalance[x][y][token];
    CT memory fb;
    fb.nonce = cb.nonce;
    fb.X = cb.X.add(lb.X);
    fb.Y = cb.Y.add(lb.Y);

    // first tx after close pending capacity.
    // todo:
    if (!pendingOpened[x][y] && pendingBalance[x][y][token].Y.X != 0) {
      // will be set to zero.
      fb.X = fb.X.add(pendingBalance[x][y][token].X);
      fb.Y = fb.Y.add(pendingBalance[x][y][token].Y);
    }

    // no in tx changes pending balance to last balance in epoch.
    if (pendingOpened[x][y] && pendingBalance[x][y][token].Y.X != 0 && block.number/epochLength[x][y] > lastPendingEpoch[x][y][token]) {
      fb.X = fb.X.add(pendingBalance[x][y][token].X);
      fb.Y = fb.Y.add(pendingBalance[x][y][token].Y);
    }

    return fb;
  }


  /*
   * @dev reduce balance of user's balance and rollover user's pending state.
   * @dev even epoch is current epoch, also rollover it.
   * @dev call before any out tx.
   */
  function rolloverAndUpdate(CT memory balance, uint x, uint y, uint token) internal {
    currentBalance[x][y][token] = balance;
    currentBalance[x][y][token].nonce = balance.nonce.add(1);

    // set last balance to zero. already be calculated.
    if (lastBalance[x][y][token].Y.X != 0) {
      lastBalance[x][y][token].X = BN128.G1Point(0, 0);
      lastBalance[x][y][token].Y = BN128.G1Point(0, 0);
    }

    // set pending balance to zero if pending not opened.
    if (!pendingOpened[x][y] && pendingBalance[x][y][token].Y.X != 0) {
      pendingBalance[x][y][token].X = BN128.G1Point(0, 0);
      pendingBalance[x][y][token].Y = BN128.G1Point(0, 0);
    }

    // first out tx in epoch.
    // no need to check lastPendingEpoch is zero or not. any in tx will update it. and without in tx, one just send any out tx.
    if (pendingOpened[x][y] && block.number/epochLength[x][y] > lastPendingEpoch[x][y][token]) {
      lastPendingEpoch[x][y][token] = block.number/epochLength[x][y];

      // this can be spent, so set it to zero.
      if (pendingBalance[x][y][token].Y.X != 0) {
        pendingBalance[x][y][token].X = BN128.G1Point(0, 0);
        pendingBalance[x][y][token].Y = BN128.G1Point(0, 0);
      }
    }
  }

  /*
   * @dev add balance to user' balance or pending.
   */
  function toBalanceOrPending(CT memory balance, uint x, uint y, uint token) internal {
    // not open pending capacity.
    if (!pendingOpened[x][y]) {
      CT storage cb = currentBalance[x][y][token];
      cb.X = cb.X.add(balance.X);
      cb.Y = cb.Y.add(balance.Y);

      // at this point, the last balance and pending balance may not zero, but it will be calculated when user burn/transfer.

      return;
    }

    CT storage lb = lastBalance[x][y][token];
    CT storage pb = pendingBalance[x][y][token];
    // try to pending this tx.
    uint currentEpoch = block.number / epochLength[x][y];
    // just pending ct to pending state.
    if (currentEpoch == lastPendingEpoch[x][y][token]) {
      pb.X = pb.X.add(balance.X);
      pb.Y = pb.Y.add(balance.Y);

      return;
    }

    // currentEpoch is bigger than last pending epoch. so just add pending balance to last balance and set pending balance to ct.
    // currentEpoch always >= lastpendingepoch.
    lb.X = lb.X.add(pb.X);
    lb.Y = lb.Y.add(pb.Y);
    // reset pending state.
    pb.X = balance.X;
    pb.Y = balance.Y;
    // reset last pending epoch.
    lastPendingEpoch[x][y][token] = currentEpoch;
  }
}
