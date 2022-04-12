package service

import (
	"github.com/satanaroom/calendar"
	"github.com/satanaroom/calendar/pkg/repository"
)

type Event interface {
	CreateEvent(event *calendar.Event) (int, error)
	UpdateEvent(event *calendar.Event) error
	DeleteEvent(eventId int) error
	EventsForDay(event *calendar.Event) ([]calendar.ResultEvent, error)
	EventsForWeek(event *calendar.Event) ([]calendar.ResultEvent, error)
	EventsForMonth(event *calendar.Event) ([]calendar.ResultEvent, error)
}

type Service struct {
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Event: NewCalendarService(repos.Calendar),
	}
}
