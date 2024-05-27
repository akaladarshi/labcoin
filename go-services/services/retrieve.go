package services

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/akaladarshi/labcoin/config"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	retrieveFormDataURL = "https://docs.google.com/spreadsheets/d/%s/export?format=csv&gid=%d"
	tmpFilePath         = "../tmp/fom_data_%d.csv"
)

type RetrievalService struct {
	cfg        *config.Config
	counter    int
	FormDataCh <-chan *FormData
	APIService *sheets.Service // TODO: either remove this or find a way to retrieve using API
}

func NewRetrievalService(cfg *config.Config, dataCh <-chan *FormData) (*RetrievalService, error) {
	_, err := sheets.NewService(context.Background(), option.WithAPIKey(cfg.GoogleAPIKEY))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	return &RetrievalService{
		cfg:        cfg,
		FormDataCh: dataCh,
	}, nil
}

func (r *RetrievalService) Start(ctx context.Context) {
	for {
		select {
		case data := <-r.FormDataCh:
			url := fmt.Sprintf(retrieveFormDataURL, data.SpreadSheetID, data.SheetID)
			filePath := fmt.Sprintf(tmpFilePath, r.counter)
			err := downloadCSV(url, filePath)
			if err != nil {
				fmt.Printf("failed to download CSV: %v", err)
			}

			r.counter++
		case <-ctx.Done():
			return
		}
	}
}

// downloadCSV downloads a CSV file from the given URL and saves it to the specified file path
func downloadCSV(url, filepath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}

	return nil
}
