// Generic word count program
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type OperationResult int
type Operation func(arg int) OperationResult

// readArgsFromStdin - reads args from standard input
func readArgsFromStdin() []int {
	args := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		err := scanner.Scan()
		fmt.Println(err)

		text := scanner.Text()
		if len(text) < 1 {
			break
		}
		for _, arg := range strings.Split(text, " ") {
			if item, err := strconv.Atoi(arg); err == nil {
				args = append(args, item)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return args
}

func main() {
	args := readArgsFromStdin()

	fmt.Println(args)

	var b byte = 32
	a := "ass"

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
