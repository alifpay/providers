package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alifpay/providers/app"
	"github.com/alifpay/providers/models"
	"github.com/shopspring/decimal"
)

func main() {
	decimal.MarshalJSONWithoutQuotes = true
	var cfg models.Config

	cfg.Host = os.Getenv("HOST")
	if cfg.Host == "" {
		cfg.Host = ":8088"
	}

	//version of api
	cfg.Version = os.Getenv("VERSION")
	if cfg.Version == "" {
		log.Fatalln("Environment variable Version is empty")
	}

	s := app.New(cfg)

	c, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sig
		log.Println("Shutting down")
		s.Shutdown(c)
		signal.Stop(sig)
		close(sig)
		cancel()
	}()

	log.Println("api is starting")
	err := s.Run()
	if err != nil {
		log.Fatalln("api server:", err)
	}
}
