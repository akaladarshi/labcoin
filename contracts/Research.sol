// SPDX-License-Identifier: MIT
pragma solidity 0.8.21;

contract Research {
    struct FormData {
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
        FormData data;
    }

    uint256 public researchCounter;
    mapping(uint256 => ResearchData) public researches;

    // Event to log when a new research project is registered
    event ResearchRegistered(
        uint256 id,
        string title,
        string description,
        uint timeDuration,
        uint maxDataSetCount,
        string formLink,
        address indexed researcher
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
        uint _timeDuration
    ) public {
        researchCounter++;
        ResearchData memory _researchData = ResearchData(_title, _description, _timeDuration, msg.sender, FormData(_formLink, _spreadSheetID, _sheetID, _maxDataSetCount));
        verifyResearchData(_researchData);
        researches[researchCounter] = _researchData;

        emit ResearchRegistered(
            researchCounter,
            _title,
            _description,
            _timeDuration,
            _maxDataSetCount,
            _formLink,
            msg.sender
        );
    }


    function getAllFormsDetails() public view returns (FormData[] memory) {
        FormData[] memory _formsData = new FormData[](researchCounter);
        for (uint256 i = 1; i <= researchCounter; i++) {
            _formsData[i - 1] = researches[i].data;
        }

        return _formsData;
    }
}