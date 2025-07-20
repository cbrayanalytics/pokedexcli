package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	splittext := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return splittext
}

func main() {
	TestString := "  bulbasaur     pikachu  "

	result := cleanInput(TestString)
	fmt.Printf("Result of cleanInput: %s", result)
}
