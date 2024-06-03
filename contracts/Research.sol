// SPDX-License-Identifier: MIT
pragma solidity 0.8.21;

import {Labcoin} from "./labcoin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Research {
    IERC20 public immutable labCoin;
    uint8 public constant DECIMALS = 18; // Load from labcoin

    struct FormData {
        uint256 researchID;
        string formLink;
        string spreadSheetID;
        uint sheetID;
        uint maxDataSetCount;
    }

    struct ResearchData {
        string title;
        string description;
        uint timeDuration;
        address researcher;
        uint256 remainingAmount;
        FormData data;
    }

    uint256 public researchCounter;
    mapping(uint256 => ResearchData) public researches;
    mapping (uint256 => string) public researchIDToCID;

    constructor(address _labCoinAddress) {
        labCoin = Labcoin(_labCoinAddress);
    }

    // Event to log when a new research project is registered
    event ResearchRegistered(
        uint256 id,
        string title,
        string description,
        uint timeDuration,
        uint maxDataSetCount,
        string formLink,
        address indexed researcher,
        uint256 remainingAmount
    );

    function verifyResearchData(ResearchData memory _researchData) internal pure {
        require(bytes(_researchData.title).length > 0, "Title is required");
        require(bytes(_researchData.description).length > 0, "Description is required");
        require(_researchData.timeDuration > 0, "Time duration is required");
        require(_researchData.data.maxDataSetCount > 0, "Max Data set count is required");
        require(bytes(_researchData.data.formLink).length > 0, "Form Link is required");
        require(_researchData.data.sheetID > 0, "Sheet ID is required");
        require(bytes(_researchData.data.spreadSheetID).length > 0, "Spread sheet id is required");
    }

    // TODO: Integrate a token contract to charge a fee for registering a research project
    // TODO: Integrate a token contract to reward participant for participating in a research project.
    function registerResearch(
        string memory _title,
        string memory _description,
        string memory _formLink,
        string memory _spreadSheetID,
        uint _sheetID,
        uint _maxDataSetCount,
        uint _timeDuration,
        uint _distributionAmoutInWhole
    ) public {
        require(_distributionAmoutInWhole > 0, "Amount is required");
        uint amountInSmallestUnits = _distributionAmoutInWhole * (10 ** DECIMALS);

        // Ensure the contract is allowed to spend the specified amount on behalf of the researcher
        require(labCoin.allowance(msg.sender, address(this)) >= amountInSmallestUnits, "Allowance is less than amount");

        // Transfer the specified amount of LabPoints from the researcher to the contract
        require(labCoin.transferFrom(msg.sender, address(this), amountInSmallestUnits), "Token transfer failed");

        researchCounter++;
        ResearchData memory _researchData = ResearchData(_title, _description, _timeDuration, msg.sender, amountInSmallestUnits, FormData(researchCounter, _formLink, _spreadSheetID, _sheetID, _maxDataSetCount));
        verifyResearchData(_researchData);
        researches[researchCounter] = _researchData;

        emit ResearchRegistered(
            researchCounter,
            _title,
            _description,
            _timeDuration,
            _maxDataSetCount,
            _formLink,
            msg.sender,
            _distributionAmoutInWhole
        );
    }

    function distributeLabCoin(uint256 _researchId, address participant) public {
        ResearchData storage research = researches[_researchId];

        // distribute amount based on max data set
        uint256 amountPerParticipant = research.remainingAmount / research.data.maxDataSetCount;

        // Transfer labcoins on behalf on the researcher to the participant
        require(labCoin.transfer(participant, amountPerParticipant), "Token transfer failed");

        // Update the remaining amount to be distributed
        research.remainingAmount -= amountPerParticipant;
    }


    function getAllFormsDetails() public view returns (FormData[] memory) {
        uint256 count = 0;
        for (uint256 i = 1; i <= researchCounter; i++) {
            if (bytes(researchIDToCID[i]).length < 1) {
                count++;
            }
        }

        FormData[] memory _formsData = new FormData[](count);
        uint256 index = 0;
        for (uint256 i = 1; i <= researchCounter; i++) {
            if (bytes(researchIDToCID[i]).length < 1) {
                _formsData[index] = researches[i].data;
                index++;
            }
        }

        return _formsData;
    }

    function storeCID(uint256 _researchId, string memory _cid) public {
        require(bytes(_cid).length > 0, "CID is required");
        require(bytes(researches[_researchId].title).length > 0, "Research project does not exist");

        researchIDToCID[_researchId] = _cid;
    }
}