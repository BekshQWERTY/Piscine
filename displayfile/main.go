package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Stdout.WriteString("File name missing\n")
		return
	} else if len(args) > 1 {
		os.Stdout.WriteString("Too many arguments\n")
		return
	}
	file, err := os.Open(args[0])
	if err != nil {
		os.Stdout.WriteString("Error opening file\n")
		return
	}
	defer file.Close()
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		os.Stdout.WriteString("Error reading file\n")
		return
	}
}
