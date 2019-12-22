package main

import (
	"fmt"
	"log"
	"net/http"
)


//initial run apeared "Hi there, I love %s!monkeys" http://localhost:8080/monkeys - typo of Fprint
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}