package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, status int, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(value)
	if err != nil {
		return InternalServerError(fmt.Sprintf("Error encoding json response: %v", value))
	}
	w.WriteHeader(status)
	_, err = w.Write(b)
	return nil
}
