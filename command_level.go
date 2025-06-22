package main

import (
	"fmt"
	"os"
	"strconv"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
)

func commandLevel(cfg *Config, args ...string) error {
	// check if args are empty, if so print the current value
	if len(args) == 0 {
		fmt.Printf("level is set to %d\n", cfg.Config.Level)
		return nil
	}

	// convert args[0] to int
	inputInt, err := strconv.Atoi(args[0])
	if err != nil {
		errorhandling.LogError(err, "commandLevel")
		os.Exit(1)
	}

	// check if the inputtet number is between 1 and 20
	if inputInt < 0 || inputInt > 20 {
		fmt.Println("level needs be between 1 and 20")
		return nil
	}

	// store the new number in the config and display confirmation msg
	fmt.Printf("level set to %d\n", inputInt)
	cfg.Config.Level = inputInt

	return nil
}