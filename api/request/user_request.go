package request

import (
	"todo-app/application/dto"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type PostUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u PostUserRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Name,
			validation.Required.Error("名前は必須入力です"),
		),
		validation.Field(
			&u.Email,
			validation.Required.Error("メールアドレスは必須入力です"),
			is.Email.Error("メールアドレスを入力して下さい"),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error("パスワードは必須入力です"),
			validation.Length(8, 0).Error("パスワードは8文字以上です"),
		),
	)
}

func (r PostUserRequest) ToDto() dto.User {
	return dto.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

type PutUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u PutUserRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Email,
			is.Email.Error("メールアドレスを入力して下さい"), // 空の場合は有効
		),
		validation.Field(
			&u.Password,
			validation.Length(8, 0).Error("パスワードは8文字以上です"), // 空の場合は有効
		),
	)
}

func (r PutUserRequest) ToDto() dto.User {
	return dto.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
