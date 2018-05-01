// Package half ivides a number by two
// Returns whether or not the argument is even.
package half

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
