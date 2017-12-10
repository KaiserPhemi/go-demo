// handler for a basic web app.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// application level context for handler
type App struct{}

func (myApp App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Welcome to golang")
}

func main() {
	var myApp App

	log.Print("Listening on localhost:5000. Ctrl-c to cancel.")
	log.Fatal(http.ListenAndServe("localhost:5000", myApp))
}
