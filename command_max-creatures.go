package main

import (
	"fmt"
	"strconv"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
)

func commandMaxCreatures(cfg *Config, args ...string) error {
	// check if args are empty, if so print the current value
	if len(args) == 0 {
		// check if the MaxCreatures is 0
		// if it is, print a msg saying that it's set to any, if not, print current ammount
		if cfg.Config.MaxCreatures == 0 {
			fmt.Println("No ammount of Maximum Creatures set")
			fmt.Println("")
		} else {
			fmt.Printf("Max Creatures is set to: %d\n", cfg.Config.MaxCreatures)
			fmt.Println("")
		}
		return nil
	}

	// convert args[0] to int
	inputInt, err := strconv.Atoi(args[0])
	if err != nil {
		errorhandling.LogError(err, "commandLevel")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return nil
	}

	// check to see if the input is 0 or below
	// if so, reset the value and return
	if inputInt == 0 || inputInt < 0 {
		fmt.Println("Max Creatures has been reset")
		fmt.Println("")
		cfg.Config.MaxCreatures = 0
		return nil
	}

	// store the new max creatuers in the config and display confirmation msg
	fmt.Printf("Max Creatures set to: %d\n", inputInt)
	fmt.Println("")
	cfg.Config.MaxCreatures = inputInt

	return nil
}
