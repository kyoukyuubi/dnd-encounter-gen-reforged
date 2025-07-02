package main

import (
	"fmt"
	"strings"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func commandEnvironment(cfg *Config, args ...string) error {
	// if no args are provided display what is in the config
	if len(args) == 0 {
		// check if the store environments are empty
		if len(cfg.Config.Environments) == 0 {
			fmt.Println("Environments are set to any")
		} else {
			fmt.Println("Environments are set to: ")

			// loop through the stored types are print them in an unorded list
			for _, e := range cfg.Config.Environments {
				fmt.Printf("* %s\n", e)
			}
		}

		fmt.Println("")
		return nil
	}

	// if args is set to 'Any' reset the config to an empty slice
	if args[0] == "any" {
		fmt.Println("Environments has been reset to default value")
		fmt.Println("")
		cfg.Config.Environments = []string{}
		return nil
	}

	// load the environments from the json file
	environments, err := jsonHandler.LoadEnvirnments()
	if err != nil {
		errorhandling.LogError(err, "commandEnvironment/LoadEnvirnments")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return nil
	}

	// create an empty slice for putting the found values into
	var sliceFound []string

	// set a value to false, to be used after the loop to see if formatting needs to be printed
	errorOccurred := false

	// loop through the user inputs and see if they match the environments from the json
	for _, input := range args {
		matchFound := false
		for _, e := range environments {
			// check if what was inputtet matches what was in the json file
			// uses the value from the json file for error handling etc.
			if strings.Contains(strings.ToLower(e), input) {
				sliceFound = append(sliceFound, e)
				matchFound = true
				continue
			}
		}

		// if no matchers is found tell the user
		if !matchFound {
			fmt.Printf("%s was not found in the determined environments\n", input)
			errorOccurred = true
		}
	}

	// check if error occurred, if so print a line break for the user to make things easier to read
	if errorOccurred {
		fmt.Println("-------------------------------")
	}

	// update the config with the new slice if the new slice list is not empty
	if len(sliceFound) != 0 {
		cfg.Config.Environments = sliceFound
		fmt.Println("the environments are now set to: ")

		// loop through the created slice and print it in an unorederd list format
		for _, e := range sliceFound {
			fmt.Printf("* %s\n", e)
		}
		fmt.Println("")
	} else {
		fmt.Println("No matches were found in the determined types")
		fmt.Println("Environments list was not altered")
		fmt.Println("")
	}

	return nil
}