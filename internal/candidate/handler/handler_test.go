package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/george124816/gelection/internal/candidate/model"
	"github.com/george124816/gelection/internal/candidate/repository"
)

type MockAdapter struct {
	ReturnedCandidate model.Candidate
}

func (d MockAdapter) GetCandidate(ctx context.Context, db repository.DBQueries, id uint64) (model.Candidate, error) {
	return d.ReturnedCandidate, nil
}

func (d MockAdapter) GetAllCandidates(ctx context.Context, db repository.DBQueries) ([]model.Candidate, error) {
	return []model.Candidate{}, nil
}

func TestCandidateRetrieveUpdateDestroyHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/candidates/{id}", CandidateRetrieveUpdateDestroyHandler)

	t.Run("should return successfully", func(t *testing.T) {
		expectedCandidate := model.Candidate{
			Id:         1,
			Name:       "Zé do Caixão",
			ElectionId: 2,
		}

		adapter = MockAdapter{ReturnedCandidate: expectedCandidate}

		request, err := http.NewRequest("GET", "/candidates/24", nil)

		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()
		mux.ServeHTTP(response, request)

		returnedBody := response.Body
		expectedBody := "{\"id\":1,\"name\":\"Zé do Caixão\",\"election_id\":2}\n"

		if returnedBody.String() != expectedBody {
			t.Fatalf("expected %s returned %s", expectedBody, returnedBody)

		}

		expectedContentType := "application/json"
		returnedContentType := response.Header().Get("Content-Type")

		if returnedContentType != expectedContentType {
			t.Fatalf("expected %s returned %s", expectedContentType, returnedContentType)

		}

		expectedStatusCode := 200
		returnedStatusCode := response.Result().StatusCode

		if returnedStatusCode != expectedStatusCode {
			t.Fatalf("expected %d returned %d", expectedStatusCode, returnedStatusCode)

		}
	})

}
