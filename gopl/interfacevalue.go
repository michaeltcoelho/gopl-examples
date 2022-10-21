// Interface value

// An interface has a pointer and a value
// Whichever concrete type which satifies an error interface is an error
package main

import "fmt"

type customError struct {
	msg string
	ctx map[string]string
}

func (ce customError) Error() string {
	return fmt.Sprintf("Error: %s", ce.msg)
}

func cError(msg string) *customError {
	return &customError{}
}

func main() {
	// if cError returns &customError{}, the interface value type is not nil, but its value is, it points to an empty struct, customError{}, which means next validation != nil will return true.
	// if cError returns nil, the interface value type is nil, its value is also nil then next validation != nil returns false
	var err error = cError("Blah")
	if err != nil {
		println("Works.")
	} else {
		println("Ops.")
	}
}
