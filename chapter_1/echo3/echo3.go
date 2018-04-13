// Echo3 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds()) // added for exercise 1.3
}
