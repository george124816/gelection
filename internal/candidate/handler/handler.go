package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/george124816/gelection/internal/candidate/model"
	"github.com/george124816/gelection/internal/candidate/repository"
	engine "github.com/george124816/gelection/internal/db"
)

func CandidateHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET" && r.Pattern == "/candidate/{id}":
		inputId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Fprintln(w, err)
		}

		candidate, err := repository.GetCandidate(context.Background(), engine.Engine, uint64(inputId))
		result, err := json.Marshal(candidate)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		fmt.Fprintln(w, string(result))
	case r.Method == "GET" && r.Pattern == "/candidates":
		candidates, err := repository.GetAll(context.Background(), engine.Engine)

		if err != nil {
			fmt.Fprintln(w, err)
		}

		resultJson, err := json.Marshal(candidates)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintln(w, string(resultJson))

	case r.Method == "UPDATE" && r.Pattern == "/candidate/{id}":
		var requestCandidate model.Candidate
		inputId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
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

	case r.Method == "POST" && r.Pattern == "/candidate":
		var requestCandidate model.Candidate
		bodyRequest, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println("body request failed")
		}

		err = json.Unmarshal(bodyRequest, &requestCandidate)

		if err != nil {
			log.Println("failed to decode json")
		}

		err = repository.Create(requestCandidate)

		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "created")
	case r.Method == "DELETE" && r.Pattern == "/candidate/{id}":
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
