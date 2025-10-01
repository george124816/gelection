package handler

import (
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Not Allowed")
	}
}
