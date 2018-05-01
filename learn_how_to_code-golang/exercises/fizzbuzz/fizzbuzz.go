// Implementing FizzBuzz in Go
package main

import "fmt"

func main() {
	fizzString := "Fizz"
	buzzString := "Buzz"
	for i := 1; i < 101; i++ {
		if i%15 == 0 {
			fmt.Print(fizzString + buzzString)
		} else if i%3 == 0 {
			fmt.Print(fizzString)
		} else if i%5 == 0 {
			fmt.Print(buzzString)
		} else {
			fmt.Printf("%d", i)
		}
		if i != 100 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
