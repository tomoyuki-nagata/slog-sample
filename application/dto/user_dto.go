package dto

import "todo-app/domain/model"

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

func GenerateUsers(users []model.User) []User {
	result := []User{}
	for _, user := range users {
		result = append(result, GenerateUser(user))
	}
	return result
}

func GenerateUser(user model.User) User {
	return User{
		Id:       user.Id(),
		Name:     user.Name(),
		Email:    user.Email(),
		Password: user.Password(),
	}
}
