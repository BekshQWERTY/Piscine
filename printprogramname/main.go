package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programPath := os.Args[0]
	lastSlash := 0
	for i, c := range programPath {
		if c == '/' {
			lastSlash = i + 1
		}
	}
	programName := programPath[lastSlash:]
	for _, r := range programName {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
