// SPDX-License-Identifier: MIT
pragma solidity 0.8.21;

contract Research {
    struct FormData {
        string link;
        string dataLink;
    }

    struct ResearchData {
        string title;
        string description;
        uint timeDuration;
        uint dataSetCount;
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
        uint dataSetCount,
        address indexed researcher
    );

    function verifyResearchData(ResearchData memory _researchData) internal pure {
        require(bytes(_researchData.title).length > 0, "Title is required");
        require(bytes(_researchData.description).length > 0, "Description is required");
        require(_researchData.timeDuration > 0, "Time duration is required");
        require(_researchData.dataSetCount > 0, "Data set count is required");
        require(bytes(_researchData.data.link).length > 0, "Link is required");
        require(bytes(_researchData.data.dataLink).length > 0, "Data link is required");
    }

    // TODO: Integrate a token contract to charge a fee for registering a research project
    // TODO: Integrate a token contract to reward participant for participating in a research project.
    function registerResearch(
        string memory _title,
        string memory _description,
        string memory _link,
        string memory _dataLink,
        uint _dataSetCount,
        uint _timeDuration
    ) public {
        researchCounter++;
        ResearchData memory _researchData = ResearchData(_title, _description, _timeDuration, _dataSetCount, msg.sender, FormData(_link, _dataLink));
        verifyResearchData(_researchData);
        researches[researchCounter] = _researchData;

        emit ResearchRegistered(
            researchCounter,
            _title,
            _description,
            _timeDuration,
            _dataSetCount,
            msg.sender
        );
    }


    function getAllResearches() public view returns (FormData[] memory) {
        FormData[] memory _formsData = new FormData[](researchCounter);
        for (uint256 i = 1; i <= researchCounter; i++) {
            _formsData[i - 1] = researches[i].data;
        }

        return _formsData;
    }
}