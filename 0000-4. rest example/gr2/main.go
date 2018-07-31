package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	username string `json:"username"`
	password string `json:"password"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/echo", echo).Methods("GET")

	log.Println("starting a server port 8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}

func echo(w http.ResponseWriter, r *http.Request) {

	log.Println("/echo")

	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	user.password = "1234"

	// return
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Println("error marshall")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)

}
