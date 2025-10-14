package repository

import (
	"context"
	"log/slog"
	"os"

	"github.com/george124816/gelection/internal/candidate/model"

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

type Adapter interface {
	GetCandidate(context.Context, DBQueries, uint64) (model.Candidate, error)
	GetAllCandidates(ctx context.Context, db DBQueries) ([]model.Candidate, error)
}

type DefaultAdapter struct {
}

func (d DefaultAdapter) GetCandidate(ctx context.Context, db DBQueries, id uint64) (model.Candidate, error) {
	var candidate model.Candidate
	sqlStatement := `SELECT * FROM candidates WHERE id = $1`

	err := db.QueryRow(ctx, sqlStatement, id).Scan(&candidate.Id, &candidate.Name, &candidate.ElectionId)

	if err != nil {
		slog.Error("failed to get candidate", "error", err)
		return candidate, err
	}
	GetCandidateCount.Add(context.Background(), 1)

	return candidate, nil

}

func (d DefaultAdapter) GetAllCandidates(ctx context.Context, db DBQueries) ([]model.Candidate, error) {
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
			slog.Error("failed to scan next candidate", "error", err)
			os.Exit(1)
		}
		candidates = append(candidates, c)
	}

	GetAllCandidateCount.Add(context.Background(), 1)

	return candidates, nil
}
