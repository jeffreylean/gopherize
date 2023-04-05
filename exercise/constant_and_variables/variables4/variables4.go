package main

// Make it compile!
// Scopes of variables and named constants

// I AM NOT DONE

var a bool

func main() {
	println(a) // false

	var a = 1
	{
		a := a + 1
		y := a + 1
	}
	println(a) // 1
	println(y)
}
