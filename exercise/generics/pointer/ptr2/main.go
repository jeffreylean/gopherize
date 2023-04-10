package main

type Human interface {
	HasAge() bool
	Age() int8
}

type user struct {
	age int8
}

func (u *user) HasAge() bool {
	return u.age > 0
}

func (u *user) Age() int8 {
	return u.age
}

func getAge[T Human](v T) int8 {
	if v.HasAge() {
		return v.Age()
	}
	return 0
}

func main() {
	u := user{}
	// Make it compile!
	// You need to pass the correct value that
	// implements the interface.
	getAge(u)
}
