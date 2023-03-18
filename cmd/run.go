package cmd

import (
	"fmt"

	"github.com/jeffreylean/gopherize/internal/executor"
	"github.com/jeffreylean/gopherize/internal/exercise"
	"github.com/spf13/cobra"
)

// Run specific exercise
func runCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "run [exercise]",
		Short: "Run exercise",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Search for the exercise
			exs, err := exercise.Search(args[0])
			if err != nil {
				return err
			}
			res, err := executor.Execute(*exs)
			if err != nil {
				return err
			}

			fmt.Println(res.Output)
			return nil
		},
	}

	return cmd
}
