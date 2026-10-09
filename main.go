package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	newMap := make_map()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		cmd, ok := newMap[words[0]]
		if ok {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n")
	fmt.Print("\n")
	for _, v := range make_map() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func make_map() map[string]cliCommand {
	exitMap := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	return exitMap
}
