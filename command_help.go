package main

import (
	"fmt"
)

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the D&D Encounter Generator!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
		fmt.Println("")
	}

	return nil
}