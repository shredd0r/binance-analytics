package sqlite

import (
	"context"
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shredd0r/binance-analytics/storage"
	sqlite "github.com/shredd0r/binance-analytics/storage/sqlite/query"
)

type SqliteStorage[T any] struct {
	logger    *slog.Logger
	formatter sqlite.QueryFormatting[T]
	dbConn    *sql.Conn
}

func NewSqliteStorage[T any](logger *slog.Logger, dbConn *sql.Conn, formatter sqlite.QueryFormatting[T]) storage.Storage[T] {
	return &SqliteStorage[T]{
		logger:    logger,
		formatter: formatter,
		dbConn:    dbConn,
	}
}

func (s *SqliteStorage[T]) Save(ctx context.Context, t T) error {
	s.logger.Debug("start save entity in sqlite")
	query, args := s.formatter.FormatForSave(t)
	_, err := s.dbConn.ExecContext(ctx, query, args...)

	if err != nil {
		s.logger.Error("error when try save entity, err: %s", err.Error())
		return err
	}

	return nil
}

func (s *SqliteStorage[T]) SaveAll(ctx context.Context, t *[]*T) {

}
func (s *SqliteStorage[T]) Get(ctx context.Context, id string) (*T, error) {
	return nil, nil
}
func (s *SqliteStorage[T]) GetAll(ctx context.Context) (*[]*T, error) {
	return nil, nil
}
func (s *SqliteStorage[T]) Delete(ctx context.Context, id string) error {
	return nil
}
