// Finds the difference between the sum of squares and square of sum
// of the first 100 natural numbers
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumSquare())
}

func sumSquare() int {
	var squaresFloat float64
	var sumsFloat float64
	for i := 1.0; i < 101; i++ {
		squaresFloat += math.Pow(i, 2)
		sumsFloat += i
	}
	sumsFloat = math.Pow(sumsFloat, 2)
	sums := int(sumsFloat)
	squares := int(squaresFloat)
	return sums - squares
}

/*
Finds the difference between the sums of squares and square of sums
for all natural numbers between 1 and 100
*/
