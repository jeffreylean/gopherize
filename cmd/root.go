package cmd

import (
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "gopherize <command>",
		Short: "Let's gopherize you into gopher!",
	}

	rootCmd.AddCommand(runCmd())
	return rootCmd
}
