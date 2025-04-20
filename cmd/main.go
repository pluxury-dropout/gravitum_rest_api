package main

import (
	"database/sql"
	"gravitum_rest_api/users_db"
	"log"
	"net/http"
	"os"

	"gravitum_rest_api/internal"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env not found: %v", err)
	}
	addr := os.Getenv("ADDR")
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("No DSN env")
	}

	db, err := openDB(dsn)
	if err != nil {
		log.Fatalf("Failed to open DB, %v", err)
	}
	defer db.Close()

	userModel := &users_db.UserModel{DB: db}
	userHandler := &internal.UserInfo{UsersModel: userModel}

	srv := &http.Server{
		Addr:    addr,
		Handler: internal.SetupRoutes(userHandler),
	}

	log.Printf("Starting server on %s", addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
