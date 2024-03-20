package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
		Venue:       "Some Venue",
		Address:     "Some Address",
		TicketPrice: 1000.0,
	}

	event2 := models.Event{
		ID:          2,
		OrganizerID: 1,
		Name:        "Art Exhibition",
		Description: "Whatever",
		Category:    "Art",
		DateTime:    time.Date(2023, time.January, 1, 10, 0, 0, 0, time.UTC),
		Venue:       "Some Art Gallery",
		Address:     "Some Address",
		TicketPrice: 500.0,
	}

	participant1, err = repo.CreateParticipant(participant1)
	if err != nil {
		return
	}

	participant2, err = repo.CreateParticipant(participant2)
	if err != nil {
		return
	}

	event1, err = repo.CreateEvent(event1)
	if err != nil {
		return
	}

	event2, err = repo.CreateEvent(event2)
	if err != nil {
		return
	}

	feedback1 := models.EventFeedback{
		ID:       1,
		EventID:  1,
		SenderID: 1,
		Text:     "Super cool event! Loved it!",
	}

	feedback2 := models.EventFeedback{
		ID:       2,
		EventID:  2,
		SenderID: 2,
		Text:     "Disappointing. Not worth the money.",
	}

	err = repo.CreateEventFeedback(feedback1)
	if err != nil {
		return
	}

	err = repo.CreateEventFeedback(feedback2)
	if err != nil {
		return
	}

	userEventLink1 := models.UserEventLink{
		UserID:   1,
		EventID:  1,
		LinkType: "like",
	}

	userEventLink2 := models.UserEventLink{
		UserID:   2,
		EventID:  2,
		LinkType: "going",
	}

	err = repo.CreateUserEventLink(userEventLink1)
	if err != nil {
		return
	}

	err = repo.CreateUserEventLink(userEventLink2)
	if err != nil {
		return
	}

	ticket1 := models.Ticket{
		ID:           1,
		EventID:      1,
		OwnerID:      1,
		PurchaseDate: time.Now(),
		SeatNumber:   "A1",
	}

	ticket2 := models.Ticket{
		ID:           2,
		EventID:      2,
		OwnerID:      2,
		PurchaseDate: time.Now(),
		SeatNumber:   "B2",
	}

	ticket1, err = repo.CreateTicket(ticket1)

	if err != nil {
		return
	}

	ticket2, err = repo.CreateTicket(ticket2)

	if err != nil {
		return
	}
}
