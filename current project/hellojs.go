/*
not currently compiling, error "inconsistent vendoring in C:\Go\src:
        go.mod requires github.com/gorilla/mux v1.7.3 but vendor/modules.txt does not include it.
        run 'go mod tidy; go mod vendor' to sync" further errors upon the suggested run comand
*/

package main

//packages, other Go source files you want to use are added with import paths

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//this struct groups together the data for the type event, including the ID, title and description
type event struct {
	ID          string `json:"ID,omitempty"`
	Title       string `json:"Title,omitempty"`
	Description string `json:"Description,omitempty"`
}

//why a type and not a var? either way its a slice
type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Makeup collection application",
		Description: "This is an app to store your makeup and stuff or whatever"
	},
}

func createEvent(w http.ResponseWriter, r *http.Request){
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "title and description only for update, thenk you!")
	}

	json.Unmarchal(reqBody, &newEvent)
	event = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

//this is a function called homeLink will display "welcome home!" 
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//this is the main function, it servers homeLink on port 8080
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
