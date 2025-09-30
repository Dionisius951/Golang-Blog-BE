package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Dionisius951/Golang-Blog-BE/internal/helper"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			helper.WriteJSON(w, http.StatusUnauthorized, models.ApiResponse{
				Success: false,
				Message: "Unauthorized",
			})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		user, err := helper.ValidateToken(token)
		if err != nil {
			helper.WriteJSON(w, http.StatusUnauthorized, models.ApiResponse{
				Success: false,
				Message: "Error When Validating Token",
				Error:   err.Error(),
			})
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
