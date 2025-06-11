package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	upper := false
	start := 1
	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		upper = true
		start = 2
	}

	for i := start; i < len(os.Args); i++ {
		n := 0
		valid := true
		for _, c := range os.Args[i] {
			if c < '0' || c > '9' {
				valid = false
				break
			}
			n = n*10 + int(c-'0')
		}

		if valid && n > 0 && n < 27 {
			if upper {
				z01.PrintRune(rune('A' + n - 1))
			} else {
				z01.PrintRune(rune('a' + n - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}

	if len(os.Args) > start {
		z01.PrintRune('\n')
	}
}
