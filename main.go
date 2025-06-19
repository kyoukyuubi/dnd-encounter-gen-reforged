package main

import (
	"os"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func main() {
	err := jsonHandler.Check()
	if err != nil {
		errorhandling.LogError(err, "Check()")
		os.Exit(1)
	}

	config.Init()

	config, err := config.Read()
	if err != nil {
		errorhandling.LogError(err, "Read()")
		os.Exit(1)
	}

	cfg := &Config{
		Config: config,
	}

	startRepl(cfg)
}