package repository

import (
	"context"
	"order-management/internal/model"
)

func (repo *repo) CreateUser(
	/*req*/ ctx context.Context, request model.User) (
	/*res*/ err error,
) {
	query := "INSERT INTO users (email, password, name) VALUES ($1, $2, $3)"

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(request.Email, request.Password, request.Name)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) GetUserByEmail(
	/*req*/ ctx context.Context, email string) (
	/*res*/ user model.User, err error,
) {
	query := "SELECT email, password FROM users WHERE email = $1"

	statemnt, err := repo.db.Prepare(query)
	if err != nil {
		return user, err
	}

	row, err := statemnt.Query(email)
	if err != nil {
		return user, err
	}

	for row.Next() {
		var email, password string

		err = row.Scan(&email, &password)
		if err != nil {
			return user, err
		}

		user = model.User{
			Email:    email,
			Password: password,
		}
	}

	return user, nil
}
