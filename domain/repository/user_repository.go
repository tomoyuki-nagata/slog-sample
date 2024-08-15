package repository

import "todo-app/domain/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindById(id string) (model.User, error)
	FindByEmail(email string) (model.User, error)
	CreateUser(model.User) error
	UpdateUser(model.User) error
	DeleteUser(id string) error
}
