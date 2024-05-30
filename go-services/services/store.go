package services

import (
	"context"
	"fmt"
	"sync"

	"github.com/akaladarshi/labcoin/clients"
	"github.com/akaladarshi/labcoin/config"
)

type StoreService struct {
	cfg            *config.Config
	contractClient *clients.ResearchContractClient
	storeCIDCh     <-chan *StoreCID
	cache          *sync.Map
}

func NewStoreService(cfg *config.Config, storeCID <-chan *StoreCID, cache *sync.Map) (*StoreService, error) {
	client, err := clients.NewResearchClient(fmt.Sprintf(expressServerURL, cfg.ExpressServerPort))
	if err != nil {
		return nil, fmt.Errorf("failed to create research client: %v", err)
	}

	return &StoreService{
		cfg:            cfg,
		storeCIDCh:     storeCID,
		contractClient: client,
		cache:          cache,
	}, nil
}

func (s *StoreService) Start(ctx context.Context) {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case storeCID := <-s.storeCIDCh:
			data, err = storeCID.Marshal()
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}

			txHash, err := s.contractClient.StoreCID(data)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}

			fmt.Printf("Transaction hash: %s\n", txHash)
			if txHash != "" {
				s.cache.Delete(storeCID.CID)
			}

		case <-ctx.Done():
			return
		}
	}
}
