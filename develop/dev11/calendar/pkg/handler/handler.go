package handler

import (
	"net/http"

	"github.com/satanaroom/calendar/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.eventsForDay)
	mux.HandleFunc("/events_for_week", h.eventsForWeek)
	mux.HandleFunc("/events_for_month", h.eventsForMonth)

	return mux
}
