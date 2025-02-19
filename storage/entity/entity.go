package entity

type TradeType uint8

const (
	TradeTypeBuy = iota
	TradeTypeSell
)

type C2COrder struct {
	Id          string
	OrderNumber string
	TradeType   TradeType
	Asset       string
	Fiat        string
	Amount      float64
	TotalPrice  float64
	CreateTime  uint64
}
