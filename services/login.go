package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//keypairs like ---  {"username": "zcollin1", "password":"test", "number": "734515513", "email":"collinswhatever@live"}

var userIdCounter uint32 = 0
var userStore = []User{}

type User struct {
	Id           uint32 	`json:"id"`
	Username     string 	`json:"username"`
	Password     string 	`json:"password"`
	Number	     string 	`json:"number"`
	Email 	     string	`json:"email"`
	//Photo byte?
	//Calendar   NA
	//Contacts     []&User 	`json:"contacts"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	u := User{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if (checkCredentials(u.Username, u.Password) != nil) {
		w.Write(u)
	}

	userIdCounter += 1
}

//check for user and password in array
func checkCredentials(username string, password string) User {
	for _, u := range userStore {
		if u.Username == username && u.Password == password {
			return u
		}
	}
	return nil
}