package main

import (
	"fmt"
	"os"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"
)

func commandExit(cfg *Config, args ...string) error {
	err := config.Update(cfg.Config)
	if err != nil {
		fmt.Printf("Exit code 1: %e\n", err)
		os.Exit(1)
		return err
	}

	fmt.Println("Closing the Generator... Goodbye!")
	os.Exit(0)

	return nil
}