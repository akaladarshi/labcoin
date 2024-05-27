package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/akaladarshi/labcoin/config"
	"github.com/akaladarshi/labcoin/services"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	service, err := services.NewService(cfg)
	if err != nil {
		panic(err)
	}

	service.Start(ctx)
	waitForStopSignal()

}

func waitForStopSignal() {
	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)
	<-killSignal
}
