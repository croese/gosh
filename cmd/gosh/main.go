package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const defaultPrompt = "gosh>"

func main() {
	fmt.Printf("%s ", defaultPrompt)
	input, err := readUserInput()
	if err != nil {
		reportError(err)
	}
	fmt.Printf("gosh: got input: '%s'\n", input)
}

func reportError(err error) {
	fmt.Fprintf(os.Stderr, "gosh: input error: %s\n", err)
}

func readUserInput() (string, error) {
	stdin := bufio.NewScanner(os.Stdin)
	input := ""
	for stdin.Scan() {
		line := stdin.Text()
		if !strings.HasSuffix(line, "\\") {
			input += line
			break
		} else {
			input += strings.TrimSuffix(line, "\\")
			fmt.Print("> ")
		}
	}
	return input, stdin.Err()
}
