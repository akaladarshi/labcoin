package services

import (
	"context"
	"fmt"

	"github.com/akaladarshi/labcoin/clients"
	"github.com/akaladarshi/labcoin/config"
)

type StoreService struct {
	cfg            *config.Config
	contractClient *clients.ResearchContractClient
	storeCIDCh     <-chan *StoreCID
}

func NewStoreService(cfg *config.Config, storeCID <-chan *StoreCID) (*StoreService, error) {
	client, err := clients.NewResearchClient(fmt.Sprintf(expressServerURL, cfg.ExpressServerPort))
	if err != nil {
		return nil, fmt.Errorf("failed to create research client: %v", err)
	}

	return &StoreService{
		cfg:            cfg,
		storeCIDCh:     storeCID,
		contractClient: client,
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

			_, err = s.contractClient.StoreCID(data)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}
