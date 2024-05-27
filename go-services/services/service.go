package services

import (
	"context"
	"fmt"

	"github.com/akaladarshi/labcoin/config"
	"github.com/akaladarshi/labcoin/utils"
)

type Service struct {
	qServc     *QueryService
	retrvServc *RetrievalService
	storeServc *StoreService
}

func NewService(cfg *config.Config) (*Service, error) {
	_, researchBinding, err := utils.IntializeEthereumClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize Ethereum client: %v", err)
	}

	dataChan := make(chan *FormData, 1)
	queryService, err := NewQueryService(cfg, researchBinding, dataChan)
	if err != nil {
		return nil, fmt.Errorf("unable to create query service: %v", err)
	}

	storeCIDCh := make(chan *StoreCID, 1)
	retrievalService, err := NewRetrievalService(cfg, dataChan, storeCIDCh)
	if err != nil {
		return nil, fmt.Errorf("unable to create retrieval service: %v", err)
	}

	storeService, err := NewStoreService(cfg, storeCIDCh)
	if err != nil {
		return nil, fmt.Errorf("failed to create store service: %w", err)
	}

	return &Service{
		qServc:     queryService,
		retrvServc: retrievalService,
		storeServc: storeService,
	}, nil
}

func (s *Service) Start(ctx context.Context) {
	go s.qServc.Start(ctx)
	go s.retrvServc.Start(ctx)
	go s.storeServc.Start(ctx)
}
