package handler

import (
	"encoding/json"
	"net/http"
)

// HealthResponse — тело ответа эндпоинта /health.
type HealthResponse struct {
	Status string `json:"status"`
}

// Register вешает все HTTP-обработчики на mux.
// Роутинг живёт в одном месте, main.go не знает про конкретные пути.
func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /version", version)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := HealthResponse{Status: "ok"}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
