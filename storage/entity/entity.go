package entity

type C2COrder struct {
	Id          string
	OrderNumber string
	TradeType   string
	Fiat        string
	Amount      float64
	TotalPrice  float64
	CreateTime  uint64
}
