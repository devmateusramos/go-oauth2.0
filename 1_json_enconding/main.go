package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))

	p1Json, _ := json.Marshal(p1)
	fmt.Println(p1Json)
	fmt.Println(string(p1Json))

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

	http.ListenAndServe(":8080", nil)
}

type person struct {
	First string
}

func foo(w http.ResponseWriter, r *http.Request) {

}

func bar(w http.ResponseWriter, r *http.Request) {

}
