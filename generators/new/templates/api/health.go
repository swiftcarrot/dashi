package api

import "net/http"

func (server *Server) HealthCheck(w http.ResponseWriter, r *http.Request) error {
	return WriteJSONResponse(w, http.StatusOK, map[string]bool{
		"ok": true,
	})
}
