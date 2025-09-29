package services

import (
	"context"
	"fmt"

	"github.com/Dionisius951/Golang-Blog-BE/internal/db"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
)

func AddComment(ctx context.Context, data *models.Comment) error {
	sql :=
		`
	INSERT INTO "comment" (post_id,username_id,content) VALUES ($1,$2,$3)
		`
	_, err := db.Pool.Query(ctx, sql, data.PostID, data.UsernameID, data.Comment)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
