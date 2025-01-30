package users

import (
	"cmd/main.go/internal/repositories"
	"context"
)

/*
	inmemory es una implementacion del repositorio en Memoria
	Implementamos por tanto la interfaz UserReop, almacenando usuarios
	en un slice de memoria (Users []repositories.User(
*/

type ImUserStore struct {
	Users []repositories.UserRepo
}

//NewImUsersStore es un constructor para el slice de users
func NewImUsersStore(ctx context.Context) *ImUserStore {
	return &ImUserStore{
		Users: []repositories.User{},
	}
}

//AddNewUser adds users inside the slice
func (S *ImUserStore) AddNewUser(ctx context.Context, user repositories.User) error {
	S.Users = append(S.Users, user)
	return nil
}
