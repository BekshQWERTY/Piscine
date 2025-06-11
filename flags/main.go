package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	insertStr := ""
	order := false
	help := false
	mainStr := ""

	// Parse flags
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if len(arg) > 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--insert" || arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if arg == "--help" || arg == "-h" {
			help = true
		} else if mainStr == "" {
			mainStr = arg
		}
	}

	// Show help if needed
	if help || len(args) == 0 {
		printHelp()
		return
	}

	// Apply insert
	result := mainStr + insertStr

	// Apply order
	if order {
		runes := []rune(result)
		// Bubble sort implementation
		for i := 0; i < len(runes); i++ {
			for j := 0; j < len(runes)-1; j++ {
				if runes[j] > runes[j+1] {
					runes[j], runes[j+1] = runes[j+1], runes[j]
				}
			}
		}
		result = string(runes)
	}

	// Print result
	printStr(result)
	z01.PrintRune('\n')
}

func printHelp() {
	printStr("--insert\n")
	printStr("  -i\n")
	printStr("\t This flag inserts the string into the string passed as argument.\n")
	printStr("--order\n")
	printStr("  -o\n")
	printStr("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
