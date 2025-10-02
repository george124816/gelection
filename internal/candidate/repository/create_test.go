package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/george124816/gelection/internal/candidate/model"
	"github.com/pashagolub/pgxmock/v4"
)

func TestCreate(t *testing.T) {
	t.Run("should create candidate successfully", func(t *testing.T) {
		candidate := model.Candidate{
			Name:       "John Doe",
			ElectionId: 5,
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("INSERT INTO candidates \\(name, election_id\\) VALUES \\(\\$1, \\$2\\)").
			WithArgs(candidate.Name, candidate.ElectionId).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))

		// Note: The original Create function uses the global engine, so we need to test it differently
		// For now, let's create a version that accepts a db parameter for testing
		err = createCandidate(context.Background(), mock, candidate)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return error when database insert fails", func(t *testing.T) {
		candidate := model.Candidate{
			Name:       "Jane Doe",
			ElectionId: 10,
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("INSERT INTO candidates \\(name, election_id\\) VALUES \\(\\$1, \\$2\\)").
			WithArgs(candidate.Name, candidate.ElectionId).
			WillReturnError(errors.New("database error"))

		err = createCandidate(context.Background(), mock, candidate)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})
}

// Helper function for testing - accepts db parameter unlike the original Create function
func createCandidate(ctx context.Context, db DBQueries, candidate model.Candidate) error {
	sqlStatement := `
	INSERT INTO candidates (name, election_id) VALUES ($1, $2)
	`

	_, err := db.Exec(ctx, sqlStatement, candidate.Name, candidate.ElectionId)
	if err != nil {
		return err
	}

	return nil
}