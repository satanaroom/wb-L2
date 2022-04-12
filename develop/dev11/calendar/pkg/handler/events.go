package handler

import (
	"net/http"
	"strconv"
)

// StatusBadRequest                    = 400
// StatusInternalServerError           = 500
// StatusServiceUnavailable            = 503

func (h *Handler) createEvent(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		newErrorResponse(res, http.StatusServiceUnavailable, "wrong method")
		return
	}
	input, err := jsonParser(req)
	if err != nil {
		newErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateEvent(input)
	if err != nil {
		newErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(res, http.StatusOK, map[string]interface{}{
		"result": id,
	})
}

func (h *Handler) updateEvent(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		newErrorResponse(res, http.StatusServiceUnavailable, "wrong method")
		return
	}
	input, err := jsonParser(req)
	if err != nil {
		newErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.UpdateEvent(input)
	if err != nil {
		newErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}
	responseId := strconv.Itoa(input.Id)
	newResponse(res, http.StatusOK, map[string]interface{}{
		"result": responseId + " updated",
	})
}

func (h *Handler) deleteEvent(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		newErrorResponse(res, http.StatusServiceUnavailable, "wrong method")
		return
	}
	id, err := getEventId(req)
	if err != nil {
		newErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.DeleteEvent(id)
	if err != nil {
		newErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}
	responseId := strconv.Itoa(id)
	newResponse(res, http.StatusOK, map[string]interface{}{
		"result": responseId + " deleted",
	})
}

func (h *Handler) eventsForDay(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		newErrorResponse(res, http.StatusServiceUnavailable, "wrong method")
		return
	}
	input, err := queryParser(req)
	if err != nil {
		newErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	eventsForDay(input)
	events, err := h.services.EventsForDay(input)
	if err != nil {
		newErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(res, http.StatusOK, eventsForWeekResponse{
		Data: events,
	})
}

func (h *Handler) eventsForWeek(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		newErrorResponse(res, http.StatusServiceUnavailable, "wrong method")
		return
	}
	input, err := queryParser(req)
	if err != nil {
		newErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	eventsForWeek(input)
	events, err := h.services.EventsForWeek(input)
	if err != nil {
		newErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(res, http.StatusOK, eventsForWeekResponse{
		Data: events,
	})
}

func (h *Handler) eventsForMonth(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		newErrorResponse(res, http.StatusServiceUnavailable, "wrong method")
		return
	}
	input, err := queryParser(req)
	if err != nil {
		newErrorResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	eventsForMonth(input)
	events, err := h.services.EventsForMonth(input)
	if err != nil {
		newErrorResponse(res, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(res, http.StatusOK, eventsForWeekResponse{
		Data: events,
	})
}
