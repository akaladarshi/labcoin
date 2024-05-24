package utils

import (
	"fmt"
	"log"

	"github.com/akaladarshi/labcoin/bindings"
	"github.com/akaladarshi/labcoin/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IntializeEthereumClient(cfg *config.Config) (*ethclient.Client, *bindings.Research, error) {
	client, err := ethclient.Dial(cfg.NetCfg.URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress(cfg.ResearchContractAddress)
	instance, err := bindings.NewResearch(contractAddress, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to instantiate a Research contract: %v", err)
	}

	return client, instance, nil
}
