package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dionisius951/Golang-Blog-BE/internal/helper"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
	"github.com/Dionisius951/Golang-Blog-BE/internal/services"
	"github.com/gorilla/mux"
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

	helper.WriteJSON(w, http.StatusCreated, models.ApiResponse{
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
		Data:    data,
	})
}

func GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ContentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}
	data, err := services.GetPostById(r.Context(), ContentID)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}

	if len(data) <= 0 {
		helper.WriteJSON(w, http.StatusOK, models.ApiResponse{
			Success: true,
			Message: "Success Get Content",
			Data:    "No Content Found",
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "Successfully Get Content",
		Data:    data,
	})
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ContentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}

	var UpdatePost models.UpdatePost
	err = json.NewDecoder(r.Body).Decode(&UpdatePost)

	err = services.UpdatePost(r.Context(), UpdatePost.Title, UpdatePost.Content, ContentID)
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
		Message: "Successfully Update Content!",
	})
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ContentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}

	err = services.DeletePost(r.Context(), ContentID)
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
		Message: "Successfully Delete Content!",
	})
}
