package main

import (
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	// Only the path goes here, method is handled inside the handler
	mux.HandleFunc("/users", api.getUserHandler)
	mux.HandleFunc("/user", api.createuserHandler)

	srv.ListenAndServe()
}
