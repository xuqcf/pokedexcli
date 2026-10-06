package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	temp_string := strings.ToLower(text)
	words := strings.Fields(temp_string)
	return words
}
