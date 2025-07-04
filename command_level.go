package main

import (
	"fmt"
	"strconv"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
)

func commandLevel(cfg *Config, args ...string) error {
	// check if args are empty, if so print the current value
	if len(args) == 0 {
		fmt.Printf("level is set to: %d\n", cfg.Config.Level)
		fmt.Println("")
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

	// check if the inputtet number is between 1 and 20
	if inputInt < 0 || inputInt > 20 {
		fmt.Println("level needs be between 1 and 20")
		fmt.Println("")
		return nil
	}

	// store the new level in the config and display confirmation msg
	fmt.Printf("level set to: %d\n", inputInt)
	fmt.Println("")
	cfg.Config.Level = inputInt

	return nil
}
