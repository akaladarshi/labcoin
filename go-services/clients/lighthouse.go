package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akaladarshi/labcoin/config"
)

const nodeURL = "https://node.lighthouse.storage/api/v0"

type LighthouseClient struct {
	cfg     *config.Config
	BaseURL string
	Client  *http.Client
}

func NewLighthouseClient(cfg *config.Config) *LighthouseClient {
	return &LighthouseClient{
		cfg:     cfg,
		BaseURL: nodeURL,
		Client:  &http.Client{},
	}
}

// UploadFile uploads a file to Lighthouse storage
func (c *LighthouseClient) UploadFile(filePath string) (*LighthouseResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a new multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	writer.Close()

	// Create a new HTTP request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/add", c.BaseURL), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.cfg.LightHouseAPIKEY)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to upload file, status code: %d, response: %s", resp.StatusCode, respBody)
	}

	var response LighthouseResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the reponse: %w", err)
	}

	return &response, nil
}
