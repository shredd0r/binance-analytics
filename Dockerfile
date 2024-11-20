FROM golang:1.22.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /binance-reader

EXPOSE 8080

# Run
CMD ["/binance-reader"]