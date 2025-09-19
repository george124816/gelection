package http

import (
	"log"
	"net/http"

	"github.com/george124816/gelection/internal/candidate/handler"
	"github.com/george124816/gelection/internal/configs"
	engine "github.com/george124816/gelection/internal/db"
)

func Start() {
	config := configs.HttpConfig{Port: 4000}

	engine.Connect()

	router := http.NewServeMux()

	router.HandleFunc("/candidate", handler.CandidateHandler)
	router.HandleFunc("/candidates", handler.CandidateHandler)
	router.HandleFunc("/candidate/{id}", handler.CandidateHandler)

	log.Println("starting server on port", config.Port)

	http.ListenAndServe(config.GetStringPort(), router)

}
