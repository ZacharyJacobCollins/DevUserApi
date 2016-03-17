package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

//keypairs like ---  {"username": "zcollin1", "password":"test", "number": "734515513", "email":"collinswhatever@live"}

var userIdCounter uint32 = 0
var userStore = []StoredUser{}

type InputUser struct {
	Username     string 	`json:"username"`
	Password     string 	`json:"password"`
	Number	     string 	`json:"number"`
	Email 	     string	`json:"email"`
	//Photo byte?
	//Calendar   NA
	//Contacts     []&User 	`json:"contacts"`
}

type StoredUser struct {
	Id           uint32 `json:"id"`
	Username     string 	`json:"username"`
	Password     string 	`json:"password"`
	Number	     string 	`json:"number"`
	Email 	     string	`json:"email"`
}

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

	err = validateUniqueness(p.Username)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := StoredUser{
		Id:           	userIdCounter,
		Username:     	p.Username,
		Number: 	p.Number,
	}

	userStore = append(userStore, u)

	userIdCounter += 1

	w.WriteHeader(http.StatusCreated)
}

func validateUniqueness(username string) error {
	for _, u := range userStore {
		if u.Username == username {
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