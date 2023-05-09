package main

import "log"

type Map[K comparable] map[K]string

type Map2[K comparable, V any] map[K]V

func main() {
	// Make it compile!
	// Fix the type of the custom map
	m := make(Map[string]string)
	log.Println(m)

	// Fix the type of the custom map
	m2 := make(Map2[string]bool)
	log.Println(m2)
}
