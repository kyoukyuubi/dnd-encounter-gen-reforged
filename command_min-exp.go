package main

import (
	"fmt"
	"strconv"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
)

func commandMinExp(cfg *Config, args ...string) error {
	// check if args are empty, if so print the current value
	if len(args) == 0 {
		// check if MinExperience is set to 0
		// if it is, print a msg telling the user that no minimum exp is set
		if cfg.Config.MinExperience == 0 {
			fmt.Println("No ammount of Minimum Experience set")
			fmt.Println("")
		} else {
			fmt.Printf("Minimum Experience is set to: %d\n", cfg.Config.MinExperience)
			fmt.Println("")
		}
		return nil
	}

	// convert args[0] to int
	inputInt, err := strconv.Atoi(args[0])
	if err != nil {
		errorhandling.LogError(err, "commandMinExp")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return nil
	}

	// check to see if the input is 0 or below
	// if so, reset the value and return
	if inputInt == 0 || inputInt < 0 {
		fmt.Println("Minimum Experience has been reset")
		cfg.Config.MinExperience = 0
		return nil
	}

	// calc the budget
	budget, err := calcBudget(cfg)
	if err != nil {
		errorhandling.LogError(err, "commandMinExp/calcBudget")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return nil
	}

	// see if the inputtet number is bigger than the budget
	// if it is, warn the user and set to min to the budget. If not, set to the inputtet number
	if inputInt > budget {
		fmt.Println("Your input is larger than your budget, setting Minimum Experience to budget.")
		fmt.Println("")
		fmt.Printf("Input: %d\n", inputInt)
		fmt.Printf("Your current budget: %d\n", budget)
		fmt.Println("")

		// set the nim-exp to budget to prevent the user from going over
		cfg.Config.MinExperience = budget
	} else {
		fmt.Printf("Setting Minimum Experience set to: %d\n", inputInt)
		fmt.Println("")
		cfg.Config.MinExperience = inputInt
	}

	return nil
}
