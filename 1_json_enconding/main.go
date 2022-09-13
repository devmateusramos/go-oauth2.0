package main

import (
	"encoding/json"
	"fmt"
	"log"
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
}

type person struct {
	First string
}
