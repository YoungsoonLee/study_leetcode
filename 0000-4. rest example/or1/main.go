package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	UserName  string    `json:"username"`
	Password  string    `json:"password"` // save encripted password
	IsAdmin   bool      `json:"isadmin"`
	CreatedAt time.Time `json:"createdat"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/echo", echoHandler)

	log.Println("start server port 5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		log.Fatalln("sever start error, ", err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
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
