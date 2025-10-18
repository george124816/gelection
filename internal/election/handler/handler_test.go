package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/george124816/gelection/internal/election/model"
	"github.com/george124816/gelection/internal/election/repository"
)

type MockAdapter struct {
	GetAllElectionsReturn struct {
		Elections []model.Election
		err       error
	}

	GetElectionReturn struct {
		election *model.Election
		err      error
	}

	CreateReturn error
}

func (m MockAdapter) GetAllElections(ctx context.Context, db repository.DBQueries) ([]model.Election, error) {
	return m.GetAllElectionsReturn.Elections, m.GetAllElectionsReturn.err
}

func (m MockAdapter) GetElection(ctx context.Context, db repository.DBQueries, id int) (*model.Election, error) {
	return m.GetElectionReturn.election, m.GetElectionReturn.err
}

func (m MockAdapter) Create(ctx context.Context, db repository.DBQueries, election model.Election) error {
	return m.CreateReturn
}

func TestElectionListCreateHandler(t *testing.T) {
	t.Run("GET", func(t *testing.T) {
		t.Run("should return error", func(t *testing.T) {
			adapter = MockAdapter{GetAllElectionsReturn: struct {
				Elections []model.Election
				err       error
			}{
				err: errors.New("repository is not available"),
			}}

			request, err := http.NewRequest("GET", "/elections", nil)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()

			ElectionListCreateHandler(response, request)

			expected := "repository is not available\n"
			result := response.Body.String()

			if expected != result {
				t.Fatalf("expected %v, returned %v", expected, result)
			}

		})
		t.Run("election list handler returns empty", func(t *testing.T) {
			adapter = MockAdapter{GetAllElectionsReturn: struct {
				Elections []model.Election
				err       error
			}{}}

			request, err := http.NewRequest("GET", "/elections", nil)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()

			ElectionListCreateHandler(response, request)

			expectedElection := []model.Election{}

			var returnedBody []model.Election

			body, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = json.Unmarshal(body, &returnedBody)

			if !slices.Equal(expectedElection, returnedBody) {
				t.Fatalf("expected %v, returned %v", expectedElection, returnedBody)
			}
		})
		t.Run("election list handler returns items", func(t *testing.T) {
			adapter = MockAdapter{
				GetAllElectionsReturn: struct {
					Elections []model.Election
					err       error
				}{
					Elections: []model.Election{
						model.Election{
							Id:   1,
							Name: "A",
						},
						model.Election{
							Id:   2,
							Name: "B",
						},
						model.Election{
							Id:   3,
							Name: "C",
						},
					},
				}}

			request, err := http.NewRequest("GET", "/elections", nil)
			if err != nil {
				t.Fatal(err)
			}

			response := httptest.NewRecorder()

			ElectionListCreateHandler(response, request)

			expectedElection := []model.Election{
				model.Election{
					Id:   1,
					Name: "A",
				},
				model.Election{
					Id:   2,
					Name: "B",
				},
				model.Election{
					Id:   3,
					Name: "C",
				},
			}

			var returnedBody []model.Election

			body, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}

			err = json.Unmarshal(body, &returnedBody)

			if !slices.Equal(expectedElection, returnedBody) {
				t.Fatalf("expected %v, returned %v", expectedElection, returnedBody)
			}
		})
	})

	t.Run("POST", func(t *testing.T) {
		t.Run("should create successfully", func(t *testing.T) {
			adapter = MockAdapter{
				CreateReturn: nil,
			}

			payload := bytes.NewBuffer([]byte(`{"name": "Crazy"}`))
			request, err := http.NewRequest("POST", "/elections", payload)

			if err != nil {
				t.Error(err)
			}

			response := httptest.NewRecorder()
			ElectionListCreateHandler(response, request)

			returnedBody := response.Body
			expectedBody := "created\n"

			if returnedBody.String() != expectedBody {
				t.Fatalf("expected %s returned %s", expectedBody, returnedBody)
			}

			expectedStatusCode := 201
			returnedStatusCode := response.Result().StatusCode

			if returnedStatusCode != expectedStatusCode {
				t.Fatalf("expected %d returned %d", expectedStatusCode, returnedStatusCode)

			}

		})

		t.Run("should fail when database returns error", func(t *testing.T) {
			adapter = MockAdapter{
				CreateReturn: errors.New("failed to create in database"),
			}

			payload := bytes.NewBuffer([]byte(`{"name": "Crazy"}`))
			request, err := http.NewRequest("POST", "/elections", payload)

			if err != nil {
				t.Error(err)
			}

			response := httptest.NewRecorder()
			ElectionListCreateHandler(response, request)

			returnedBody := response.Body
			expectedBody := "failed to create in database\n"

			if returnedBody.String() != expectedBody {
				t.Fatalf("expected %s returned %s", expectedBody, returnedBody)
			}

			expectedStatusCode := 409
			returnedStatusCode := response.Result().StatusCode

			if returnedStatusCode != expectedStatusCode {
				t.Fatalf("expected %d returned %d", expectedStatusCode, returnedStatusCode)

			}

		})

	})
}

func TestElectionRetrieveHandler(t *testing.T) {

	t.Run("GET", func(t *testing.T) {
		t.Run("should return election successfully", func(t *testing.T) {
			adapter = MockAdapter{
				GetElectionReturn: struct {
					election *model.Election
					err      error
				}{
					election: &model.Election{
						Id:   8788,
						Name: "Some Election",
					},
				},
			}

			request, err := http.NewRequest("GET", "/elections/8788", nil)
			if err != nil {
				t.Error(err)
			}

			response := httptest.NewRecorder()

			mux := http.NewServeMux()
			mux.HandleFunc("/elections/{id}", ElectionRetrieveHandler)
			mux.ServeHTTP(response, request)

			returnedBody := response.Body.String()
			expectedBody := "{\"id\":8788,\"name\":\"Some Election\"}\n"

			if returnedBody != expectedBody {
				t.Fatalf("expected %s returned %s", expectedBody, returnedBody)
			}
		})

		t.Run("should return not found", func(t *testing.T) {
			adapter = MockAdapter{
				GetElectionReturn: struct {
					election *model.Election
					err      error
				}{
					election: nil,
					err:      errors.New("election not found"),
				},
			}

			request, err := http.NewRequest("GET", "/elections/8788", nil)
			if err != nil {
				t.Error(err)
			}

			response := httptest.NewRecorder()

			mux := http.NewServeMux()
			mux.HandleFunc("/elections/{id}", ElectionRetrieveHandler)
			mux.ServeHTTP(response, request)

			returnedBody := response.Body.String()
			expectedBody := "election_not_found\n"

			if returnedBody != expectedBody {
				t.Fatalf("expected %s returned %s", expectedBody, returnedBody)
			}

			expectedStatusCode := 404
			returnedStatusCode := response.Result().StatusCode

			if returnedStatusCode != expectedStatusCode {
				t.Fatalf("expected %d returned %d", expectedStatusCode, returnedStatusCode)

			}
		})

	})
}
