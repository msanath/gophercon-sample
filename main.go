package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// find port selected by service, default to 8080
	var port string
	var found bool
	port, found = os.LookupEnv("PORT")
	if found == true {
		// prepend a colon
		port = fmt.Sprintf(":%s", port)
	} else {
		port = ":8080"
	}

	http.HandleFunc("/", tester)

	log.Fatal(http.ListenAndServe(port, nil))
}

// tester
func tester(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World %s!", html.EscapeString(
		r.URL.Query().Get("name")))
	log.Printf("request received, details follow:\n%+v\n", r)
}
