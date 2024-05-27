package services

import "time"

const (
	defaultQueryInterval = 40 * time.Second
	expressServerURL     = "http://localhost:%s"
	retrieveFormDataURL  = "https://docs.google.com/spreadsheets/d/%s/export?format=csv&gid=%d"
	tmpFilePath          = "tmp/%s_%d.csv"
)
