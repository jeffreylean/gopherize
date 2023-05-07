package struct3

// Make it compile!
// The fields of an addressable struct are also addressable. The fields of an unaddressable struct are also unaddressable.
// The fields of unaddressable structs can't be modified. All composite literals, including struct composite literals are unaddressable.

// I AM NOT DONE

type User struct {
	name string
	age  int
}

func main() {
	// User{} is unaddressable
	// Assign User{} to a variable first to make it addressable, then only assign value through the struct variable. v.age = 26.
	User{}.age = 25

	// Addresses can be taken from composite literals.
	user := &User{"Jeffrey", 25}
	user.age = 26

}
