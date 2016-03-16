package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

//keypairs like ---  {"username": "dennis", "balance": 200}

type InputUser struct {
	Name	     string `json:"name"`
	Number	     string `json:"number"`
}

type StoredUser struct {
	Id           uint32 `json:"id"`
	Name 	     string `json:"name"`
	Number	     string `json:"number"`
}

var userIdCounter uint32 = 0
var userStore = []StoredUser{}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	p := StoredUser{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = validateUniqueness(p.Name)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := StoredUser{
		Id:           	userIdCounter,
		Name:     	p.Name,
		Number: 	p.Number,
	}

	userStore = append(userStore, u)

	userIdCounter += 1

	w.WriteHeader(http.StatusCreated)
}

func validateUniqueness(username string) error {
	for _, u := range userStore {
		if u.Name == username {
			return errors.New("Username is already used")
		}
	}

	return nil
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := json.Marshal(userStore)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(users)
}

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users", listUsersHandler).Methods("GET")
	return r
}