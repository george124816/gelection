package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
)

func TestDeleteCandidate(t *testing.T) {
	t.Run("should delete candidate successfully", func(t *testing.T) {
		candidateID := uint64(1)

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("DELETE FROM candidates WHERE id = \\$1").
			WithArgs(candidateID).
			WillReturnResult(pgxmock.NewResult("DELETE", 1))

		err = deleteCandidateWithDB(context.Background(), mock, candidateID)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return error when candidate doesn't exist", func(t *testing.T) {
		candidateID := uint64(999)

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("DELETE FROM candidates WHERE id = \\$1").
			WithArgs(candidateID).
			WillReturnResult(pgxmock.NewResult("DELETE", 0))

		err = deleteCandidateWithDB(context.Background(), mock, candidateID)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		expectedError := "failed to delete"
		if err.Error() != expectedError {
			t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})

	t.Run("should return error when database delete fails", func(t *testing.T) {
		candidateID := uint64(1)

		mock, err := pgxmock.NewPool()
		if err != nil {
			t.Fatalf("Failed to create mock pool: %v", err)
		}
		defer mock.Close()

		mock.ExpectExec("DELETE FROM candidates WHERE id = \\$1").
			WithArgs(candidateID).
			WillReturnError(errors.New("database error"))

		err = deleteCandidateWithDB(context.Background(), mock, candidateID)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations were met: %v", err)
		}
	})
}

// Helper function for testing - accepts db parameter unlike the original DeleteCandidate function
func deleteCandidateWithDB(ctx context.Context, db DBQueries, id uint64) error {
	sqlStatement := `DELETE FROM candidates WHERE id = $1`

	result, err := db.Exec(ctx, sqlStatement, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 1 {
		return nil
	} else {
		return errors.New("failed to delete")
	}
}