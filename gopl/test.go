package main

import "fmt"

type Ar interface {
	Foo()
}

type A struct {
	s string
}

func (a *A) Foo() {
	fmt.Printf("Foo...%s\n", a.s)
}

func RunFoo(a Ar) {
	a.Foo()
}

func main() {
	a := &A{"deu"}
	RunFoo(a)

	var aa Ar
	fmt.Println(aa)
}
