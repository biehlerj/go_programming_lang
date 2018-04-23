// Myecho1 prints the name of the command that invoked it as well as whatever args are passed to it
// Corresponds to exercise 1.1 (pg 8)
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0], strings.Join(os.Args[1:], " "))
}
