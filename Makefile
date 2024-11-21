build-linux:
	go build -ldflags "-s -w" -o ./bin/binance-reader -a ./main.go

load-env:
	export $(grep -v '^#' .env | xargs)

