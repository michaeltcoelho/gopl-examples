package main

import (
    "os"
    "fmt"
)

func main() {
    filename := os.Args[1]

    fmt.Printf("%s\n", basename(filename))
}

func basename(filename string) string {
    for i := len(filename) - 1; i > 0; i-- {
        if filename[i] == '/' {
            filename = filename[i+1:]
            break
        }
    }

    for i := len(filename) - 1; i > 0; i-- {
        if filename[i] == '.' {
            filename = filename[:i]
            break
        }
    }
    return filename
}
