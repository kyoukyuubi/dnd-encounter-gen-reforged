package main

import (
	"fmt"
	"strings"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func commandType(cfg *Config, args ...string) error {
	// if no args are provided display what is in the config
	if len(args) == 0 {
		// check if the stored types are empty
		if len(cfg.Config.Types) == 0 {
			fmt.Println("Types are set to any")
		} else {
			fmt.Println("Types are set to: ")

			// loop through the stored types and print them in an unordered list
			for _, t := range cfg.Config.Types {
				fmt.Printf("* %s\n", t)
			}
		}

		fmt.Println("")
		return nil
	}

	// if args is set to 'Any' reset the config to an empty slice
	if args[0] == "any" {
		fmt.Println("types has been reset to default value")
		fmt.Println("")
		cfg.Config.Types = []string{}
		return nil
	}

	// load the types from the json file
	types, err := jsonHandler.LoadTypes()
	if err != nil {
		errorhandling.LogError(err, "commandType")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return nil
	}

	// create an empty slice for putting the found values into
	var sliceTypes []string

	// set a value to false, to be used after the loop to see if formatting needs to be printed
	errorOccurred := false

	// loop through the user inputs and see if they match the types from the json
	for _, input := range args {
		matchFound := false
		for _, t := range types {
			// check if what was inputtet matches what was in the json file
			// uses the value from the json file for error handling etc.
			if strings.Contains(strings.ToLower(t), input) {
				sliceTypes = append(sliceTypes, t)
				matchFound = true
				continue
			}
		}

		// if no matches is found tell the user
		if !matchFound {
			fmt.Printf("%s was not found in the determined types\n", input)
			errorOccurred = true
		}
	}
	
	// check if error occurred, if so print a line break for the user to make things easier to read
	if errorOccurred {
		fmt.Println("-------------------------------")
	}

	// update the config with the new slice if the new slice list is not empty
	if len(sliceTypes) != 0 {
		cfg.Config.Types = sliceTypes
		fmt.Println("The types are now set to:")

		// loop through the created slice and print it in an unorederd list format
		for _, t := range sliceTypes {
			fmt.Printf("* %s\n", t)
		}
		fmt.Println("")
	} else {
		fmt.Println("No matches were found in the determined types")
		fmt.Println("types list was not altered")
		fmt.Println("")
	}

	return nil
}