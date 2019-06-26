pragma solidity >=0.4.21 <0.6.0;

contract TokenConverter {

  struct Token {
    // name of token. (optional)
    string name;
    // address of token contract.
    address tokenAddr;
    // 1 pgc => amount of this token.
    uint precision;
  }

  mapping(address => Token) tokens;

  function convertToPGC(uint tokenAmount, address tokenAddr) public view returns(uint) {
    Token memory token = tokens[tokenAddr];
    require(tokenAddr == token.tokenAddr, "invalid token addr");
    require(token.precision >= 1, "invalid precision of token");
    require(tokenAmount/token.precision*token.precision == tokenAmount, "invalid amount precision");

    return tokenAmount/token.precision;
  }

  function convertToToken(uint pgcAmount, address tokenAddr) public view returns(uint) {
    Token memory token = tokens[tokenAddr];
    require(tokenAddr == token.tokenAddr, "invalid token addr");
    require(token.precision >= 1, "invalid precision of token");

    return pgcAmount * token.precision;
  }
}
