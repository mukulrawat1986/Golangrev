// Package stringutil contains utility function for working with strings
package stringutil

// Reverse returns its argument string reversed rune-wise left to right
func Reverse(s string) string {
	return reverseTwo(s)
}
