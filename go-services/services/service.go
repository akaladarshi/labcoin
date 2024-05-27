package services

import (
	"context"
	"fmt"

	"github.com/akaladarshi/labcoin/bindings"
	"github.com/akaladarshi/labcoin/config"
	"github.com/akaladarshi/labcoin/utils"
)

type Service struct {
	cfg             *config.Config
	researchBinding *bindings.Research
	formDataCh      chan *FormData
	qServc          *QueryService
	retrvServc      *RetrievalService
}

func NewService(cfg *config.Config) (*Service, error) {
	_, researchBinding, err := utils.IntializeEthereumClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize Ethereum client: %v", err)
	}

	dataChan := make(chan *FormData, 100)
	queryService, err := NewQueryService(cfg, researchBinding, dataChan)
	if err != nil {
		return nil, fmt.Errorf("unable to create query service: %v", err)
	}

	retrievalService, err := NewRetrievalService(cfg, dataChan)
	if err != nil {
		return nil, fmt.Errorf("unable to create retrieval service: %v", err)
	}

	return &Service{
		qServc:     queryService,
		retrvServc: retrievalService,
	}, nil
}

func (s *Service) Start(ctx context.Context) {
	go s.qServc.Start(ctx)
	go s.retrvServc.Start(ctx)
}
