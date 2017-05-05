package main

import "fmt"

func main() {
	p := new(int)
	fmt.Printf("%v", *p)
	*p = 2
	fmt.Printf("%v", *p)
}
