pragma solidity >= 0.5.0 < 0.6.0;
pragma experimental ABIEncoderV2;

import "./library/BN128.sol";

contract PublicParams {
  using BN128 for BN128.G1Point;
  using BN128 for uint;

  // g, h generator for sigma proof.
  // Warning: g is not the g base point of curve.
  BN128.G1Point public g;
  BN128.G1Point public h;

  // gBase represents curve base point on curve.
  BN128.G1Point public gBase;

  constructor() public {
    // set alt_128 g base point to g.
    h.X = 1;
    h.Y = 2;

    // set h generator.
    gBase.X = 1;
    gBase.Y = 2;

    g = generatePointByString("g generator of twisted elg");
  }

  // return g generator.
  function getG() public view returns(BN128.G1Point memory) {
    return g;
  }

  // return h generator.
  function getH() public view returns(BN128.G1Point memory) {
    return h;
  }

  // generatePointByString
  function generatePointByString(string memory s) internal view returns(BN128.G1Point memory) {
    uint scalar = uint(keccak256(abi.encodePacked(s))).mod();
    return gBase.mul(scalar);
  }
}
