package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/satanaroom/calendar"
)

type Calendar interface {
	CreateEvent(event *calendar.Event) (int, error)
	UpdateEvent(event *calendar.Event) error
	DeleteEvent(eventId int) error
	EventsForDay(event *calendar.Event) ([]calendar.ResultEvent, error)
	EventsForWeek(event *calendar.Event) ([]calendar.ResultEvent, error)
	EventsForMonth(event *calendar.Event) ([]calendar.ResultEvent, error)
}

type Repository struct {
	Calendar
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Calendar: NewEventPostgres(db),
	}
}
