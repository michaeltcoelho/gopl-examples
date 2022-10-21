package main

import (
	"encoding/json"
	"fmt"
)

type album1 struct {
	Title string `json:"title"`
}

type album2 struct {
	Title string `json:"title"`
}

func main() {
	fmt.Println("Hello, world!")

	a0 := album1{}

	a1 := album1{
		"Title 1",
	}

	a2 := album2{
		"Title 2",
	}

	fmt.Printf("%T %#[1]v\n", a0)
	fmt.Printf("%T %#[1]v\n", a1)
	fmt.Printf("%T %#[1]v\n", a2)

	payload, _ := json.Marshal(a1)
	fmt.Printf("%T %[1]v\n", string(payload))

	var a3 album1
	json.Unmarshal(payload, &a3)
	fmt.Printf("%T %#[1]v\n", a3)

	a1 = album1(a2)
	fmt.Printf("%T %#[1]v\n", a1)
}
