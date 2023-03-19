package verify

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/jeffreylean/gopherize/internal/exercise"
	"github.com/jeffreylean/gopherize/internal/prompt"
	progressbar "github.com/schollz/progressbar/v3"
)

func Verify(exercises []exercise.Exercise) error {
	// Instantiate progress bar
	bar := progressbar.NewOptions(len(exercises), progressbar.OptionEnableColorCodes(true), progressbar.OptionSetDescription("[blue][reset] progress:"), progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[blue]=[reset]",
		SaucerHead:    "[blue]>[reset]",
		SaucerPadding: "-",
		BarStart:      "[",
		BarEnd:        "]",
	}), progressbar.OptionSetWidth(60), progressbar.OptionSetPredictTime(false), progressbar.OptionSetElapsedTime(false))

	if err := bar.RenderBlank(); err != nil {
		return err
	}

	for _, exs := range exercises {
		bar.Describe(fmt.Sprintf("Compiling %s...", exs.Name))

		// Execute the exercise, check if it compile correctly.
		res := exs.Run()
		if res.Err != "" {
			prompt.Warn(fmt.Sprintf("Compiling of %s failed! Please try again. Here's the output:", exs.Name))
			fmt.Println(res.Err)
			os.Exit(1)
		}
		bar.Add(1)

		// Prompt the completion message. Won't continue verify if the exercise is in pending state.
		if !Completion(exs, res.Output) {
			return nil
		}
	}

	// Done verifying all exercises, this means the whole exercise is completed. Prompt the success message.
	fmt.Println("")
	color.Green.Println("Congratulations!!!! You had completed all the exercise, now go out there and shine as a gopher!!")
	return nil
}

func Completion(exs exercise.Exercise, output string) bool {
	// If it is done, return.
	if exs.State() == exercise.Done {
		return true
	}

	// Prompt success message.
	prompt.Success(fmt.Sprintf("Successfully compiled %s", exs.Name))
	fmt.Println("")
	// If there's stdoutput from the exercise, print it.
	if output != "" {
		fmt.Println("Output:")
		fmt.Println(output)
	}
	fmt.Println("")
	fmt.Println("You can keep working on the exercise,")
	fmt.Printf("or jump into the next one by removing %s comment.", color.Bold.Sprint("I AM NOT DONE"))

	return false
}

//func Loop(exercises []exercise.Exercise) error {
//	// Initiate watcher
//	watcher, err := fsnotify.NewWatcher()
//	if err != nil {
//		return err
//	}
//	defer watcher.Close()
//
//	p, err := os.Getwd()
//	if err != nil {
//		return err
//	}
//
//	// Recursively watch the whole exercise directory.
//	err = filepath.Walk(p+"/exercise", func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//		if info.IsDir() {
//			err = watcher.Add(path)
//			if err != nil {
//				return err
//			}
//		}
//		return nil
//	})
//	if err != nil {
//		return err
//	}
//
//	// Instantiate progress bar
//	bar := progressbar.NewOptions(100, progressbar.OptionEnableColorCodes(true), progressbar.OptionSetDescription("[blue][reset] progress:"), progressbar.OptionSetTheme(progressbar.Theme{
//		Saucer:        "[blue]=[reset]",
//		SaucerHead:    "[blue]>[reset]",
//		SaucerPadding: "-",
//		BarStart:      "[",
//		BarEnd:        "]",
//	}), progressbar.OptionSetWidth(60), progressbar.OptionSetPredictTime(false), progressbar.OptionSetElapsedTime(false))
//	bar.RenderBlank()
//
//	// Boolean channel to check if the exercise is done completely, if so will stop looping.
//	done := make(chan bool)
//	// Exercise counter
//	i := 0
//
//	go func(count int, exercises []exercise.Exercise, bar *progressbar.ProgressBar) {
//		defer close(done)
//		for i < len(exercises) {
//			exs := exercises[i]
//			prompt(bar, exs)
//
//			select {
//			case event, ok := <-watcher.Events:
//				if !ok {
//					return
//				}
//				if event.Has(fsnotify.Create) || event.Has(fsnotify.Chmod) || event.Has(fsnotify.Write) {
//					// Execute the exercise, check if it compile correctly.
//					res, err := executor.Execute(exs)
//					if err != nil {
//						log.Println(res.Err)
//						continue
//					}
//					// Print out the output.
//					fmt.Println(res.Output)
//
//					// Check if gopher is done or not with the exercise
//					content, err := ioutil.ReadFile(exs.File)
//					if err != nil {
//						log.Println(err)
//						return
//					}
//					re := regexp.MustCompile(Regex)
//					// Increment the counter if user done with it.
//					if re.Match(content) {
//						continue
//					}
//					i++
//					bar.Add(1)
//				}
//			case err, ok := <-watcher.Errors:
//				if !ok {
//					return
//				}
//				log.Println("error:", err)
//			}
//		}
//	}(i, exercises, bar)
//
//	<-done
//	bar.Close()
//	return nil
//}
//
////func prompt(bar *progressbar.ProgressBar, exs exercise.Exercise) {
////	// Clears the terminal with an ANSI escape code.
////	// Works in UNIX and newer Windows terminals.
////	fmt.Println("\x1Bc")
////	fmt.Println(bar.String())
////	fmt.Println("Compiling ", exs.File)
////}
