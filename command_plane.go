package main

import (
	"fmt"
	"strings"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func commandPlane(cfg *Config, args ...string) error {
	// if no args are provided, display what is in the config
	if len(args) == 0 {
		// check if the stored types are empty
		if len(cfg.Config.Planes) == 0 {
			fmt.Println("Planes are set to any")
		} else {
			fmt.Println("Planes are set to: ")

			// loop through the stored planes and print them in an unordered list
			for _, p := range cfg.Config.Planes {
				fmt.Printf("* %s\n", p)
			}
		}

		fmt.Println("")
		return nil
	}

	// if args is set to 'Any' rest the config to an empty slice
	if args[0] == "any" {
		fmt.Println("Planes has been reset to default value")
		fmt.Println("")
		cfg.Config.Planes = []string{}
		return nil
	}

	// load the planes from the json file
	planes, err := jsonHandler.LoadPlanes()
	if err != nil {
		errorhandling.LogError(err, "commandPlane/LoadPlanes")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return nil
	}

	// create an empty slice for the found matches
	var sliceFound []string

	// set a value to false, to be used after the loop to see if formatting needs to be printed
	errorOccurred := false

	// set previous category and subCategory for chekcs to avoid duplicate entries
	previousCategory := ""
	previousSubCategory := ""

	for _, input := range args {
		matchFound := false

		for _, value := range planes.Plane {

			// check if what was inputtet matches what was in the json file
			// uses the value from the json file for error handling etc.
			if strings.Contains(strings.ToLower(value.Name), input) {
				sliceFound = append(sliceFound, value.Name)
				matchFound = true
				continue
			} else if strings.Contains(strings.ToLower(value.Category), input) {

				// check if sliceFound contains the input, if it doesn't add it
				if jsonHandler.CheckPrevious(sliceFound, previousCategory) {
					sliceFound = append(sliceFound, value.Category)
					matchFound = true
					previousCategory = value.Category
				}
				continue
			} else if strings.Contains(strings.ToLower(value.SubCategory), input) {
				if jsonHandler.CheckPrevious(sliceFound, previousSubCategory) {
					sliceFound = append(sliceFound, value.SubCategory)
					matchFound = true
					previousSubCategory = value.SubCategory
				}
				continue
			}
		}

		// if no matches is found tell the user
		if !matchFound {
			fmt.Printf("%s was not found in the determined planes/categories\n", input)
			errorOccurred = true
		}
	}

	// check if error occurred, if so print a line break for the user to make things easier to read
	if errorOccurred {
		fmt.Println("-------------------------------")
	}

	// update the config with the new slice if the new slice list is not empty
	if len(sliceFound) != 0 {
		cfg.Config.Planes = sliceFound
		fmt.Println("Planes are now set to: ")

		// loop through the created slice and print it in an unorded list
		for _, p := range sliceFound {
			fmt.Printf("* %s\n", p)
		}
		fmt.Println("")
	} else {
		fmt.Println("No matches were found in the determined planes/categories")
		fmt.Println("Planes list was not altered")
		fmt.Println("")
	}

	return nil
}