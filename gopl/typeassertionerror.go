package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

var ErrNotExist = errors.New("file does not exist")

func IsNotExist(err error) bool {
	// type assertion for error matching
	// What's the dynamic type behind err (error) interface value?
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

// You can create structured errors with contextual data
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()

}

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err)         // open /no/such/file: no such file or directory
	fmt.Printf("%#v\n", err) // &fs.PathError{Op:"open", Path:"/no/such/file", Err:0x2}

	fmt.Println(IsNotExist(err)) // true
}
