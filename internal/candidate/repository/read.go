package repository

import (
	"context"
	"log/slog"
	"os"

	"github.com/george124816/gelection/internal/candidate/model"
	"github.com/jackc/pgx/v5"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const name = "github.com/george124816/gelection"

var (
	Meter                = otel.Meter(name)
	GetAllCandidateCount metric.Int64Counter
	GetCandidateCount    metric.Int64Counter
)

func init() {
	var err error
	GetAllCandidateCount, err = Meter.Int64Counter("gelection.get_all_candidates",
		metric.WithDescription("The Number of GetAllCandidates was called"),
	)

	GetCandidateCount, err = Meter.Int64Counter("gelection.get_candidate",
		metric.WithDescription("The Number of GetCanddiate was called"),
	)

	if err != nil {
		panic(err)
	}
}

type DBQueries interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

func GetCandidate(ctx context.Context, db DBQueries, id uint64) (model.Candidate, error) {
	var candidate model.Candidate
	sqlStatement := `SELECT * FROM candidates WHERE id = $1`

	err := db.QueryRow(ctx, sqlStatement, id).Scan(&candidate.Id, &candidate.Name, &candidate.ElectionId)

	if err != nil {
		slog.Error("failed to get candidate", err)
		return candidate, err
	}
	GetCandidateCount.Add(context.Background(), 1)

	return candidate, nil

}

func GetAllCandidates(ctx context.Context, db DBQueries) ([]model.Candidate, error) {
	slog.Debug("GetAllCandidates")

	var candidates []model.Candidate

	sqlStatement := `SELECT * FROM candidates`

	result, err := db.Query(ctx, sqlStatement)

	if err != nil {
		return nil, err
	}

	for result.Next() {
		var c model.Candidate
		if err := result.Scan(&c.Id, &c.Name, &c.ElectionId); err != nil {
			slog.Error("failed to scan next candidate", err)
			os.Exit(1)
		}
		candidates = append(candidates, c)
	}

	GetAllCandidateCount.Add(context.Background(), 1)

	return candidates, nil
}
