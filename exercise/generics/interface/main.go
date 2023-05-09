package main

import "encoding/json"

type user struct{}

// Make it compile!
// You should implements the necessary interface
// for the `user` struct.

func jsonMarshal[T interface {
	json.Marshaler
}](v T) ([]byte, error) {
	return v.MarshalJSON()
}

func main() {
	jsonMarshal(user{})
}
