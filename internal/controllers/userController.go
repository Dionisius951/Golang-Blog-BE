package controllers

import (
	"net/http"

	"github.com/Dionisius951/Golang-Blog-BE/internal/helper"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	helper.WriteJSON(w, http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "halooo",
	})
}