package packages4

// Make the test pass!
// init is a reserved work in Go for declaring an init() function. There can be multiple functions named as `init` declared in a package. `init` should not have any input parameters and return any results.
// At runtime, `init` function will be sequentially invoked once and only once before the `main` function. It's like a static initializer.

// I AM NOT DONE

var s string
var gopher string

// Complete the init() function so the variable s holds value of "Hi,Gopher"
func init() {
}

// Don't write any code in the getValue() function.
func getValue() string {
	return s
}
