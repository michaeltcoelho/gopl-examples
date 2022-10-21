package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a []int
	fmt.Println(a)
	fmt.Println(a == nil)
	fmt.Printf("%T\n", []int(nil))
	fmt.Printf("%b\n", []int(nil) == nil)

	// changing the underlining array for arr slice
	arr := []int{0, 1, 2, 3, 4, 5}
	reverse(arr[:2])
	reverse(arr[2:])
	reverse(arr)
	fmt.Println(arr)

	//append
	var runes []rune
	for _, r := range "Hello, ちは!" {
		fmt.Printf("Rune Count:. %v\n", utf8.RuneLen(r))
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	fmt.Printf("%v\n", len(runes))

	//appendInt
	//Growing a slice dynamically
	var k, l []int
	for i := 0; i < 10; i++ {
		l = appendInt(k, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(l), l)
		k = l
	}

	// array och slices
	const aa int = 1
	fmt.Printf("%v\n", [5]int{1, 2, 3})
	fmt.Printf("%v\n", [...]int{1, 2, 3})
	fmt.Printf("%v\n", [3]string{1: "a", 2: "b"})
	fmt.Printf("%v\n", []string{1: "a", 2: "b"})
	fmt.Printf("%v\n", []string{aa: "a", 2: "b"})

	// map
	fmt.Printf("%v\n", map[string]int{"a": 1, "b": 2})
	fmt.Printf("%v\n", map[string]struct{ string }{"a": {"adsf"}, "b": {"fdads"}})
	fmt.Printf("%v\n", map[string]struct{ string }{"a": {"adsf"}, "b": {"fdads"}})

	type language struct {
		name string
		year int
	}

	var _ = [...]language{
		language{"C", 1972},
		language{"Python", 1991},
		language{"Go", 2009},
	}
	// simplifying map literals
	var _ = [...]language{
		{"C", 1972},
		{"Python", 1991},
		{"Go", 2009},
	}

	// Comparing container values
	var b [16]byte
	var c []int
	var mm map[string]int

	fmt.Println(b == b)
	fmt.Println(mm == nil)
	fmt.Println(c == nil)
	fmt.Println(nil == map[string]int{})
	fmt.Println(nil == []int{})
	fmt.Println(nil == []int(nil))
	fmt.Printf("%T\n", []map[string]int{})

	s0 := []int{1, 2, 3}
	fmt.Println(s0, cap(s0))
	s1 := append(s0, 4)
	fmt.Println(s1, cap(s1)) // It updates cap to 6, double the base-slic length. Depends on the compiler

	//Besides composite literals to create maps and slices and can use make(M, n)
	m0 := make(map[string]int, 3)
	fmt.Println(m0, len(m0))
	m0["Go"] = 1
	m0["C"] = 2
	fmt.Println(m0, len(m0))

	ms0 := make([]int, 3, 5)
	fmt.Println(ms0, len(ms0), cap(ms0))
	ms0 = make([]int, 2)
	fmt.Println(ms0, len(ms0), cap(ms0))

	//You wont be able to
	type T struct{ v int }
	m1 := map[string]T{}
	m1["Python"] = T{v: 3}
	m1["Go"] = T{v: 2}

	fmt.Println(m1)
	//Wont work cause you cant change struct field in a map
	//You have to modify it as a whole
	//m1["Go"].v = 5

	m1["Go"] = T{v: 5}

	fmt.Println(m1)
}

func reverse(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

//appendInt grows a slace dynamically by allocating new array segment whose capacity is doubled the original array's length
func appendInt(s []int, i int) []int {
	var z []int
	zlen := len(s) + 1
	if zlen <= cap(s) {
		// There are still space for growing s. Then let's grow it.
		z = s[:zlen]
	} else {
		//There's not enough space. Let's allocate new segment
		// It creates new segment with capacity doubled the original array's length
		zcap := zlen
		if zcap < 2*len(s) {
			zcap = 2 * len(s)
		}
		z = make([]int, zlen, zcap)
		copy(z, s)
	}
	z[len(s)] = i
	return z
}
