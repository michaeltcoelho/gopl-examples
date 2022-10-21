package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string = "Hello, there!"
	sub1 := s[6:]
	sub2 := s[:6]
	fmt.Printf("%v\n", &s)
	fmt.Printf("%v\n", &sub1)
	fmt.Printf("%v\n", &sub2)

	s = "hello"
	fmt.Println(hasPrefix(s, "he"))
	fmt.Println(hasPrefix(s, "lo"))
	fmt.Println(hasSuffix(s, "lo"))
	fmt.Println(hasSuffix(s, "he"))

	s = "Hello, そこ"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Printf("% X\n", s)
	for _, r := range s {
		fmt.Printf("%c\n", r)
	}

	fmt.Printf("%T\n", rune('a'))
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func hasSuffix(s, suf string) bool {
	return len(s) >= len(suf) && s[len(s)-len(suf):] == suf
}
