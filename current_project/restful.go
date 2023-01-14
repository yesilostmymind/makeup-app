package main

//packages, other Go source files you want to use are added with import paths, if you need to use say json for example, you must first import it here, go is like: json, I don't know her

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//this struct groups together the data for the type event, including the ID, title and description
type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title,omitempty"`
	Description string `json:"Description,omitempty"`
}

//why a type and not a var? either way its a slice
type allEvents []event

//dummy database, please include commas at the end of lines 
var events = allEvents{
	{
		ID:          "1",
		Title:       "Glaze #1 by Inn Beauty",
		Description: "plumping gloss oil hybrid, translucent red, candy apple flavor",
	},
}

//function that are not used are meaningless, declare them in the main statment if you intend to use it. 
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "title and description only for update, thenk you!")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

//this shows a single event, see how the method is specified in the main func
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

//event updating works, I haven't gotten it to show the err yet tho
func updateEvent (w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "event update requires title and description only, please do")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

//deleteEvent working, please remember to include what verable you want used, w/o eventID printed "...ID with %v! (DELETE)..."
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted, oops and/or yay!", eventID)
		}
	}
}

//this is a function called homeLink will display "welcome home!" 
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//this is the main function, it servers homeLink on port 8080
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
