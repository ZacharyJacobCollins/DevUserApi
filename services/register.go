package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//keypairs like ---  {"username": "zcollin1", "password":"test", "number": "734515513", "email":"collinswhatever@live"}

var userIdCounter uint32 = 0
var userStore = []User{}

type User struct {
	Id           uint32 `json:"id"`
	Username     string 	`json:"username"`
	Password     string 	`json:"password"`
	Number	     string 	`json:"number"`
	Email 	     string	`json:"email"`
	//Photo byte?
	//Calendar   NA
	//Contacts     []&User 	`json:"contacts"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	p := User{}

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

	u := User{
		Id:           	userIdCounter,
		Username:     	p.Username,
		Number: 	p.Number,
	}

	userStore = append(userStore, u)
	userIdCounter += 1
	w.WriteHeader(http.StatusCreated)
}

//Mke faster, not linear search here
func validateUniqueness(username string) error {
	for _, u := range userStore {
		if u.Username == username {
			return errors.New("Username is already used")
		}
	}
	return nil
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := json.Marshal(userStore)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(users)
}