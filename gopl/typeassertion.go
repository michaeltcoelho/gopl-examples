package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	// Since we've assigned os.Stdout whose dynamic type is *os.File
	w = os.Stdout
	// w's dynamic type is identical to *os.File
	// f contains w's dynamic value
	f := w.(*os.File)
	fmt.Fprintf(f, "f is w's extracted dynamic value since it is identical to %T\n", f)

	// Since w's dynamic type is *os.File it is not equal to *bytes.Buffer
	c, ok := w.(*bytes.Buffer)
	fmt.Printf("%T is %s\n", c, ok)

	// io.ReadWriteCloser matches io.Writer interface type and also, it includes two more methods Read and Close
	var rwc io.Writer
	rwc = os.Stdout
	b, ok := rwc.(io.ReadWriteCloser)
	fmt.Printf("%#v is %s\n", b, ok)
	// rwc dynamic type matches io.ReadWriteCloser then b contains io.ReadWriteCloser extra behavior
	b.Close()
}
