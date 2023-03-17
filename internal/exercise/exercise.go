package exercise

import (
	"fmt"
	"log"
	"os"
	"sync"

	yaml "gopkg.in/yaml.v3"
)

const (
	BugFix = "bugfix"
	QnA    = "qna"
)

type Exercise struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	File string `yaml:"file"`
}

var once sync.Once
var exercises *[]Exercise
var err error

// Get list of exercises
func GetExercises() (*[]Exercise, error) {
	if exercises == nil {
		once.Do(func() {
			var yFile []byte
			yFile, err = os.ReadFile("exercise.yaml")
			if err != nil {
				return
			}

			err = yaml.Unmarshal(yFile, &exercises)
		})
	} else {
		log.Println("Exercise instance has been created")
	}

	if err != nil {
		return nil, err
	}
	return exercises, nil
}

// Search for particular exercise
func Search(name string) (*Exercise, error) {
	exs, err := GetExercises()
	if err != nil {
		return nil, err
	}
	for _, each := range *exs {
		if each.Name == name {
			return &each, nil
		}
	}
	return nil, fmt.Errorf("No exercise found")
}
