package main

import (
	"fmt"
	"os"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"
	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
)

func commandExit(cfg *Config, args ...string) error {
	// updates the config file and handles errors
	err := config.Update(cfg.Config)
	if err != nil {
		errorhandling.LogError(err, "commandExit")
		os.Exit(1)
	}

	// print a goodbye message and close the software
	fmt.Println("Closing the Generator... Goodbye!")
	os.Exit(0)

	return nil
}