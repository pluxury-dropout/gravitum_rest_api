package main

import (
	"encoding/json"
	"net/http"
)

type CreateUserForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *userInfo) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var UserForm CreateUserForm
	if err := json.NewDecoder(r.Body).Decode(&UserForm); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if UserForm.Name == "" || UserForm.Email == "" || UserForm.Password == "" {
		http.Error(w, "Name, email and password cannot be blank", http.StatusBadRequest)
		return
	}
	if err := u.userModel.CreateUser(UserForm.Name, UserForm.Email, UserForm.Password); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
	})
}
