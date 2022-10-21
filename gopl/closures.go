package main

import "fmt"

func main() {
	// f is the closure itself
	f := foo()

	fmt.Println(f(), f(), f(), f())

	// We're simulating here a common issue that might happen to you when working with closures
	// As the close keeps reference of the outer's block variables
	// at the end of the loop, i will be 4, therefore, all closures will print out 4
	// In order to fix it, you just reassign i
	funcs := make([]func(), 4)
	for i := 0; i < 4; i++ {
		i := i
		funcs[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		funcs[i]()
	}
}

func foo() func() int {
	a, b := 0, 1

	// the closure keeps reference for the outer's function variables
	return func() int {
		a, b = b, a+b
		return b
	}
}
