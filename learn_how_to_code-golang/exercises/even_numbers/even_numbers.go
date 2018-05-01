// Prints all even numbers between 0 and 100
package main

import "fmt"

func main() {
	for i := 0; i < 101; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		} else {
			continue
		}
	}
	fmt.Println()
}
