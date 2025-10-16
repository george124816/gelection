package repository

import (
	"context"

	"github.com/george124816/gelection/internal/election/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBQueries interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type Adapter interface {
	GetAllElections(ctx context.Context, db DBQueries) ([]model.Election, error)
	Create(ctx context.Context, db DBQueries, election model.Election) error
	GetElection(ctx context.Context, db DBQueries, id int) (model.Election, error)
}

type DefaultAdapter struct{}
