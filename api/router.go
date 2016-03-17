package api

import (
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", CreateUserHandler).Methods("POST")
	//only for testing purposes
	r.HandleFunc("/register", ListUsersHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("GET")
	return r
}