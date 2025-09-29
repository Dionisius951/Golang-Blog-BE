package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Dionisius951/Golang-Blog-BE/internal/helper"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
	"github.com/Dionisius951/Golang-Blog-BE/internal/services"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	var data models.Comment

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}

	err = services.AddComment(r.Context(), &data)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured When Add Data",
			Error:   err.Error(),
		})
		return
	}

	helper.WriteJSON(w, http.StatusCreated, models.ApiResponse{
		Success: true,
		Message: "Comment Created",
	})
}
