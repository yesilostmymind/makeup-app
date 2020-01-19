package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
/*jfc the reason it wasn't building was that the portion after string needed to be with a "backquote" `` its the one where the tilde key is  */
//struct are used for grouping data together, a collection of fields like first name, last name, address.. for a person
type Person struct {
	ID 		  string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

//here is the address struct
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

//the universal verable people is declated, it is a slice of
var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(&people)
}


func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ =json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Omega", Lastname: "Rae", Address: &Address{City: "Daly City", State: "CA"}})
	people = append(people, Person{ID: "2", Firstname: "Kitty", Lastname: "Cat"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":6767", router))
}