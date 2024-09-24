package main

import (
	"fmt"
	"net/http"
)


type Customer struct {
	Name string
	City string
	Zipcode string
}

func main() {
	// define route /greet with an anonymous func
	http.HandleFunc("/greet", greet)

	// starting webserver
	http.ListenAndServe("localhost:8000", nil)

}

// take anonymous func aut, and declare separate func greet
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

// TERMINAL_1:
// go build main.go -o m
// run server: ./m         => it will start web server
// TERMINAL_2:
// curl http://localhost:8000/greet  <enter>
// OUT2: Hello World!
