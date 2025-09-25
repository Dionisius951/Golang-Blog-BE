package services

import (
	"context"
	"fmt"

	"github.com/Dionisius951/Golang-Blog-BE/internal/db"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
)

func AddPost(ctx context.Context, content *models.Post) error {
	sql := `INSERT INTO "post" (title, content, username_id) VALUES ($1,$2,$3)`
	_, err := db.Pool.Query(ctx, sql, content.Title, content.Content, content.UsernameId)
	if err != nil {
		return fmt.Errorf("%v \n", err)
	}
	return nil
}

func GetAllPost(ctx context.Context) ([]models.Post, error) {
	sql := `SELECT * FROM "post"`
	rows, err := db.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("%v \n", err)
	}
	var data []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.UsernameId,
		)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		data = append(data, post)
	}
	return data, nil
}
