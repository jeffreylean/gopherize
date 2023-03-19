package cmd

import (
	"github.com/jeffreylean/gopherize/internal/exercise"
	"github.com/jeffreylean/gopherize/internal/verify"
	"github.com/spf13/cobra"
)

// Verify all the exercises
func verifyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "verify",
		Short: "Verify all exercise sequentially",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Retrieve all exercise.
			exercises, err := exercise.GetExercises()
			if err != nil {
				return err
			}
			// Verify all the exercise.
			verify.Verify(*exercises, &verify.Progress{Done: 0, Total: uint64(len(*exercises))})
			return nil
		},
	}

	return cmd
}
