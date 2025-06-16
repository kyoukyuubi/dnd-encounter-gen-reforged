package main

import "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"

func main() {
    cfg := &Config{}

	// test the init function
	config.Init()

	startRepl(cfg)
}