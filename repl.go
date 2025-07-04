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
		commandName, args := cleanInput(reader.Text())

		// skip empty input
		if commandName == "" {
			continue
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

func cleanInput(text string) (string, []string) {
	// trim the whitesapce around the text if any
	text = strings.TrimSpace(text)

	// split the first space, to seperate command from args
	parts := strings.SplitN(text, " ", 2)

	// check if args are supplied, if not return the command
	if len(parts) == 1 {
		return parts[0], []string{}
	}

	// store the command and args in 2 seperate values for easier use
	command := strings.ToLower(parts[0])
	argsString := parts[1]

	// split the args by comma and trim white space
	var args []string

	if argsString != "" {
		// split it by ","
		rawArgs := strings.Split(argsString, ",")

		// loop through and add them to the empty slice
		for _, arg := range rawArgs {
			args = append(args, strings.ToLower(strings.TrimSpace(arg)))
		}
	}

	return command, args
}

// the command struct, stores data about the command and the function, using the signature
type cliCommand struct {
	name        string
	description string
	example     string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	// returns a map that stores the command struct with the data and the function name
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this help menu.",
			example:     "help",
			callback:    commandHelp,
		},
		"list": {
			name:        "list <what to list e.g 'planes'",
			description: "Lists available values for use in filters. For example, `list types` shows all valid creature types.",
			example:     "list planes",
			callback:    commandList,
		},
		"list-filters": {
			name:        "list-filters",
			description: "Lists all currently active filters and their values.",
			example:     "list-filters",
			callback:    commandListFilters,
		},
		"level": {
			name:        "level <number (optional)>",
			description: "Sets the desired level. If no level is specified, prints the current level.",
			example:     "level 10",
			callback:    commandLevel,
		},
		"party-size": {
			name:        "party-size <number (optional)>",
			description: "Sets the party size. If no size is specified, prints the current party size.",
			example:     "party-size 6",
			callback:    commandPartySize,
		},
		"max-creatures": {
			name:        "max-creatures <number (optional)>",
			description: "Sets the maximum number of creatures allowed (use 0 for no limit). If no amount is specified, prints the current maximum.",
			example:     "max-creatures 5",
			callback:    commandMaxCreatures,
		},
		"min-exp": {
			name:        "min-exp <number (optional)>",
			description: "Sets the minimum experience a creature must have (use 0 to disable). If no amount is specified, prints the current minimum.",
			example:     "min-exp 250",
			callback:    commandMinExp,
		},
		"type": {
			name:        "type <types seperated by comma (optional)>",
			description: "Sets one or more desired creature types (separated by commas; default is any). If none are given, prints the current types. To reset, use `any`.",
			example:     "type aberration, ooze",
			callback:    commandType,
		},
		"plane": {
			name:        "plane <planes seperated by comma (optional)>",
			description: "Sets one or more planes (comma-separated; default is any). If none are given, prints the current planes. To reset, use `any`.",
			example:     "plane feywilds, lower planes",
			callback:    commandPlane,
		},
		"environment": {
			name:        "environment <environments seperated by comma (optional)>",
			description: "Sets one or more environments (comma-separated; default is any). If none are given, prints the current environments. To reset, use `any`.",
			example:     "environment urban, hill",
			callback:    commandEnvironment,
		},
		"source": {
			name:        "source <sources seperated by comma (optional)>",
			description: "Sets one or more sources (comma-separated; default is any). If none are given, prints the current sources. To reset, use `any`.",
			example:     "source Monster Manual 2024",
			callback:    commandSource,
		},
		"difficulty": {
			name:        "difficulty <desired difficulty(optional)>",
			description: "Sets the desired encounter difficulty (default is Moderate). If none is specified, prints the current difficulty.",
			example:     "difficulty high",
			callback:    commandDifficulty,
		},
		"generate": {
			name:        "generate",
			description: "Generates a random encounter based on your current filters and settings. Each use creates a new encounter (previous results are not saved).",
			example:     "generate",
			callback:    commandGenerate,
		},
		"reset": {
			name:        "reset",
			description: "Resets all filters to their default values.",
			example:     "reset",
			callback:    commandReset,
		},
		"exit": {
			name:        "exit",
			description: "Safely exits the program and saves your filters.",
			example:     "exit",
			callback:    commandExit,
		},
	}
}
