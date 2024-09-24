package main

import (
	"log"
	"net/http"
)

func Start(){
		// define route /greet with an anonymous func
		http.HandleFunc("/greet", greet)
		http.HandleFunc("/customers", getAllCustomers)
	
		// starting webserver
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	
}