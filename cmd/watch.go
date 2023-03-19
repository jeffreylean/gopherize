package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/jeffreylean/gopherize/internal/exercise"
	"github.com/jeffreylean/gopherize/internal/prompt"
	"github.com/jeffreylean/gopherize/internal/verify"
	"github.com/spf13/cobra"
)

// Auto re-run the file when user edited the exercise.
func watchCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "watch",
		Short: "Hot re-verify your exercises",
		RunE: func(cmd *cobra.Command, args []string) error {
			exercises, err := exercise.GetExercises()
			if err != nil {
				return err
			}

			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				return err
			}
			defer watcher.Close()

			p, err := os.Getwd()
			if err != nil {
				return err
			}

			// Recursively watch the whole exercise directory.
			err = filepath.Walk(p+"/exercise", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					err = watcher.Add(path)
					if err != nil {
						return err
					}
				}
				return nil
			})
			if err != nil {
				return err
			}

			// Boolean channel to check if the exercise is done completely, if so will stop looping.
			done := make(chan bool)
			progress := &verify.Progress{Done: 0, Total: uint64(len(*exercises))}
			// Verify initial state
			verify.Verify(*exercises, progress)

			go func(exs []exercise.Exercise, p *verify.Progress) {
				defer close(done)
				exercisesToVerify := exs[p.Done:]
				for len(exercisesToVerify) != 0 {
					select {
					case event, ok := <-watcher.Events:
						if !ok {
							return
						}
						if event.Has(fsnotify.Create) || event.Has(fsnotify.Chmod) || event.Has(fsnotify.Write) {
							prompt.ClearScreen()
							// Only verify exercises which are not yet completed.

							if verify.Verify(exercisesToVerify, p) {
								return
							}
							exercisesToVerify = exs[p.Done:]
						}
					case err, ok := <-watcher.Errors:
						if !ok {
							return
						}
						log.Fatal(err)
					}
				}
			}(*exercises, progress)

			<-done
			return nil
		},
	}

	return cmd
}
