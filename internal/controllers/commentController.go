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
func GetCommentByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PostID, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}
	data, err := services.GetComment(r.Context(), PostID)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured When Query Data",
			Error:   err.Error(),
		})
		return
	}

	if len(data) <= 0 {
		helper.WriteJSON(w, http.StatusOK, models.ApiResponse{
			Success: true,
			Message: "Success Get Comment",
			Data:    "No Comment Found",
		})
		return
	}

	helper.WriteJSON(w, http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "Successfully Get Comment",
		Data:    data,
	})
}

func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CommentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Error Occured",
			Error:   err.Error(),
		})
		return
	}

	var UpdateComment models.UpdateComment
	err = json.NewDecoder(r.Body).Decode(&UpdateComment)

	err = services.UpdateComment(r.Context(), UpdateComment.Comment, CommentID)
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
		Message: "Successfully Update Comment!",
	})
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
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

	err = services.DeleteComment(r.Context(), ContentID)
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
		Message: "Successfully Delete Comment!",
	})
}
