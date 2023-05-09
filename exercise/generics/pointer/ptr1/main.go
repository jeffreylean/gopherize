package main

// Make it compile!
// We should pass the pointer of string into the function.
func assignTo[T string](v T) {
	*v = "Testing"
}

func main() {
	var str = "Hello"
	assignTo(&str)
}
