package model

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func CreateModel(t *testing.T) {
	t.Run("should create vote struct", func(t *testing.T) {
		uuid := uuid.NewString()
		insertedAt := time.Now()
		electionId := 1
		candidateId := 1

		vote := Vote{
			Id:          uuid,
			InsertedAt:  insertedAt,
			ElectionId:  electionId,
			CandidateId: candidateId,
		}

		if vote.Id != uuid {
			t.Error("error")
		}
		if vote.InsertedAt != insertedAt {
			t.Error("error")
		}
		if vote.ElectionId != electionId {
			t.Error("error")
		}
		if vote.CandidateId != candidateId {
			t.Error("error")
		}

	})

}
