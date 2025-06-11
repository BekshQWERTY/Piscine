package main

import "github.com/01-edu/z01"

func main() {
	for i := 25; i >= 0; i-- {
		z01.PrintRune(rune('a' + i))
	}
	z01.PrintRune('\n')
}
