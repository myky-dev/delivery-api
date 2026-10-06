package handler

import (
	"encoding/json"
	"net/http"
)

// VersionResponse — тело ответа эндпоинта /version.
type VersionResponse struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

func version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := VersionResponse{Service: "delivery-api", Version: "0.1.0"}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
