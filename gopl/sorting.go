//Go provides the sort.Interface interface
//It exposes methods
//Len() int
//Less(i, j int) bool
//Swap(i, j int)
package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main()  {
  s := StringSlice{"michae", "ana", "nicolas", "gabi"}
  sort.Sort(s)
  fmt.Printf("Resulting %s\n", s)

  s = []string{"michae", "ana", "nicolas", "gabi"}
  sort.Strings(s)
  fmt.Printf("Resulting %s\n", s)
}
