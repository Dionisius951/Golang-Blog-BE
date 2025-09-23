package helper

import (
	"encoding/json"
	"net/http"

	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
)

func WriteJSON(w http.ResponseWriter, status int, response models.ApiResponse) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}