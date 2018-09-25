package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string `json:"id, omitempty"`
	Firstname string `json:"firstname, omitempty"`
	Lastname  string `json:"lastname, omitempty"`
}

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {

}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {

}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
