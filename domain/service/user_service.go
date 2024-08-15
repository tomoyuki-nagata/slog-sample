package service

import (
	"todo-app/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (u UserService) ExistEmail(email string) (bool, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	return user.Id() != "", nil
}
