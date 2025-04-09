package usersdb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword []byte    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Post(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users (name, email, hashed_password, created_at) VALUES ($1,$2,$3, CURRENT_TIMESTAMP)`
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		log.Printf("Failed to execute the statement: %v", err)
		return err
	}
	return nil

}

func (m *UserModel) Get(id int) (*User, error) {
	stmt := `SELECT id, name, email, created_at, updated_at FROM users where id=$1`
	row := m.DB.QueryRow(stmt, id)

	user_info := &User{}

	err := row.Scan(&user_info.ID, &user_info.Name, &user_info.Email, &user_info.CreatedAt, &user_info.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("User not found")
		} else {
			return nil, err
		}
	}
	return user_info, nil
}

func (m *UserModel) PUT(id int, new_name, new_email string) error {
	stmt := `UPDATE users SET name=$1, email=$2, updated_at=CURRENT_TIMESTAMP WHERE id=$3`
	_, err := m.DB.Exec(stmt, new_name, new_email, id)
	if err != nil {
		log.Printf("Failed to update user: %v", id)
		return err
	}
	return nil
}
