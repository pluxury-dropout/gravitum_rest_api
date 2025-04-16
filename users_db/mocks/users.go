package mocks

import (
	"gravitum_rest_api/users_db"

	"github.com/jackc/pgconn"
)

type UserModel struct {
}

func (m *UserModel) CreateUser(name, email, password string) error {
	switch email {
	case "test@test.com":
		return &pgconn.PgError{
			Code:    "23505",
			Message: "Email already taken",
		}
	default:
		return nil
	}

}

func (m *UserModel) GetUser(id int) (*users_db.User, error) {
	if id != 1 {
		return nil, &pgconn.PgError{
			Message: "User not found",
		}
	}
	return &users_db.User{
		ID:    1,
		Name:  "test",
		Email: "test@test.com",
	}, nil
}

func (m *UserModel) UpdateUser(id int, new_name, new_email string) error {
	if id != 1 {
		return &pgconn.PgError{
			Message: "ID not found",
		}
	}
	if new_email == "test@test.com" {
		return &pgconn.PgError{
			Message: "Email already taken",
		}
	}
	return nil
}

type MockInterface interface {
	CreateUser(name, email, password string) error
	GetUser(id int) error
	UpdateUser(id int, new_name, new_email string) error
}
