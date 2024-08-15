package model

import (
	"errors"
	"todo-app/core/util"

	"github.com/google/uuid"
)

// Userを新規作成時に使用
func NewUser(name, email, password string) (User, error) {
	if name == "" {
		return User{}, errors.New("名前がありません")
	}
	if email == "" {
		return User{}, errors.New("メールアドレスがありません")
	}
	if password == "" {
		return User{}, errors.New("パスワードがありません")
	}

	hashedPassword, err := util.GenerateHashPassword(password)
	if err != nil {
		return User{}, err
	}
	return User{
		id:       uuid.New().String(),
		name:     name,
		email:    email,
		password: hashedPassword,
	}, nil
}

// DBから取得したデータをUserにバインドする
func BindUser(id, name, email, password string) User {
	return User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}
}

type User struct {
	id       string
	name     string
	email    string
	password string
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) ChangeName(name string) User {
	return User{
		id:       u.id,
		name:     name,
		email:    u.email,
		password: u.password,
	}
}

func (u *User) ChangeEmail(email string) User {
	return User{
		id:       u.id,
		name:     u.name,
		email:    email,
		password: u.password,
	}
}

func (u *User) ChangePassword(password string) User {
	hashedPassword, _ := util.GenerateHashPassword(password)
	return User{
		id:       u.id,
		name:     u.name,
		email:    u.email,
		password: hashedPassword,
	}
}
