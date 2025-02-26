package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var result []string
	lower := strings.ToLower(text)
	trim := strings.TrimSpace(lower)
	result = strings.Fields(trim)
	return result
}
