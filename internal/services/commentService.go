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

func GetComment(ctx context.Context, id int) ([]models.GetComment, error) {
	sql :=
		`
		SELECT cm.id, cm.post_id,ps.title, cm.username_id, usr.username, cm.content FROM "comment" cm 
			JOIN "user" usr ON cm.username_id=usr.id 
			JOIN "post" ps ON cm.post_id=ps.id
			WHERE cm.post_id=$1
		`
	rows, err := db.Pool.Query(ctx, sql, id)
	if err != nil {
		return nil, fmt.Errorf("%v \n", err)
	}

	var data []models.GetComment

	for rows.Next() {
		var comment models.GetComment
		err := rows.Scan(
			&comment.Id,
			&comment.PostID,
			&comment.PostTitle,
			&comment.UsernameID,
			&comment.Username,
			&comment.Comment,
		)
		if err != nil {
			return nil, fmt.Errorf("%v \n", err)
		}
		data = append(data, comment)
	}
	return data, nil

}

func UpdateComment(ctx context.Context, content *string, id int) error {
	sql := `
	UPDATE "comment" SET 
		content=COALESCE($1, content)
		WHERE id=$2
		`
	cmd, err := db.Pool.Exec(ctx, sql, content, id)
	if err != nil {
		return fmt.Errorf("Error Updating Content: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("No Content Found With Username Id : %v", id)
	}

	return nil
}

func DeleteComment(ctx context.Context, id int) error {
	sql := `
	DELETE FROM "comment" WHERE id=$1
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
