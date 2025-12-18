package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		// cleaned := strings.ToLower(strings.TrimSpace(input))
		words := cleanInput(input)

		if len(words) > 0 {
			command := words[0]
			fmt.Printf("Your command was: %s\n", command)
		}
	}
}
