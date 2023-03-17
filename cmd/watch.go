package cmd

import (
	"github.com/jeffreylean/gopherize/internal/exercise"
	"github.com/jeffreylean/gopherize/internal/watcher"
	"github.com/spf13/cobra"
)

// Auto re-run the file when user edited the exercise.
func watchCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "watch",
		Short: "Watch hot reload your exercise",
		RunE: func(cmd *cobra.Command, args []string) error {
			exercises, err := exercise.GetExercises()
			if err != nil {
				return err
			}

			if err := watcher.Loop((*exercises)); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
