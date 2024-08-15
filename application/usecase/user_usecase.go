package usecase

import (
	"todo-app/application/dto"
	"todo-app/domain/model"
	"todo-app/domain/repository"
)

//go:generate mockgen -source ./user_usecase.go -destination ./mock/user_usecase_mock.generated.go -package usecase
type UserRecord interface {
	FindAll() ([]dto.User, error)
	FindUser(id string) (dto.User, error)
	CreateUser(user dto.User) (string, error)
	UpdateUser(user dto.User) error
	DeleteUser(id string) error
}

func NewUserRecordInteractor(repo repository.UserRepository) UserRecord {
	return &userInteractor{repo: repo}
}

type userInteractor struct {
	repo repository.UserRepository
}

func (u userInteractor) FindAll() ([]dto.User, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return dto.GenerateUsers(users), nil
}

func (u userInteractor) FindUser(id string) (dto.User, error) {
	user, err := u.repo.FindById(id)
	if err != nil {
		return dto.User{}, err
	}
	return dto.GenerateUser(user), nil
}

func (u userInteractor) CreateUser(userDto dto.User) (string, error) {
	user, err := model.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		return "", err
	}
	u.repo.CreateUser(user)
	return user.Id(), nil
}

func (u userInteractor) UpdateUser(userDto dto.User) error {
	user, err := u.repo.FindById(userDto.Id)
	if err != nil {
		return err
	}

	if userDto.Name != "" {
		user = user.ChangeName(userDto.Name)
	}
	if userDto.Email != "" {
		user = user.ChangeEmail(userDto.Email)
	}
	if userDto.Password != "" {
		user = user.ChangePassword(userDto.Password)
	}

	err = u.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u userInteractor) DeleteUser(id string) error {
	err := u.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
