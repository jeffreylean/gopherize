package struct1

// Make it compile!
// A struct is defined using the type keyword, followed by the struct's name,
// the struct keyword, and a set of curly braces that enclose the fields of
// the struct. Each field has a name and a data type.

// I AM NOT DONE

import "fmt"

type User struct {
	name
	age int
}

func main() {
	var user User

	fmt.Printf("User name is %s", user.name)
}
