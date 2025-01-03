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
		fmt.Print("Gokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if text == "exit" {
			fmt.Println("Thank you for using gokedex!")
			break
		}
		fmt.Println("Your command was:", cleaned[0])
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
