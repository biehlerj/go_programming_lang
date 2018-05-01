// Package foo takes in a variadic number of arguments
package foo

import "fmt"

func foo(nums ...int) {
	fmt.Println(nums)
}
