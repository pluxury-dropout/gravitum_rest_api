package users_db_test

import (
	"database/sql"
	"gravitum_rest_api/internal/assert"
	"gravitum_rest_api/users_db"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "postgres://postgres:admin@localhost:5432/gravitum_users")
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}
	_, _ = db.Exec("TRUNCATE users RESTART IDENTITY")
	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	model := &users_db.UserModel{DB: db}

	err := model.CreateUser("Test", "test@test.com", "pass123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGetUser(t *testing.T) {
	db := setupTestDB(t)
	model := &users_db.UserModel{DB: db}
	_ = model.CreateUser("Test", "test@test.com", "pass123")

	user, err := model.GetUser(1)
	if err != nil {
		t.Errorf("Expected user, got errror:%v", err)
	}
	if user.Email != "test@test.com" {
		assert.Equal(t, user.Email, "test@test.com")
		// t.Errorf("Expected email test@test.com, got %v", user.Email)
	}

}

func TestUpdateUser(t *testing.T) {
	db := setupTestDB(t)
	model := &users_db.UserModel{DB: db}

	_ = model.CreateUser("test_create", "test@test.com", "pass123")
	err := model.UpdateUser(1, "updated_name", "updated_test@test.com")
	if err != nil {
		t.Errorf("update failed: %v", err)
	}
	user, _ := model.GetUser(1)
	if user.Name != "updated_name" {
		assert.Equal(t, user.Name, "updated_name")
		// t.Errorf("expected updated_name, got %s", user.Name)
	}
}
