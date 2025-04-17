package users_db_test

import (
	"gravitum_rest_api/internal/assert"
	"gravitum_rest_api/users_db/mocks"
	"testing"
)

func TestModelCreateUser(t *testing.T) {
	model := &mocks.UserModel{}

	err := model.CreateUser("Test", "test1@test.com", "pass123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
func TestModelGetUser(t *testing.T) {
	model := &mocks.UserModel{}
	_ = model.CreateUser("Test", "test@test.com", "pass123")

	user, err := model.GetUser(1)
	if err != nil {
		t.Errorf("Expected user, got errror:%v", err)
	}
	if user.Email != "test@test.com" {
		assert.Equal(t, user.Email, "test@test.com")
	}

}
func TestModelUpdateUser(t *testing.T) {
	model := &mocks.UserModel{}

	_ = model.CreateUser("test_create", "test@test.com", "pass123")
	err := model.UpdateUser(1, "updated_name", "updated_test@test.com")
	if err != nil {
		t.Errorf("update failed: %v", err)
	}
	user, _ := model.GetUser(1)
	if user.Name != "updated_name" {
		assert.Equal(t, user.Name, "updated_name")
	}
}
