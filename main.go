package main

import (
	"context"
	"fmt"

	"github.com/caarlos0/env/v11"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/shredd0r/binance-reader/config"
)

func main() {
	var cfg *config.Config
	err := env.Parse(cfg)

	// cfg, err := config.Read()

	if err != nil {
		print(err)
	}

	GetAllOrders(cfg.Credentials.ApiKey, cfg.Credentials.SecretKey)
}

func GetAllOrders(apiKey string, secretKey string) {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	resp, err := client.NewGetAllCoinsInfoService().Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(resp))
}
