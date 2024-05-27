package services

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akaladarshi/labcoin/clients"
	"github.com/akaladarshi/labcoin/config"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type RetrievalService struct {
	cfg        *config.Config
	counter    int
	client     *clients.LighthouseClient
	formDataCh <-chan *FormData
	storeCIDCh chan<- *StoreCID
}

func NewRetrievalService(cfg *config.Config, dataCh <-chan *FormData, storeCID chan<- *StoreCID) (*RetrievalService, error) {
	_, err := sheets.NewService(context.Background(), option.WithAPIKey(cfg.GoogleAPIKEY))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	lighthouseClient := clients.NewLighthouseClient(cfg)
	return &RetrievalService{
		cfg:        cfg,
		counter:    0,
		client:     lighthouseClient,
		formDataCh: dataCh,
		storeCIDCh: storeCID,
	}, nil
}

func (r *RetrievalService) Start(ctx context.Context) {
	for {
		select {
		case data := <-r.formDataCh:
			url := fmt.Sprintf(retrieveFormDataURL, data.SpreadSheetID, data.SheetID)
			filePath, err := filepath.Abs(fmt.Sprintf(tmpFilePath, r.counter))
			if err != nil {
				fmt.Printf("invalid filepath: %s\n", err.Error())
				continue
			}

			err = downloadCSV(url, filePath)
			if err != nil {
				fmt.Printf("failed to download CSV: %v", err)
				continue
			}

			response, err := r.client.UploadFile(filePath)
			if err != nil {
				continue
			}

			r.counter++
			r.storeCIDCh <- &StoreCID{
				ResearchID: data.ResearchID,
				Name:       response.Name,
				CID:        response.Hash,
			}
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
