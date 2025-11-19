package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/vashkevichjr/go_crud/internal/config"
	"github.com/vashkevichjr/go_crud/internal/repository"
	"github.com/vashkevichjr/go_crud/internal/transport/rest"
)

func main() {
	r := chi.NewRouter()

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := config.Load()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("connected to database")

	defer func(conn *pgx.Conn) {
		if err := conn.Close(ctx); err != nil {
			log.Printf("failed to close connection: %v", err)
		}
	}(conn)

	repo := repository.NewPostgresRepo(conn)
	handler := rest.NewHandler(repo)

	r.Post("/", handler.SaveAndGetNumber)

	log.Println("starting server on port 8080")
	err = http.ListenAndServe(cfg.Port, r)
}
