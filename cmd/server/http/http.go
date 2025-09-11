package http

import (
	"log"
	"net/http"

	"github.com/george124816/gelection/internal/candidate"
)

func Start() {
	router := http.NewServeMux()

	router.HandleFunc("/candidate", candidate.CandidateHandler)
	router.HandleFunc("/candidate/{id}", candidate.CandidateHandler)

	log.Println("serving...")

	http.ListenAndServe(":4000", router)
}
