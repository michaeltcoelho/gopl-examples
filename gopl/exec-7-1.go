package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WordCounter int
type LineCounter int

func Counter(sf bufio.SplitFunc, s string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(s)))
	scanner.Split(sf)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func (cw *WordCounter) Write(p []byte) (int, error) {
	count, err := Counter(bufio.ScanWords, string(p))
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*cw += WordCounter(count)
	return count, nil
}

func (cl *LineCounter) Write(p []byte) (int, error) {
	count, err := Counter(bufio.ScanLines, string(p))
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*cl += LineCounter(count)
	return count, nil
}

func main() {
	text := "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

	var cw WordCounter
	cw.Write([]byte(text))
	fmt.Println(cw)

	var cl LineCounter
	cl.Write([]byte(text))
	fmt.Println(cl)
}
