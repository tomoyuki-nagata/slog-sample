package datasource

import (
	"database/sql"
	"todo-app/domain/model"
	"todo-app/domain/repository"
)

type userRecordDatasource struct {
	*sql.DB
}

func NewUserRecordDatasource(db *sql.DB) repository.UserRepository {
	return &userRecordDatasource{db}
}

func (u *userRecordDatasource) FindAll() ([]model.User, error) {
	rows, err := u.Query(`SELECT * FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user, err := toUsers(rows)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRecordDatasource) FindById(id string) (model.User, error) {
	row := u.QueryRow(`SELECT * FROM users WHERE id = $1`, id)

	user, err := toUser(row)
	if err != nil {
		// TODO レコード0件の場合のエラーハンドリング
		if err == sql.ErrNoRows {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, nil
}

func (u *userRecordDatasource) FindByEmail(email string) (model.User, error) {
	row := u.QueryRow(`SELECT * FROM users WHERE email = $1`, email)

	user, err := toUser(row)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRecordDatasource) CreateUser(user model.User) error {
	_, err := u.Exec("INSERT INTO users (id, name, email,password) VALUES ($1, $2, $3, $4)", user.Id(), user.Name(), user.Email(), user.Password())
	if err != nil {
		return err
	}
	return nil
}

func (u *userRecordDatasource) UpdateUser(user model.User) error {
	_, err := u.Exec("UPDATE users set name = $1, email = $2, password = $3 WHERE id = $4 ", user.Name(), user.Email(), user.Password(), user.Id())
	if err != nil {
		return err
	}
	return nil
}

func (u *userRecordDatasource) DeleteUser(id string) error {
	_, err := u.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func toUsers(rows *sql.Rows) ([]model.User, error) {
	var users []model.User
	for rows.Next() {
		var id string
		var name string
		var email string
		var password string

		if err := rows.Scan(&id, &name, &email, &password); err != nil {
			return nil, err
		}
		users = append(users, model.BindUser(id, name, email, password))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func toUser(row *sql.Row) (model.User, error) {
	var id string
	var name string
	var email string
	var password string

	err := row.Scan(&id, &name, &email, &password)
	if err != nil {
		return model.User{}, err
	}

	return model.BindUser(id, name, email, password), nil
}
