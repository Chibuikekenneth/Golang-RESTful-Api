package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
