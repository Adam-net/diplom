package server

import (
	"net/http"
	"os"
)

func Run() error {
	port := getPort("TODO_PORT")
	http.Handle("/", http.FileServer(http.Dir("web")))
	return http.ListenAndServe(":"+port, nil)
}

func getPort(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return "7540"
}
