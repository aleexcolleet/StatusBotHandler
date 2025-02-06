package repositories

import "context"

type User struct {
	Name     string
	Id       int
	Password int
	Email    string
}

/*
	UserRepo is an interface that acts as a port:
		Defines how the core domain interacts with external systems.
		It sayis which operations are available without implementing them.

		This way, core domain remains agnostic from the different implementation
		of other systems.
*/

type UserRepo interface {
	CreateUser(ctx context.Context, user User) error
	ReadUser(ctx context.Context, userId int) (User, error)
}
