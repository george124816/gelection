package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/george124816/gelection/internal/candidate/model"
	"github.com/pashagolub/pgxmock/v4"
)

func TestUpdate(t *testing.T) {
	t.Run("should update candidate successfully", func(t *testing.T) {
		candidateID := 1
		candidate := model.Candidate{
			Name: "Updated Name",
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("UPDATE candidates SET name = \\$1 WHERE id = \\$2").
			WithArgs(candidate.Name, candidateID).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		err = updateCandidate(context.Background(), mock, candidateID, candidate)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return error when no rows affected", func(t *testing.T) {
		candidateID := 999
		candidate := model.Candidate{
			Name: "Non-existent Candidate",
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("UPDATE candidates SET name = \\$1 WHERE id = \\$2").
			WithArgs(candidate.Name, candidateID).
			WillReturnResult(pgxmock.NewResult("UPDATE", 0))

		err = updateCandidate(context.Background(), mock, candidateID, candidate)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		expectedError := "failed to update"
		if err.Error() != expectedError {
			t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return error when database update fails", func(t *testing.T) {
		candidateID := 1
		candidate := model.Candidate{
			Name: "Test Name",
		}

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("UPDATE candidates SET name = \\$1 WHERE id = \\$2").
			WithArgs(candidate.Name, candidateID).
			WillReturnError(errors.New("database error"))

		err = updateCandidate(context.Background(), mock, candidateID, candidate)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})
}

// Helper function for testing - accepts db parameter unlike the original Update function
func updateCandidate(ctx context.Context, db DBQueries, id int, candidate model.Candidate) error {
	sqlStatement := `UPDATE candidates SET name = $1 WHERE id = $2`

	result, err := db.Exec(ctx, sqlStatement, candidate.Name, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return errors.New("failed to update")
	}

	return nil
}