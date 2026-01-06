package main

import (
	"net/http"
)

type api struct{
	addr string
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Users list ..."))
}

func (a *api) createuserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Created User"))
}

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()
	
	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("Post /user", api.createuserHandler)

	srv.ListenAndServe()
}