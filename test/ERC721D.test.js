const { expect } = require("chai");
const { ethers } = require("hardhat");

const parseEther = ethers.utils.parseEther;

let owner;
let account;
let erc721D;


describe("ERC721D", function () {
  before(async function () {
    [owner, account] = await ethers.getSigners();
    const ERC721D = await ethers.getContractFactory("ERC721D");
    erc721D = await ERC721D.deploy();
    await erc721D.deployed();

    await erc721D.setRatio(70,50);
    console.log("erc721D mintRatio:",await erc721D.mintRatio())
    console.log("erc721D royaltyRatio:",await erc721D.royaltyRatio())
  });

  it("Should be able to deploy", async function () {});

  it("Should be able to mint and request a refund", async function () {

    console.log("owner balance:", (await owner.getBalance()).toString());
    console.log("erc721D balance:", (await erc721D.signer.getBalance()).toString());
    console.log("erc721D founderIncome:",await erc721D.founderIncome())
    console.log("erc721D communityIncome:",await erc721D.communityIncome())
    
    await account.sendTransaction({ to: erc721D.address, value: parseEther("1") });
    
    console.log("erc721D balance:", (await erc721D.signer.getBalance()).toString());
    console.log("erc721D founderIncome:",await erc721D.founderIncome())
    console.log("erc721D communityIncome:",await erc721D.communityIncome())

  });
});
