package main

import (
	"io"
	"os"
	"strings"
)

type MyLimitReader struct {
	r        io.Reader
	limit, b int64
}

func (lr *MyLimitReader) Read(p []byte) (n int, err error) {
	n, err = lr.r.Read(p[:lr.limit])
	lr.b += int64(n)
	if lr.b >= lr.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &MyLimitReader{r: r, limit: n, b: 0}
}

func main() {
	s := "Hi there, my name is Michael. I'm from Brazil. Brazil is a large country located in South America."

	lr := LimitReader(strings.NewReader(s), 8)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		os.Exit(1)
	}
}
