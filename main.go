package main

import (
	"fmt"
	"os"

	"github.com/jeffreylean/gopherize/cmd"
)

func main() {
	// Check if user is in the gopherize directory
	if _, err := os.Stat("exercise.yaml"); err != nil {
		fmt.Print("gopherize must be run from the gopherize directory!\nTry `cd gopherize/`.")
		os.Exit(1)
	}

	rootCmd := cmd.Root()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
