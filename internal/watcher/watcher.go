package watcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/fsnotify/fsnotify"
	"github.com/jeffreylean/gopherize/internal/executor"
	"github.com/jeffreylean/gopherize/internal/exercise"
	progressbar "github.com/schollz/progressbar/v3"
)

// Regex pattern for //I AM NOT DONE
const Regex = `\/\/\s*I\s+AM\s+NOT\s+DONE\b`

// Loop spawn a goroutine to keep watching on list of exercises. It's like hot loading for exercises, whenever user write the file,
// fsnotify will capture the events and executor will execute the exercises to see if the exercise done correctly.
// If it is correct, and //I AM NOT DONE is being removed, will move on the the next exercises else it will stay at the current exercise.
func Loop(exercises []exercise.Exercise) error {
	// Initiate watcher
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

	// Instantiate progress bar
	bar := progressbar.NewOptions(100, progressbar.OptionEnableColorCodes(true), progressbar.OptionSetDescription("[blue][reset] progress:"), progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[blue]=[reset]",
		SaucerHead:    "[blue]>[reset]",
		SaucerPadding: "-",
		BarStart:      "[",
		BarEnd:        "]",
	}), progressbar.OptionSetWidth(60), progressbar.OptionSetPredictTime(false), progressbar.OptionSetElapsedTime(false))
	bar.RenderBlank()

	// Boolean channel to check if the exercise is done completely, if so will stop looping.
	done := make(chan bool)
	// Exercise counter
	i := 0

	go func(count int, exercises []exercise.Exercise, bar *progressbar.ProgressBar) {
		defer close(done)
		for i < len(exercises) {
			exs := exercises[i]
			prompt(bar, exs)

			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) || event.Has(fsnotify.Chmod) || event.Has(fsnotify.Write) {
					// Execute the exercise, check if it compile correctly.
					res, err := executor.Execute(exs)
					if err != nil {
						log.Println(res.Err)
						continue
					}
					// Print out the output.
					fmt.Println(res.Output)

					// Check if gopher is done or not with the exercise
					content, err := ioutil.ReadFile(exs.File)
					if err != nil {
						log.Println(err)
						return
					}
					re := regexp.MustCompile(Regex)
					// Increment the counter if user done with it.
					if re.Match(content) {
						continue
					}
					i++
					bar.Add(1)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}(i, exercises, bar)

	<-done
	return nil
}

func prompt(bar *progressbar.ProgressBar, exs exercise.Exercise) {
	// Clears the terminal with an ANSI escape code.
	// Works in UNIX and newer Windows terminals.
	fmt.Println("\x1Bc")
	fmt.Println(bar.String())
	fmt.Println("Compiling ", exs.File)
}
