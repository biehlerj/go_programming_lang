// Implementing FizzBuzz in Go
package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		result := ""
		if i%3 == 0 {
			result += "Fizz"
		}
		if i%5 == 0 {
			result += "Buzz"
		}
		if result != "" {
			fmt.Print(result)
			if i != 100 {
				fmt.Print(" ")
			}
			continue
		}
		fmt.Print(i)
		if i != 100 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
