package verify

import (
	"fmt"
	"log"

	"github.com/gookit/color"
	"github.com/jeffreylean/gopherize/internal/exercise"
	"github.com/jeffreylean/gopherize/internal/prompt"
	progressbar "github.com/schollz/progressbar/v3"
)

type Progress struct {
	Done  uint64
	Total uint64
}

func Verify(exercises []exercise.Exercise, progress *Progress) bool {
	// Instantiate progress bar
	bar := progressbar.NewOptions(int(progress.Total), progressbar.OptionEnableColorCodes(true), progressbar.OptionSetDescription("[blue][reset] progress:"), progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[blue]=[reset]",
		SaucerHead:    "[blue]>[reset]",
		SaucerPadding: "-",
		BarStart:      "[",
		BarEnd:        "]",
	}), progressbar.OptionSetWidth(60), progressbar.OptionSetPredictTime(false), progressbar.OptionSetElapsedTime(false))

	bar.Set(int(progress.Done))

	if err := bar.RenderBlank(); err != nil {
		log.Fatal(err)
	}

	for _, exs := range exercises {
		bar.Describe(fmt.Sprintf("Compiling %s...", exs.Name))

		// Execute the exercise, check if it compile correctly.
		res := exs.Run()
		if res.Err != "" {
			fmt.Println("")
			prompt.Warn(fmt.Sprintf("Compiling of %s failed! Please try again. Here's the output:", exs.Name))
			fmt.Println(res.Err)
			return false
		}
		// Prompt the completion message. Won't continue verify if the exercise is in pending state.
		if !Completion(exs, res.Output) {
			return false
		}
		bar.Add(1)
		progress.Done++
	}

	// Done verifying all exercises, this means the whole exercise is completed. Prompt the success message.
	fmt.Println()
	color.Green.Println("Congratulations!!!! You had completed all the exercise, now go out there and shine as a gopher!!")
	bar.Close()
	return true
}

func Completion(exs exercise.Exercise, output string) bool {
	// If it is done, return.
	if exs.State() == exercise.Done {
		return true
	}

	// Prompt success message.
	fmt.Println()
	prompt.Success(fmt.Sprintf("Successfully compiled %s", exs.Name))
	fmt.Println()
	// If there's stdoutput from the exercise, print it.
	if output != "" {
		fmt.Println("Output:")
		fmt.Println(output)
	}
	fmt.Println()
	fmt.Println("You can keep working on the exercise,")
	fmt.Printf("or jump into the next one by removing %s comment.", color.Bold.Sprint("I AM NOT DONE"))

	return false
}
