package main

import (
	"fmt"
)

// Make it compile!
// In Go, we can do type switch to discover the type of any/interface{}. any/interface{} in golang represent "any" value.

// I AM NOT DONE
func main() {
	types := []any{true, 1, "I am a string", 2.2}

	for _, value := range types {
		switch t := value.(type) {
		case bool:
			fmt.Printf("%v is a bool \n", t)
        case
		default:
			fmt.Println("Unrecognize type")
		}
	}
}
