package candidate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func CandidateHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Pattern == "/candidate/{id}" && r.Method == "GET":
		inputId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Fprintln(w, err)
		}
		candidate, err := GetCandidate(uint64(inputId))
		fmt.Fprintln(w, candidate)
	case r.Method == "GET":
		fmt.Fprintln(w, "GET")
	case r.Method == "POST" && r.Pattern == "/candidate":
		var requestCandidate Candidate
		bodyRequest, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println("body request failed")
		}

		err = json.Unmarshal(bodyRequest, &requestCandidate)

		if err != nil {
			log.Println("failed to decode json")
		}

		err = CreateCandidate(requestCandidate)

		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "created")
	default:
		fmt.Println("not found")
	}
}
