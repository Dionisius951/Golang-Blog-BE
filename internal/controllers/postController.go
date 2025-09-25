package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Dionisius951/Golang-Blog-BE/internal/helper"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
	"github.com/Dionisius951/Golang-Blog-BE/internal/services"
)

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	var content models.Post

	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}

	err = services.AddPost(r.Context(), &content)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "Content Created",
	})
}

func GetAllPostHandler(w http.ResponseWriter, r *http.Request) {
	data, err := services.GetAllPost(r.Context())
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "Success Get Content",
		Data: data,
	})
}
