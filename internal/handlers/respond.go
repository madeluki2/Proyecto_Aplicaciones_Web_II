package handlers

import (
	"encoding/json"
	"net/http"
)

// respondJSON serializa `data` como JSON y escribe el status code en la respuesta.
func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError envía un JSON de error con el formato estándar { "error": "mensaje" }.
func respondError(w http.ResponseWriter, status int, mensaje string) {
	respondJSON(w, status, map[string]string{"error": mensaje})
}
