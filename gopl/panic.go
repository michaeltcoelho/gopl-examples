package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// defer printStack()
	// defer func() {
	// 	if p := recover(); p != nil {
	// 		fmt.Printf("internal error: %v", p)
	// 	}
	// }()

	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) //panic if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
