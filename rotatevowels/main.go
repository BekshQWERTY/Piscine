package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	var allRunes []rune
	for i, arg := range os.Args[1:] {
		if i > 0 {
			allRunes = append(allRunes, ' ')
		}
		allRunes = append(allRunes, []rune(arg)...)
	}

	var vowels []rune
	var positions []int
	for pos, char := range allRunes {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
			char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			vowels = append(vowels, char)
			positions = append(positions, pos)
		}
	}

	for i := 0; i < len(vowels)/2; i++ {
		pos1 := positions[i]
		pos2 := positions[len(vowels)-1-i]
		allRunes[pos1], allRunes[pos2] = allRunes[pos2], allRunes[pos1]
	}

	for _, r := range allRunes {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
