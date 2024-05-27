package services

import (
	"context"
	"fmt"
	"time"

	"github.com/akaladarshi/labcoin/bindings"
	"github.com/akaladarshi/labcoin/config"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type QueryService struct {
	cfg             *config.Config
	researchBinding *bindings.Research
	formDataCh      chan<- *FormData
	APIService      *sheets.Service
}

func NewQueryService(cfg *config.Config, researchBinding *bindings.Research, dataCh chan<- *FormData) (*QueryService, error) {
	srv, err := sheets.NewService(context.Background(), option.WithAPIKey(cfg.GoogleAPIKEY))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	return &QueryService{
		cfg:             cfg,
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
	formDetails, err := q.researchBinding.GetAllFormsDetails(nil)
	if err != nil {
		return err
	}

	for _, form := range formDetails {
		sheetName, err := getSheetName(form.SpreadSheetID, form.SheetID.Int64(), q.cfg.GoogleAPIKEY)
		if err != nil {
			return fmt.Errorf("failed to get sheet name: %v", err)
		}

		spreadSheetRange := sheetName + "!A2:A"
		res, err := q.APIService.Spreadsheets.Values.Get(form.SpreadSheetID, spreadSheetRange).Do()
		if err != nil {
			return fmt.Errorf("failed to retrieve data from sheet: %v", err)
		}

		// Max data set count is reached, send the data to the retrieval service
		if uint64(len(res.Values)) >= form.MaxDataSetCount.Uint64() {
			q.formDataCh <- &FormData{
				SpreadSheetID: form.SpreadSheetID,
				SheetID:       form.SheetID.Uint64(),
			}
		}
	}

	return err
}

// getSheetName retrieves the sheet name using the spreadsheet ID and sheet ID
func getSheetName(spreadsheetID string, sheetID int64, apiKey string) (string, error) {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
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
