package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("You have not caught any Pokemon yet, try the catch command")
	}
	for name := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
