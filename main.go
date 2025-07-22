// Package main demonstrates basic string cleaning and formatting in Go.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cleanInput takes a string, trims leading/trailing whitespace,
// converts it to lowercase, and splits it by any whitespace.
// It returns a slice of cleaned, lowercase words.
//
// Example:
//
//	input: "  BULBASAUR   pikachu  "
//	output: ["bulbasaur", "pikachu"]
func cleanInput(text string) []string {
	splittext := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return splittext
}
bootdev run 2d4ed658-b546-4bf1-b569-57e5e321c1ef -s
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		for scanner.Scan() {
			input := cleanInput(scanner.Text())
			fmt.Printf("Your command was: %s\n", input[0])
			break
		}

	}
}
