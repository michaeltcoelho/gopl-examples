package main

import (
	"bytes"
	"fmt"
	"io"
)

type MyReader struct {
	s string
}

func (mr *MyReader) Read(p []byte) (n int, err error) {
	n = copy(p, mr.s)
	println(n)
	mr.s = mr.s[n:]
	if len(mr.s) == 0 {
		err = io.EOF
	}
	return
}

func main() {
	s := "hi there"

	b := &bytes.Buffer{}
	n, _ := b.ReadFrom(&MyReader{s})

	fmt.Printf("Read from MyReader, %d\n", n)
	fmt.Printf("s bytes, %d\n", len(s))
}
