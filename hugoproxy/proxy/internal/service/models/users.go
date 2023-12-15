package models

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Conditions struct {
	limit  string
	offset string
}

func NewUser(name string, age int) *User {
	return &User{
		Name: "",
		Age:  0,
	}
}
