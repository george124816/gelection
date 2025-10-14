package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/george124816/gelection/internal/candidate/model"
	"github.com/george124816/gelection/internal/candidate/repository"
	engine "github.com/george124816/gelection/internal/db"
)

var adapter repository.Adapter = repository.DefaultAdapter{}

func CandidateRetrieveUpdateDestroyHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		inputId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			error_to_return := "failed to cast input"
			slog.Error(error_to_return)
			fmt.Fprintln(w, error_to_return)
			return
		}

		candidate, err := adapter.GetCandidate(context.Background(), engine.Engine, uint64(inputId))

		if err != nil {
			error_to_return := "failed to query database result"
			slog.Error(error_to_return)
			fmt.Fprintln(w, error_to_return)
			return
		}

		result, err := json.Marshal(candidate)
		if err != nil {
			slog.Error("failed to marshal to return param", "error", err)
			fmt.Fprintln(w, err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(result))

	case r.Method == "UPDATE":
		var requestCandidate model.Candidate
		inputId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			slog.Error("failed to convert the input", "error", err)
			fmt.Fprintln(w, err)
			return
		}

		bodyRequest, err := io.ReadAll(r.Body)

		err = json.Unmarshal(bodyRequest, &requestCandidate)

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		repository.Update(inputId, requestCandidate)

		w.WriteHeader(http.StatusGone)
		fmt.Fprintln(w, "UPDATED")
	case r.Method == "DELETE":
		inputId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Fprintln(w, err)
		}
		err = repository.DeleteCandidate(uint64(inputId))

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintln(w, "deleted")
	default:
		fmt.Println("not found")
	}
}

func CandidateListCreateHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "POST":
		var requestCandidate model.Candidate
		bodyRequest, err := io.ReadAll(r.Body)

		if err != nil {
			slog.Error("body reqeust failed", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}

		err = json.Unmarshal(bodyRequest, &requestCandidate)

		if err != nil {
			slog.Error("failed to decode json", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}

		err = repository.Create(requestCandidate)

		if err != nil {
			slog.Error("failed to create candidate", "error", err)
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "created")
	case r.Method == "GET":

		candidates, err := adapter.GetAllCandidates(context.Background(), engine.Engine)

		if err != nil {
			fmt.Fprintln(w, err)
		}

		resultJson, err := json.Marshal(candidates)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(resultJson))
	default:
		fmt.Println("not found")
	}
}
