// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import '@openzeppelin/contracts/utils/Address.sol';

contract ERC721D is Ownable {
    uint256 public constant collectionSize = 1000;
    uint256 public constant price = 0.5 ether;

    event MintIncome(address indexed sender, uint indexed income);
    event RoyaltyIncome(address indexed sender, uint indexed income);  
    event FounderWithdraw(uint income);
    event CommunityWithdraw(uint income);

    struct DistributionRatio {
        uint256 founder;
        uint256 community;
    }

    DistributionRatio public mintRatio;
    DistributionRatio public royaltyRatio;

    uint256 public founderIncome;
    uint256 public communityIncome;

    function setRatio(uint256  mintToFounder, uint256  royaltyToFounder) external onlyOwner {
        require(mintToFounder <= 100,"must be <= 100");
        require(royaltyToFounder <= 100,"must be <= 100");
        mintRatio.founder = mintToFounder;
        mintRatio.community = 100 - mintRatio.founder;
        royaltyRatio.founder = royaltyToFounder;
        royaltyRatio.community = 100 - mintRatio.founder;
    }

    function _mintSettlement(uint256 value) internal {
        uint256 community = value * mintRatio.community / 100;
        uint256 founder = value - community;
        founderIncome += founder;
        communityIncome += community;
        emit MintIncome(msg.sender,value);
    }

    function founderWithdraw() external onlyOwner {
        Address.sendValue(payable(owner()), founderIncome);
        founderIncome = 0;
        emit FounderWithdraw(founderIncome);
    }
    function communityWithdraw() external onlyOwner {
        Address.sendValue(payable(owner()), communityIncome);
        communityIncome = 0;
        emit CommunityWithdraw(communityIncome);
    }
 
    fallback () external payable{
        _royaltySettlement();
    }

    receive () external payable{
        _royaltySettlement();
    }

    function _royaltySettlement() internal{
        if(msg.value <= 0){
            return;
        }
        uint256 community = msg.value * royaltyRatio.community / 100;
        uint256 founder = msg.value - community;
        founderIncome += founder;
        communityIncome += community;
        emit RoyaltyIncome(msg.sender,msg.value);
    }
}
