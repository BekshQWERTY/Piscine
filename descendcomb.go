package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := 99
	last := 0
	firstPrinted := false

	// Iterate through the range for the first number i
	for i := first; i >= 1; i-- {
		// Iterate through the range for the second number j, smaller than i
		for j := i - 1; j >= last; j-- {
			// Check if this is not the first combination
			if firstPrinted {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			// if i < 10 {
			// 	z01.PrintRune('0')
			// }
			printTwoDigits(i)  // Print the first number i
			z01.PrintRune(' ') // Space between the two numbers
			printTwoDigits(j)  // Print the second number j
			firstPrinted = true
		}
	}
}

// Helper function to print a two-digit number with leading zeros if necessary
func printTwoDigits(n int) {
	// Check if the number is less than 10 and add a leading zero if necessary
	// Print the tens place
	z01.PrintRune(rune('0' + n/10))
	// Print the ones place
	z01.PrintRune(rune('0' + n%10))
}
