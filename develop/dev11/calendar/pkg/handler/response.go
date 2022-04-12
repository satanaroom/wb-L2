package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/satanaroom/calendar"
)

const encodingJSON = "[json] error encoding json"

type errorMessage struct {
	Error string `json:"error"`
}

type eventsForWeekResponse struct {
	Data []calendar.ResultEvent `json:"result"`
}

type Id struct {
	Id int `json:"id"`
}

func newResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		fmt.Fprintf(w, "%s", errors.New(encodingJSON))
	}
}

func newErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	errorMsg := errorMessage{
		Error: errMsg,
	}
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(errorMsg)
	if err != nil {
		fmt.Fprintf(w, "%s", errors.New(encodingJSON))
	}

}
