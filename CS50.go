package main

import (
	"bufio"
	"os"
	"strings"
)

// getLine prompts the user for a line of text from standard input and returns it as a string,
// sans trailing line ending. It supports CR (\r), LF (\n), and CRLF (\r\n) as line endings.
func getLine(prompt string) string {
	// Prompt user
	if prompt != "" {
		print(prompt)
	}

	// Read a line from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		// Trim trailing newline characters
		line = strings.TrimRight(line, "\r\n")
		return line
	}

	// If no input is provided, return an empty string
	return ""
}
