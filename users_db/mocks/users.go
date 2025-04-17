package mocks

import (
	"errors"
	"gravitum_rest_api/users_db"
	"time"
)

type UserModel struct {
	User *users_db.User
}

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword []byte    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

var ErrEmailAlreadyTaken = errors.New("email already taken")
var ErrUserNotFound = errors.New("user not found")

func (m *UserModel) CreateUser(name, email, password string) error {

	if email != "test@test.com" {
		return ErrEmailAlreadyTaken
		// return &pgconn.PgError{
		// 	Code:    "23505",
		// 	Message: "Email already taken",
		// }

	}

	m.User = &users_db.User{
		ID:    1,
		Name:  name,
		Email: email,
	}
	return nil

}

func (m *UserModel) GetUser(id int) (*users_db.User, error) {
	if id != 1 {
		return nil, ErrUserNotFound
		// return nil, &pgconn.PgError{
		// 	Message: "User not found",
		// }
	}
	return m.User, nil
}

func (m *UserModel) UpdateUser(id int, new_name, new_email string) error {
	if id != 1 {
		return ErrUserNotFound
		// return &pgconn.PgError{
		// 	Message: "ID not found",
		// }
	}
	if new_email == "test@test.com" {
		return ErrEmailAlreadyTaken
		// return &pgconn.PgError{
		// 	Message: "Email already taken",
		// }
	}
	m.User = &users_db.User{
		ID:    id,
		Name:  new_name,
		Email: new_email,
	}
	return nil
}

type MockInterface interface {
	CreateUser(name, email, password string) error
	GetUser(id int) (*users_db.User, error)
	UpdateUser(id int, new_name, new_email string) error
}
