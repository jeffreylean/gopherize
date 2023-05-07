package struct2

// Make it compile!
// To create an instance of a struct, you can use the struct literal syntax, which involves providing values for the fields in the order they are defined.
// You can also use named field initializers for clarity and flexibility.

// I AM NOT DONE

type User struct {
	name string
	age  int
}

func main() {
    // Using named field initializers
    person2 := User{name:,age:}
    // Using struct literal syntax
	person1 := User{"Jeffrey"}
}
