// Sorting a group of people using the built in sort method
package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println(studyGroup)
}
