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
			return nil, fmt.Errorf("%v \n", err)
		}
		data = append(data, post)
	}
	return data, nil
}

func GetPostById(ctx context.Context, id int) ([]models.Post, error) {
	sql := `SELECT * FROM "post" WHERE username_id=$1`
	rows, err := db.Pool.Query(ctx, sql, id)
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
			return nil, fmt.Errorf("%v \n", err)
		}
		data = append(data, post)
	}

	return data, nil
}

func UpdatePost(ctx context.Context, title *string, content *string, id int) error {
	sql := `
	UPDATE "post" SET 
		title=COALESCE($1, title), 
		content=COALESCE($2, content)
		WHERE id=$3
		`
	cmd, err := db.Pool.Exec(ctx, sql, title, content, id)
	if err != nil {
		return fmt.Errorf("Error Updating Content: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("No Content Found With Username Id : %v", id)
	}

	return nil
}

func DeletePost(ctx context.Context, id int) error {
	sql := `
	DELETE FROM "post" WHERE id=$1
		`
	cmd, err := db.Pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("Error Deleting Content: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("No Content Found With Username Id : %v", id)
	}

	return nil
}
