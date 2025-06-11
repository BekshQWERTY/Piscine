package main

import (
	"io"
	"os"
)

func main() {
	exitCode := 0
	defer func() { os.Exit(exitCode) }()

	args := os.Args[1:]

	if len(args) == 0 {

		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			exitCode = 1
		}
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			exitCode = 1
			return
		}

		_, err = io.Copy(os.Stdout, file)
		file.Close()
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			exitCode = 1
			return
		}
	}
}
