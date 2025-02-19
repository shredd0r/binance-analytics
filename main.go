package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shredd0r/binance-analytics/config"
	binance_connector "github.com/shredd0r/binance-connector-go"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	// db, _ := sql.Open("sqlite3", sqlite.Path_to_file)
	// db.Exec(sqlite.SchemaSQL)

	// dbConn, _ := db.Conn(context.Background())

	// str := sqlite.NewC2COrderStorage(logger, dbConn)
	// str.Save(context.Background(), entity.C2COrder{
	// 	Id:          "123",
	// 	OrderNumber: "123",
	// 	TradeType:   "123",
	// 	Fiat:        "123",
	// 	Amount:      13.34,
	// 	TotalPrice:  123.01,
	// 	CreateTime:  1231412,
	// })

	cfg, _ := config.Read()

	GetAllOrders(logger, cfg.Credentials.ApiKey, cfg.Credentials.SecretKey)
}

func GetAllOrders(logger *slog.Logger, apiKey string, secretKey string) {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// r, err := client.NewGetAccountService().Do(context.Background())

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(binance_connector.PrettyPrint(r))

	resp, err := client.NewGetC2CTradeHistoryService().StartTimestamp(1731940691000).Timestamp(uint64(time.Now().UnixMilli())).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(resp))
}
