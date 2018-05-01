// Prompts the user for a small number and larger number
// and prints the remainder of division of the numbers.
package main

import "fmt"

func main() {
	var smallNumber int
	var largeNumber int
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&smallNumber)
	fmt.Print("Please enter a larger number: ")
	fmt.Scan(&largeNumber)
	fmt.Println(largeNumber, "%", smallNumber, "=", largeNumber%smallNumber)
}
