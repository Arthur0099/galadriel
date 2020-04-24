const IPVerifier = artifacts.require("IPVerifier");
const PublicParams = artifacts.require("PublicParams");

contract("inner product", function(account) {
  before(async function() {
    // this.PublicParamsIns = await PublicParams.new();
    console.log(IPVerifier);
    this.IPVerifierIns = await IPVerifier.new(
      "0x5A0b54D5dc17e0AadC383d2db43B0a0D3E029c4c"
    );
    console.log("ip address", this.IPVerifierIns.address);
  });

  it("verify inner product proof", async function() {});
});
