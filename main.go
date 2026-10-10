package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	initConfig := &config{make_map()}
	repl(initConfig)
}

func repl(initConfig *config) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		cmd, ok := initConfig.registry[words[0]]
		if ok {
			err := cmd.callback(initConfig)
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

func commandExit(initConfig *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error //cliCommand.callback field
}

func commandHelp(initConfig *config) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n")
	fmt.Print("\n")
	for _, v := range initConfig.registry {
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

type config struct {
	registry map[string]cliCommand //holds a commands map
}
