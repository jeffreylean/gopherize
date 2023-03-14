package main

import (
	"fmt"
	"os"

	"github.com/jeffreylean/gopherize/cmd"
)

func main() {
	rootCmd := cmd.Root()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
