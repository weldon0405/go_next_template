package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
}

// mux.HandleFunc("/api/go/users", app.GetUsers).Methods("GET")
// mux.HandleFunc("/api/go/users/{id}", app.GetUser).Methods("GET")
// mux.HandleFunc("/api/go/users", app.CreateUser).Methods("POST")
// mux.HandleFunc("/api/go/users/{id}", app.UpdateUser).Methods("PUT")
// mux.HandleFunc("/api/go/users/{id}", app.DeleteUser).Methods("DELETE")
