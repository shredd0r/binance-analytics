module github.com/shredd0r/binance-analytics

go 1.22.6

require (
	github.com/shredd0r/binance-connector-go v0.0.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/bitly/go-simplejson v0.5.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/mattn/go-sqlite3 v1.14.24
)

replace github.com/binance/binance-connector-go => github.com/shredd0r/binance-connector-go v0.0.2
