package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://postgres:admin@localhost:5432/gravitum_users", "postgresql source name")

	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatalf("Failed to open DB, %v", err)
	}
	defer db.Close()

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:      *addr,
		TLSConfig: tlsConfig,
	}

	log.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
