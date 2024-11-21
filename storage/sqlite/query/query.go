package query

import "github.com/shredd0r/binance-analytics/storage/entity"

type QueryFormatting[T any] interface {
	FormatForSave(t T) (string, []any)
	FormatForGet(id string) (string, []any)
	FormatForGetAll() (string, []any)
	FormatForDelete(id string) (string, []any)
}

type C2CQueryFormatting struct {
}

func (f *C2CQueryFormatting) FormatForSave(e entity.C2COrder) (string, []any) {
	return "INSERT INTO c2c_orders (id, order_number, trade_type, fiat, amount, total_price, create_time) VALUES (?, ?, ?, ?, ?, ?, ?);",
		[]any{e.Id, e.OrderNumber, e.TradeType, e.Fiat, e.Amount, e.TotalPrice, e.CreateTime}
}

func (f *C2CQueryFormatting) FormatForGet(id string) (string, []any) {
	return "", []any{""}
}

func (f *C2CQueryFormatting) FormatForGetAll() (string, []any) {
	return "", []any{""}
}

func (f *C2CQueryFormatting) FormatForDelete(id string) (string, []any) {
	return "", []any{""}
}
