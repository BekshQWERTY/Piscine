package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args

	for i := range arg {
		if i != 0 {
			for j := range arg[i] {
				runes := []rune(arg[i])
				{
					z01.PrintRune(runes[j])
				}
			}
			z01.PrintRune('\n')
		}
	}
}
