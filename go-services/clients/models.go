package clients

// FormDetailsResponse is an auto generated low-level Go binding around an user-defined struct.
type FormDetailsResponse struct {
	ResearchID      string `json:"research_id"`
	FormLink        string `json:"form_link"`
	SpreadSheetID   string `json:"spread_sheet_id"`
	SheetID         string `json:"sheet_id"`
	MaxDataSetCount string `json:"max_data_set_count"`
}

type LighthouseResponse struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

type StoreCIDResp struct {
	TxHash string `json:"tx_hash"`
}
