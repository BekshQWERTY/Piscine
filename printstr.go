package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, char := range s {
		if char == '\n' {
			z01.PrintRune('\\')
			z01.PrintRune('n')
		} else {
			z01.PrintRune(char)
		}
	}
}
