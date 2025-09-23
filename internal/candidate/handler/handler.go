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

func CandidateRetrieveUpdateDestroyHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
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
	case r.Method == "UPDATE":
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
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	switch {
	case r.Method == "POST":
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
	case r.Method == "GET":
		candidates, err := repository.GetAll()

		if err != nil {
			fmt.Fprintln(w, err)
		}

		resultJson, err := json.Marshal(candidates)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintln(w, string(resultJson))
	default:
		fmt.Println("not found")
	}
}
