package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/akaladarshi/labcoin/bindings"
)

type ResearchContractClient struct {
	BaseURL *url.URL
	Client  *http.Client
}

type FromDetailsResponse struct {
	AllFormDetails []*bindings.ResearchFormData `json:"AllFormDetails"`
}

// NewResearchClient creates a new ResearchContractClient with the given base URL
func NewResearchClient(baseURL string) (*ResearchContractClient, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	return &ResearchContractClient{
		BaseURL: parsedURL,
		Client:  &http.Client{},
	}, nil
}

// GetAllFormDetails fetches all the form details from the ResearchContractClient
func (c *ResearchContractClient) GetAllFormDetails() ([]*ResearchFormData, error) {
	path := c.BaseURL.JoinPath("/allformdetails")

	resp, err := c.Client.Get(path.String())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var researches []*ResearchFormData
	if err = json.Unmarshal(body, &researches); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return researches, nil
}
