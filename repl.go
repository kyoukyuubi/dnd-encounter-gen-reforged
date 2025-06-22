package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"
)

// the config struct, that usually stores a timeout etc. incase you use an api
// here we are storing the JsonConfig data
type Config struct {
	Config config.JsonConfig
}

func startRepl(config *Config) {
	// make a new reader using bufio to make a cli program
	reader := bufio.NewScanner(os.Stdin)

	for {
		// text to show while waiting for user input
		fmt.Printf("D&D Encounter Gen > ")

		// wait for user input
		reader.Scan()

		// store what was written after going through the "cleanInput" function
		words := cleanInput(reader.Text())

		// skip empty input
		if len(words) == 0 {
			continue
		}

		// parse command structure: first word is command, rest are arguments  
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		// check if the command exists by calling the "GetCommands" function
		command, exists := getCommands()[commandName]
		if exists {
			// if it does, use the command and input the config and the args, handle errors
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			// tell the loop to restart
			continue
		} else {
			// command not found, tell the user and continue
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	// normalize input: lowercase and split on whitespace for consistent parsing
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// the command struct, stores data about the command and the function, using the signature
type cliCommand struct {
	name string
	description string
	callback func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	// returns a map that stores the command struct with the data and the function name
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "displays this menu",
			callback: commandHelp,
		},
		"level": {
			name: "level <number (optional)>",
			description: "set the desired level, default is 1. If no level is specified, prints the current set level",
			callback: commandLevel,
		},
		"exit": {
			name: "exit",
			description: "safely exit the program",
			callback: commandExit,
		},
	}
}