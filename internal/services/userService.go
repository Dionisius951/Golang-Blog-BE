package services

import (
	"context"
	"fmt"

	"github.com/Dionisius951/Golang-Blog-BE/internal/db"
	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(ctx context.Context, data *models.Users) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	if err != nil {
		fmt.Printf("Error While Generate Password, error : %v", err.Error())
	}

	sql := `INSERT INTO "user" (username, email, password, role_id) VALUES ($1,$2,$3,2)`
	_, err = db.Pool.Query(ctx, sql, data.Username, data.Email, string(hashPassword))
	if err != nil {
		return fmt.Errorf("Error Occured : %v \n", err)
	}
	return nil
}

func LoginUser(ctx context.Context, email string, password string) (*models.Users, error) {
	var user models.Users

	sql := `SELECT * FROM "user" WHERE email=$1`
	err := db.Pool.QueryRow(ctx, sql, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, fmt.Errorf("Email Not Register")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Password Incorect!")
	}

	return &user, nil
}
