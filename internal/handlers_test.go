package internal_test

import (
	"bytes"
	"encoding/json"
	"gravitum_rest_api/internal"
	"gravitum_rest_api/internal/assert"
	"gravitum_rest_api/users_db/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	mockModel := &mocks.UserModel{}
	handler := &internal.UserInfo{UsersModel: mockModel}

	userData := internal.CreateUserForm{
		Name:     "test_name",
		Email:    "test@test.com",
		Password: "pass123",
	}

	payload, err := json.Marshal(userData)
	if err != nil {
		t.Fatalf("failed to marshall user data: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.CreateUser(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)
}
