package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Prompt asks a question on the command line with a default value
func Prompt(reader *bufio.Reader, question, defaultValue string) string {
	fmt.Printf("%s [%s]: ", question, defaultValue)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if text == "" {
		text = defaultValue
	}
	return strings.ToLower(text)
}
