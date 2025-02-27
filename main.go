package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)

		if len(cleanedInput) > 0 {
			fmt.Println("Your command was:", cleanedInput[0])
		}
	}
}

func cleanInput(text string) []string {
	var result []string
	lower := strings.ToLower(text)
	trim := strings.TrimSpace(lower)
	result = strings.Fields(trim)
	return result
}
