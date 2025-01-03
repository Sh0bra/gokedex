package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Pokemon name: ")
		text, _ := reader.ReadString('\n')
		cleaned := cleanInput(text)
		if text == "exit\n" {
			fmt.Println("Thank you for using gokedex!")
			break
		}
		fmt.Println(cleaned)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
