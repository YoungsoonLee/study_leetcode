package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	UserName  string    `json:"username"`
	Password  string    `json:"password"` // save encripted password
	IsAdmin   bool      `json:"isadmin"`
	CreatedAt time.Time `json:"createdat"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/echo", echo).Methods("GET")
	r.HandleFunc("/echo/{id}", echo2).Methods("GET")

	log.Println("server start with port 8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}

func echo2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println(id)

	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	user := User{}

	// parse json request body and use it to set fields on user
	// note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	user.CreatedAt = time.Now().Local()

	// marshall or convert user object back to json and write to response
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
