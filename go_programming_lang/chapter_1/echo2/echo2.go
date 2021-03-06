// Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds()) // added for exercise 1.3
}
