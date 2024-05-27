class ResearchFormData {
    constructor(formLink, spreadSheetID, sheetID, maxDataSetCount) {
        this.formLink = formLink;
        this.spreadSheetID = spreadSheetID;
        this.sheetID = sheetID;
        this.maxDataSetCount = maxDataSetCount;
    }

    // Method to handle BigInt serialization
    static fromContractData(data) {
        return new ResearchFormData(
            data[0],
            data[1],
            data[2].toString(), // Convert BigInt to string
            data[3].toString()  // Convert BigInt to string
        );
    }

    // Method to convert instance to a JSON-friendly format
    toJSON() {
        return {
            form_link: this.formLink,
            spread_sheet_id: this.spreadSheetID,
            sheet_id: this.sheetID,
            max_data_set_count: this.maxDataSetCount
        };
    }
}

module.exports = ResearchFormData;