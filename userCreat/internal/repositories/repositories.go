package repositories

import "context"

tyoe User struct {
	Name string
	Id   string
	Password int
	Email string
}
//UserRepo define el contrato que cualquier repo (base de datos, memoria, erc...(
//debe cumplir. Con ello, podemos cambiar de base de datos sin afectar
//a la logica de negocio

type UserRepo interface {
	CreteUser(ctx context.Context, user User) error
	ReadUser(ctx context.Context, userId int) (User, error)
}
