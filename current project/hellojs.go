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
