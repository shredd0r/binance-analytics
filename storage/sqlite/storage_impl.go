package sqlite

import (
	"database/sql"
	"log/slog"

	"github.com/shredd0r/binance-analytics/storage"
	"github.com/shredd0r/binance-analytics/storage/entity"
	"github.com/shredd0r/binance-analytics/storage/sqlite/query"
)

type C2COrderStorage struct {
	storage.Storage[entity.C2COrder]
}

func NewC2COrderStorage(logger *slog.Logger, dbConn *sql.Conn) storage.Storage[entity.C2COrder] {
	s := NewSqliteStorage(
		logger,
		dbConn,
		&query.C2CQueryFormatting{},
	)

	return &C2COrderStorage{
		Storage: s,
	}
}
