package ch03

// User is struct
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// NewUser is constructor of User
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}
