package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func main() {
	// define route /greet with an anonymous func
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting webserver
	http.ListenAndServe("localhost:8000", nil)

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

// replace func greet with func getAllCusomers(.w,r)
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Bob", "NYC", "110075"},
		{"Rob", "Van", "v6c c6c"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// TERMINAL_1:
// go build main.go -o m
// run server: ./m         => it will start web server
// TERMINAL_2:
// curl http://localhost:8000/greet  <enter>
// OUT2: Hello World!


// TERMINAL_!:
// go run main.go
//
// TERMINAL_2:
// @ntrajic ➜ /workspaces/Building-Microservices-API-in-Go (main) $ curl http://localhost:8000/customers
// OUT:
// [{"Name":"Bob","City":"NYC","Zipcode":"110075"},{"Name":"Rob","City":"Van","Zipcode":"v6c c6c"}]
