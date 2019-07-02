pragma solidity >=0.4.21 <0.6.0;

import "./library/SafeMath.sol";

contract TokenConverter {

  using SafeMath for uint;

  struct Token {
    string name;
    // added or not.
    bool added;
    // 1 pgc => amount of this token.
    // Warning: be careful with decimal with ether and token. i. e. 1 pgc == 1 ether means 1 gpc == 10**18.
    uint precision;
  }

  mapping(address => Token) tokens;

  constructor() public {
    Token memory etherToken;
    etherToken.name = "ether";
    etherToken.added = true;

    uint decimal = 2;
    uint rate = 10**decimal;
    etherToken.precision = 1 ether / rate;

    // set token address (0) as ether.
    tokens[address(0)] = etherToken;
  }

  /*
   *
   */
  function addToken(address token, uint precision, string memory name) public returns(bool) {
    require(uint(token) != 0, "invalid token address");
    require(!tokens[token].added, "token already added");
    require(precision > 0, "invalid precision");

    tokens[token].added = true;
    tokens[token].precision = precision;
    tokens[token].name = name;

    return true;
  }

  /*
   * @dev convert token amount to pgc amount.
   */
  function convertToPGC(address tokenAddr, uint tokenAmount) public view returns(uint) {
    Token memory token = tokens[tokenAddr];
    require(tokenAmount > 0, "amount can't be zero");
    require(token.added, "token not support currently");
    require(token.precision >= 1, "token's precision not set right");
    require(tokenAmount.div(token.precision).mul(token.precision) == tokenAmount, "invalid amount precision");

    return tokenAmount.div(token.precision);
  }

  /*
   * @dev convert pgc amount to token amount.
   */
  function convertToToken(address tokenAddr, uint pgcAmount) public view returns(uint) {
    Token memory token = tokens[tokenAddr];
    require(pgcAmount > 0, "pgc amount can't be zero");
    require(token.added, "token not support currently");
    require(token.precision >= 1, "token's precision not set right");

    return pgcAmount.mul(token.precision);
  }

  /*
   * @dev checks whether token is supported or not.
   */
  function isSupported(address tokenAddr) public view returns(bool) {
    return tokens[tokenAddr].added;
  }
}
