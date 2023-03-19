package exercise

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

const (
	BugFix = "bugfix"
	QnA    = "qna"
)

// Enum for exercise's state
const (
	Done State = iota
	Pending
)

// Regex pattern for //I AM NOT DONE
const Regex = `\/\/\s*I\s+AM\s+NOT\s+DONE\b`

type State int
type Exercise struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	File string `yaml:"file"`
}
type Result struct {
	Output string
	Err    string
}

// Get state of the exercise by checking if "I AM NOT DONE" is being removed from the exercise or not.
func (e *Exercise) State() State {
	content, err := ioutil.ReadFile(e.File)
	if err != nil {
		log.Fatalf("failed to get state of exercise: %s", err.Error())
	}
	re := regexp.MustCompile(Regex)
	// Increment the counter if user done with it.
	if re.Match(content) {
		return State(Pending)
	}
	return State(Done)
}

func (e *Exercise) Run() Result {
	var stdout, stderr bytes.Buffer

	exec := exec.Command("go", "run", "./"+e.File)
	exec.Stderr = &stderr
	exec.Stdout = &stdout

	exec.Run()
	return Result{Output: stdout.String(), Err: stderr.String()}
}

// Get list of exercises
func GetExercises() (*[]Exercise, error) {
	var yFile []byte
	var exercises []Exercise

	yFile, err := os.ReadFile("exercise.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read exercise.yaml file: %w", err)
	}

	if err = yaml.Unmarshal(yFile, &exercises); err != nil {

		return nil, err
	}

	return &exercises, nil
}

// Search for particular exercise
func Search(name string) (*Exercise, error) {
	exs, err := GetExercises()
	if err != nil {
		return nil, fmt.Errorf("failed to get all exercises: %w", err)
	}
	for _, each := range *exs {
		if strings.EqualFold(each.Name, name) {
			return &each, nil
		}
	}
	return nil, fmt.Errorf("no exercise found")
}
