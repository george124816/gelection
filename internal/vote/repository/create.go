package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/george124816/gelection/internal/vote/model"
	"github.com/google/uuid"
)

func Create(ctx context.Context, db DBQueries, vote model.Vote) error {
	sqlStatement := `
	INSERT INTO
            votes (id, inserted_at, election_id, candidate_id)
        VALUES ($1, $2, $3, $4)
	`

	_, err := db.Exec(ctx, sqlStatement, uuid.NewString(), time.Now(), vote.ElectionId, vote.CandidateId)

	if err != nil {
		slog.Error("failed to create vote", "error", err)

		return err
	}

	return nil

}
