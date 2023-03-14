package executor

import (
	"bytes"
	"os/exec"

	"github.com/jeffreylean/gopherize/internal/exercise"
)

type Result struct {
	Output string
	Err    string
	Exs    exercise.Exercise
}

func Execute(e exercise.Exercise) (Result, error) {
	var stdout, stderr bytes.Buffer

	exec := exec.Command("go", "run", "./"+e.File)
	exec.Stderr = &stderr
	exec.Stdout = &stdout

	err := exec.Run()
	return Result{Exs: e, Output: stdout.String(), Err: stderr.String()}, err
}
