package exercise

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetExercises(t *testing.T) {
	p, _ := os.Getwd()
	os.Chdir("../..")
	t.Cleanup(func() {
		os.Chdir(p)
	})

	execs, err := GetExercises()
	require.NoError(t, err)
	require.NotEmpty(t, execs)
}

func TestSearch(t *testing.T) {
	p, _ := os.Getwd()
	os.Chdir("../..")
	t.Cleanup(func() {
		os.Chdir(p)
	})

	t.Run("Search with lowercase should works", func(t *testing.T) {
		exec, err := Search("helloworld")
		require.NoError(t, err)
		require.NotNil(t, exec)
	})

	t.Run("Search with uppercase should works", func(t *testing.T) {
		exec, err := Search("helLOworLd")
		require.NoError(t, err)
		require.NotNil(t, exec)
	})

	t.Run("Search with unknown exercise", func(t *testing.T) {
		exec, err := Search("unknown")
		require.Error(t, err)
		require.Nil(t, exec)
	})
}
