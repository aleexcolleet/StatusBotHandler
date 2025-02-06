package repositories

import "context"

/*
User is a struct
*/
type User struct {
	Name     string
	Id       int
	Password string
	Mail     string
}

/*
UserRepository interface acts as a port because it defines
the methods that any repository must implement

CreateUser -> pretends to create a user on whatever place
we keep data, but not return it.
ReadUser -> Returns a copy by value of a User struct.
*/
type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
	ReadUser(ctx context.Context, userId int) (User, error)
}
