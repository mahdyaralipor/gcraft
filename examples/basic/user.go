package models

// User is a sample struct to demonstrate gcraft generation.
//
//go:generate gcraft generate -type User -src ./user.go
type User struct {
	ID       int
	Name     string
	Email    string
	Age      int
	Tags     []string
	IsActive bool
}

// UserRepository is a sample interface to demonstrate mock generation.
//
//go:generate gcraft generate -type UserRepository -src ./user.go
type UserRepository interface {
	FindByID(id int) (*User, error)
	FindAll() ([]*User, error)
	Save(u *User) error
	Delete(id int) error
}
