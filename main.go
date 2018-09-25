package main

import (
	"encoding/json"
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
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Chibuike", Lastname: "Kenneth", Address: &Address{City: "Yaba", State: "Lagos"}})
	people = append(people, Person{ID: "1", Firstname: "Anthony", Lastname: "Kenneth"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
