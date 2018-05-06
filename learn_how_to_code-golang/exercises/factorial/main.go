// Finding the factorial of a number using go routines and channels
package main

import "fmt"

func main() {
	c := factorial(33)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	factorial := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		factorial <- total
		close(factorial)
	}()
	return factorial
}
