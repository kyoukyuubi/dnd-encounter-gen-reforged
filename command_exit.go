package main

import (
	"fmt"
	"os"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"
)

func commandExit(cfg *Config, args ...string) error {
	// updates the config file and handles errors
	err := config.Update(cfg.Config)
	if err != nil {
		fmt.Printf("Exit code 1: %e\n", err)
		os.Exit(1)
		return err
	}

	// print a goodbye message and close the software
	fmt.Println("Closing the Generator... Goodbye!")
	os.Exit(0)

	return nil
}