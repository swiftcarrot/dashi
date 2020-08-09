package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		WriteError(w, err)
	}
}

func WriteJSONResponse(w http.ResponseWriter, status int, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(value)
	if err != nil {
		return internalServerError(fmt.Sprintf("Error encoding json response: %v", value))
	}
	w.WriteHeader(status)
	_, err = w.Write(b)
	return nil
}
