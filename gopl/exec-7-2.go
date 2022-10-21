package main

import (
	"bytes"
	"fmt"
	"io"
)

type MyWriter struct {
	w io.Writer
	b int64
}

func (c *MyWriter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.b += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &MyWriter{w: w, b: 0}
	return c, &c.b
}

func main() {
	b := &bytes.Buffer{}
	w, n := CountingWriter(b)
	data := []byte("Hello, world!")
	w.Write(data)

	fmt.Printf("CountingWriter %d\n", *n)
	fmt.Printf("Data %d\n", int64(len(data)))
}
