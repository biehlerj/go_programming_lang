// Myecho2 prints the index and value of each of its argument, one per line
// Corresponds to exercise 1.2 (pg 8)
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}
