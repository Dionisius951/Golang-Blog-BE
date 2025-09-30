package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Dionisius951/Golang-Blog-BE/internal/helper"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
	"github.com/Dionisius951/Golang-Blog-BE/internal/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var data models.Users

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error while process data",
			Error: err.Error(),
		})
		return
	}

	err = services.RegisterUser(r.Context(), &data)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error while process data",
			Error: err.Error(),
		})
		return
	}

	helper.WriteJSON(w, http.StatusCreated, models.ApiResponse{
		Success: true,
		Message: "Success Register User!",
	})
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data models.Users

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error while process data",
			Error: err.Error(),
		})
		return
	}

	err = services.LoginUser(r.Context(), &data)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error while process data",
			Error: err.Error(),
		})
		return
	}
	usr, err := helper.GenerateToken(&data)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error while process data",
			Error: err.Error(),
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "Login Success!",
		Data : usr,
	})
}