package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"
)

type Config struct {
	Config config.JsonConfig
}

func startRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("D&D Encounter Gen > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name string
	description string
	callback func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "displays this menu",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "safely exit the program",
			callback: commandExit,
		},
	}
}