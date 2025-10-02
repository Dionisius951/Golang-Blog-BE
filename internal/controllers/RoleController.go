package controllers

import (
	"context"
	"fmt"

	"github.com/Dionisius951/Golang-Blog-BE/internal/helper"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
)

func CheckRoleHandler(ctx context.Context, RequiredRole []int) error {
	claims := ctx.Value("user").(*models.MyCustomClaims)

	err := helper.CheckRole(RequiredRole, claims.Role)
	if err != nil {
		return fmt.Errorf("Forbidden Action")
	}

	return nil
}
