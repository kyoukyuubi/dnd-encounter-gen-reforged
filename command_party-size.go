package main

import (
	"fmt"
	"strconv"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
)

func commandPartySize(cfg *Config, args ...string) error {
	// check if args are empty, if so print the current value
	if len(args) == 0 {
		fmt.Printf("Party size is set to: %d\n", cfg.Config.NumPlayers)
		return nil
	}

	// convert args[0] to int
	inputInt, err := strconv.Atoi(args[0])
	if err != nil {
		errorhandling.LogError(err, "commandPartySize")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return nil
	}

	// check if the party-size is below 0
	if inputInt <= 0 {
		fmt.Println("Party Size must be above 0")
		return nil
	}

	// store the new party size in the config and display confirmation msg
	fmt.Printf("Party size set to: %d\n", inputInt)
	cfg.Config.NumPlayers = inputInt

	return nil
}
