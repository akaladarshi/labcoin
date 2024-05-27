class ResearchFormData {
    constructor(researchID, formLink, spreadSheetID, sheetID, maxDataSetCount) {
        this.researchID = researchID;
        this.formLink = formLink;
        this.spreadSheetID = spreadSheetID;
        this.sheetID = sheetID;
        this.maxDataSetCount = maxDataSetCount;
    }

    // Method to handle BigInt serialization
    static fromContractData(data) {
        return new ResearchFormData(
            data.researchID.toString(),
            data.formLink,
            data.spreadSheetID,
            data.sheetID.toString(),
            data.maxDataSetCount.toString()
        );
    }

    // Method to convert instance to a JSON-friendly format
    toJSON() {
        return {
            research_id: this.researchID,
            form_link: this.formLink,
            spread_sheet_id: this.spreadSheetID,
            sheet_id: this.sheetID,
            max_data_set_count: this.maxDataSetCount
        };
    }
}

class StoreCID {
    constructor(researchID, cid, name) {
        this.researchID = researchID;
        this.cid = cid;
        this.name = name
    }
}

module.exports = {
    ResearchFormData,
    StoreCID
};
