package internal

import (
	"net/http"
)

func SetupRoutes(userHandler *UserInfo) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users/", userHandler.GetUser)
	mux.HandleFunc("PUT /users/", userHandler.UpdateUser)

	return mux
}
