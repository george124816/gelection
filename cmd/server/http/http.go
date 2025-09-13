package http

import (
	"log"
	"net/http"

	"github.com/george124816/gelection/internal/candidate/handler"
)

func Start() {
	router := http.NewServeMux()

	router.HandleFunc("/candidate", handler.CandidateHandler)
	router.HandleFunc("/candidate/{id}", handler.CandidateHandler)

	log.Println("serving...")

	http.ListenAndServe(":4000", router)
}
