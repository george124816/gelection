package repository

import (
	"context"
	"log/slog"

	"github.com/george124816/gelection/internal/vote/model"
)

func GetAllVotes(ctx context.Context, db DBQueries) ([]model.Vote, error) {
	var votes []model.Vote

	sqlStatement := `SELECT * FROM votes`

	result, err := db.Query(ctx, sqlStatement)

	if err != nil {
		return nil, err
	}

	for result.Next() {
		var vote model.Vote
		if err := result.Scan(&vote.Id, &vote.InsertedAt, &vote.ElectionId, &vote.CandidateId); err != nil {
			slog.Error("failed to scan result from database", "error", err)
		}
		votes = append(votes, vote)
	}

	return votes, nil
}
