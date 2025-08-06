package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)

func main() {
	if len(os.Args) > 1 { // Read code from file
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		repl.Start(file, os.Stdout) // because in is an io.Reader we can just pass a file to the existing REPL function
		return
	} else { // Start REPL
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
