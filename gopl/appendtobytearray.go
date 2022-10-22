package main

import "fmt"

func main() {
	a := []byte("ba")
	fmt.Printf("%s - %#v -> len(%d) - cap(%d)\n", a, a, len(a), cap(a))

	a1 := append(a, 'd')
	fmt.Printf("%s - %#v -> len(%d) - cap(%d)\n", a1, a1, len(a1), cap(a1))

	a2 := append(a, 'g')
	fmt.Printf("%s - %#v -> len(%d) - cap(%d)\n", a2, a2, len(a2), cap(a2))
}
