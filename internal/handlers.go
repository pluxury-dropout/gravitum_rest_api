package internal

import (
	"encoding/json"
	"errors"
	"gravitum_rest_api/users_db"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgconn"
)

type UserInfo struct {
	UsersModel users_db.UsersModelInterface
}

type CreateUserForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserForm struct {
	ID       int    `json:"id"`
	NewName  string `json:"name"`
	NewEmail string `json:"email"`
}

type GetUserForm struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *UserInfo) CreateUser(w http.ResponseWriter, r *http.Request) {
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
	if err := u.UsersModel.CreateUser(UserForm.Name, UserForm.Email, UserForm.Password); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
	})
}

func (u *UserInfo) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var updateForm UpdateUserForm

	if err := json.NewDecoder(r.Body).Decode(&updateForm); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}
	if updateForm.ID == 0 {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	if updateForm.NewName == "" && updateForm.NewEmail == "" {
		http.Error(w, "Nothing to update", http.StatusBadRequest)
		return
	}
	if err := u.UsersModel.UpdateUser(updateForm.ID, updateForm.NewName, updateForm.NewEmail); err != nil {
		if err.Error() == "User not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			http.Error(w, "Email alredy in use", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	updatedUser, err := u.UsersModel.GetUser(updateForm.ID)
	if err != nil {
		http.Error(w, "Failed to fetch updated user data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(updatedUser)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"message": "User updated successfully",
	// 	"user": map[string]interface{}{
	// 		"id":         updatedUser.ID,
	// 		"name":       updatedUser.Name,
	// 		"email":      updatedUser.Email,
	// 		"updated_at": updatedUser.UpdatedAt,
	// 	},
	// })
}

func (u *UserInfo) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}
	user, err := u.UsersModel.GetUser(id)
	if err != nil {
		if err.Error() == "User not found" {
			http.Error(w, "User not found", http.StatusNotFound)

		} else {
			http.Error(w, "Faield to get user", http.StatusInternalServerError)

		}
		return
	}
	userForm := GetUserForm{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userForm); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}

}
