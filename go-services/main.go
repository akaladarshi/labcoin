package main

import (
	"github.com/akaladarshi/labcoin/config"
	"github.com/akaladarshi/labcoin/utils"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	utils.IntializeEthereumClient(cfg)
}
