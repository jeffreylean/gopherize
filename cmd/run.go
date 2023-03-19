package cmd

import (
	"fmt"
	"os"

	"github.com/jeffreylean/gopherize/internal/exercise"
	"github.com/jeffreylean/gopherize/internal/prompt"
	"github.com/jeffreylean/gopherize/internal/verify"
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
			res := exs.Run()
			if res.Err != "" {
				prompt.Warn(fmt.Sprintf("Compiling of %s failed! Please try again. Here's the output:", exs.Name))
				fmt.Println(res.Err)
				os.Exit(1)
			}

			// Prompt completion message and verify the state of the exercise.
			if verify.Completion(*exs, res.Output) {
				fmt.Println(res.Output)
				prompt.Success("Successfully ran " + exs.Name)
			}
			return nil
		},
	}

	return cmd
}
