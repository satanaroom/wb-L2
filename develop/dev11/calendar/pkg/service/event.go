package service

import (
	"github.com/satanaroom/calendar"
	"github.com/satanaroom/calendar/pkg/repository"
)

type EventService struct {
	repo repository.Calendar
}

func NewCalendarService(repo repository.Calendar) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(event *calendar.Event) (int, error) {
	return s.repo.CreateEvent(event)
}
func (s *EventService) UpdateEvent(event *calendar.Event) error {
	return s.repo.UpdateEvent(event)
}
func (s *EventService) DeleteEvent(eventId int) error {
	return s.repo.DeleteEvent(eventId)
}
func (s *EventService) EventsForDay(event *calendar.Event) ([]calendar.ResultEvent, error) {
	return s.repo.EventsForDay(event)
}
func (s *EventService) EventsForWeek(event *calendar.Event) ([]calendar.ResultEvent, error) {
	return s.repo.EventsForWeek(event)
}
func (s *EventService) EventsForMonth(event *calendar.Event) ([]calendar.ResultEvent, error) {
	return s.repo.EventsForMonth(event)
}
