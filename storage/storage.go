package storage

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("data not found")
)

type Storage[T any] interface {
	Save(ctx context.Context, t T) error
	SaveAll(ctx context.Context, t *[]*T)
	Get(ctx context.Context, id string) (*T, error)
	GetAll(ctx context.Context) (*[]*T, error)
	Delete(ctx context.Context, id string) error
}
