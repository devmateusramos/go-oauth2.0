package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

	http.ListenAndServe(":8080", nil)
}

type person struct {
	First string
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Mateus",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	if err := json.NewDecoder(r.Body).Decode(&p1); err != nil {
		log.Println("Decode bad data", err)
	}

	log.Println(p1)
}
