package services

import (
	"encoding/json"
	"fmt"
)

type FormData struct {
	ResearchID    uint64
	SpreadSheetID string
	SheetID       uint64
}

type StoreCID struct {
	ResearchID uint64 `json:"research_id"`
	Name       string `json:"name"`
	CID        string `json:"cid"`
}

func (s *StoreCID) Marshal() ([]byte, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall storecid: %w", err)
	}

	return data, nil
}
