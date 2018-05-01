// Package greatestNumber is a variadic function that takes in an array of numbers
// Searches for and returns the largest number in the list
package greatestNumber

func greatestNumber(nums ...int) int {
	var largest int
	for _, v := range nums {
		if v > largest {
			largest = v
		}
	}
	return largest
}
