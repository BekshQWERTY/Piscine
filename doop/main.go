package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	arg1 := os.Args[1]
	operator := os.Args[2]
	arg2 := os.Args[3]

	// Parse first number
	num1, ok1 := atoi(arg1)
	if !ok1 {
		return
	}

	// Parse second number
	num2, ok2 := atoi(arg2)
	if !ok2 {
		return
	}

	// Perform operation
	switch operator {
	case "+":
		result, overflow := add(num1, num2)
		if overflow {
			return
		}
		printInt(result)
	case "-":
		result, overflow := sub(num1, num2)
		if overflow {
			return
		}
		printInt(result)
	case "*":
		result, overflow := mul(num1, num2)
		if overflow {
			return
		}
		printInt(result)
	case "/":
		if num2 == 0 {
			printString("No division by 0\n")
			return
		}
		result, overflow := div(num1, num2)
		if overflow {
			return
		}
		printInt(result)
	case "%":
		if num2 == 0 {
			printString("No modulo by 0\n")
			return
		}
		result := num1 % num2
		printInt(result)
	default:
		return
	}
}

// atoi converts string to int, returns (value, success)
func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := 1
	i := 0

	// Handle sign
	if s[0] == '-' {
		sign = -1
		i = 1
	} else if s[0] == '+' {
		i = 1
	}

	if i >= len(s) {
		return 0, false
	}

	result := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		digit := int(s[i] - '0')

		// Check for overflow before multiplication
		if result > (9223372036854775807-digit)/10 {
			return 0, false
		}

		result = result*10 + digit
	}

	result *= sign

	// Check for underflow
	if sign == -1 && result > 0 {
		return 0, false
	}

	return result, true
}

// printInt prints an integer to stdout
func printInt(n int) {
	if n == 0 {
		os.Stdout.Write([]byte("0\n"))
		return
	}

	var digits []byte
	negative := n < 0

	if negative {
		n = -n
	}

	for n > 0 {
		digits = append(digits, byte(n%10)+'0')
		n /= 10
	}

	if negative {
		os.Stdout.Write([]byte("-"))
	}

	// Print digits in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		os.Stdout.Write([]byte{digits[i]})
	}
	os.Stdout.Write([]byte("\n"))
}

// printString prints a string to stdout
func printString(s string) {
	os.Stdout.Write([]byte(s))
}

// add performs addition with overflow detection
func add(a, b int) (int, bool) {
	result := a + b
	// Check for overflow
	if (b > 0 && a > 9223372036854775807-b) || (b < 0 && a < -9223372036854775808-b) {
		return 0, true
	}
	return result, false
}

// sub performs subtraction with overflow detection
func sub(a, b int) (int, bool) {
	result := a - b
	// Check for overflow
	if (b < 0 && a > 9223372036854775807+b) || (b > 0 && a < -9223372036854775808+b) {
		return 0, true
	}
	return result, false
}

// mul performs multiplication with overflow detection
func mul(a, b int) (int, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}

	result := a * b

	// Check for overflow by dividing back
	if result/b != a {
		return 0, true
	}

	return result, false
}

// div performs division with overflow detection
func div(a, b int) (int, bool) {
	// Check for the special case of dividing the minimum int by -1
	if a == -9223372036854775808 && b == -1 {
		return 0, true
	}

	return a / b, false
}
