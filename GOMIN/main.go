package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	data := Employee{
		FirstName: "Amit",
		LastName:  "Kumar",
		Email:     "baddhanamit@gmail.com",
		Age:       22,
	}
	json.NewEncoder(w).Encode(data)

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", GetBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}
