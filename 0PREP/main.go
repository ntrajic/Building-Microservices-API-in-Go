package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string	`json:"full_name" xml:"name"` 
	City    string	`json:"city" xml:"city"`
	Zipcode string	`json:"zip_code" xml:"zipcode"`
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
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

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

// Changing Tags/field_names in the response - directly in the structure: new tags are "full_name", "city", "zip_code"
// TERMINAL_1:
// go run main.go     // start the server
//
// TERMINAL_2:
//@ntrajic ➜ /workspaces/Building-Microservices-API-in-Go (main) $ curl http://localhost:8000/customers
// [{"full_name":"Bob","city":"NYC","zip_code":"110075"},{"full_name":"Rob","city":"Van","zip_code":"v6c c6c"}]


//XML MARSHALING
//TERMINAL_1: 
//go run main.go //start web server
//
//TERMINAL_2:
//@ntrajic ➜ /workspaces/Building-Microservices-API-in-Go (main) $ curl http://localhost:8000/customers
//OUT:
//<Customer><name>Bob</name><city>NYC</city><zipcode>110075</zipcode></Customer><Customer><name>Rob</name><city>Van</city><zipcode>v6c c6c</zipcode></Customer>
