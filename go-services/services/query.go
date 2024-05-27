package services

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/akaladarshi/labcoin/bindings"
	"github.com/akaladarshi/labcoin/clients"
	"github.com/akaladarshi/labcoin/config"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type QueryService struct {
	cfg             *config.Config
	researchBinding *bindings.Research
	contractClient  *clients.ResearchContractClient
	formDataCh      chan<- *FormData
	APIService      *sheets.Service
}

func NewQueryService(cfg *config.Config, researchBinding *bindings.Research, dataCh chan<- *FormData) (*QueryService, error) {
	srv, err := sheets.NewService(context.Background(), option.WithAPIKey(cfg.GoogleAPIKEY))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	client, err := clients.NewResearchClient(fmt.Sprintf(expressServerURL, cfg.ExpressServerPort))
	if err != nil {
		return nil, fmt.Errorf("failed to create research client: %v", err)
	}

	return &QueryService{
		cfg:             cfg,
		contractClient:  client,
		researchBinding: researchBinding,
		formDataCh:      dataCh,
		APIService:      srv,
	}, nil
}

func (q *QueryService) Start(ctx context.Context) {
	ticker := time.NewTicker(defaultQueryInterval)
	for {
		select {
		case <-ticker.C:
			// Query the contract
			err := q.queryContract(ctx)
			if err != nil {
				fmt.Printf("failed to query contract: %v", err) // TODO: use logger
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (q *QueryService) queryContract(_ context.Context) error {
	formDetails, err := q.contractClient.GetAllFormDetails()
	if err != nil {
		return err
	}

	for _, form := range formDetails {
		sheetID, err := strconv.ParseInt(form.SheetID, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse sheet ID: %v", err)
		}

		maxDataSetCount, err := strconv.ParseUint(form.MaxDataSetCount, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse max data set count: %v", err)
		}

		researchID, err := strconv.ParseUint(form.ResearchID, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse research ID: %v", err)
		}

		sheetName, err := getSheetName(form.SpreadSheetID, sheetID, q.cfg.GoogleAPIKEY)
		if err != nil {
			return fmt.Errorf("failed to get sheet name: %v", err)
		}

		spreadSheetRange := sheetName + "!A2:A"
		res, err := q.APIService.Spreadsheets.Values.Get(form.SpreadSheetID, spreadSheetRange).Do()
		if err != nil {
			return fmt.Errorf("failed to retrieve data from sheet: %v", err)
		}

		// Max data set count is reached, send the data to the retrieval service
		if uint64(len(res.Values)) >= maxDataSetCount {
			q.formDataCh <- &FormData{
				ResearchID:    researchID,
				SpreadSheetID: form.SpreadSheetID,
				SheetID:       uint64(sheetID),
				SheetName:     sheetName,
			}
		}
	}

	return err
}

// getSheetName retrieves the sheet name using the spreadsheet ID and sheet ID
func getSheetName(spreadsheetID string, sheetID int64, apiKey string) (string, error) {
	srv, err := sheets.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	spreadsheet, err := srv.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve spreadsheet data: %v", err)
	}

	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.SheetId == sheetID {
			return sheet.Properties.Title, nil
		}
	}

	return "", fmt.Errorf("sheet ID %d not found in spreadsheet", sheetID)
}
