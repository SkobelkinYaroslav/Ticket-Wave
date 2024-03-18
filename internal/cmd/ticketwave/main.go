package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"ticket_wave/internal/models"
	"ticket_wave/internal/repository"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %q", err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	defer db.Close()

	repo := repository.NewRepository(db)

	// Контрольная точка 2

	participant1 := models.Participant{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		BirthDate: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Username:  "johndoe",
		Password:  "password1",
		Role:      "user",
	}

	participant2 := models.Participant{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
		BirthDate: time.Date(1992, time.February, 2, 0, 0, 0, 0, time.UTC),
		Username:  "janesmith",
		Password:  "password2",
		Role:      "organizer",
	}

	event1 := models.Event{
		ID:          1,
		OrganizerID: 1,
		Name:        "Some Band Concert",
		Description: "Whatever",
		Category:    "Music",
		DateTime:    time.Date(2022, time.December, 31, 20, 0, 0, 0, time.UTC),
		Venue:       "Стадион",
		Address:     "Улица Пушкина, дом Колотушкина",
		TicketPrice: 1000.0,
	}

	event2 := models.Event{
		ID:          2,
		OrganizerID: 2,
		Name:        "Выставка",
		Description: "Выставка современного искусства",
		Category:    "Искусство",
		DateTime:    time.Date(2023, time.January, 1, 10, 0, 0, 0, time.UTC),
		Venue:       "Галерея",
		Address:     "Проспект Мира, дом Мира",
		TicketPrice: 500.0,
	}

}
