package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"ticket_wave/internal/handler"
	"ticket_wave/internal/repository"
	"ticket_wave/internal/service"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %q", err)
	}
}

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	handler := handler.NewHandler(service)

	handler.Run(":8080")
}
