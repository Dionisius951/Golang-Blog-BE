package models

type Comment struct {
	Id int `json:"id"`
	PostID int `json:"post_id"`
	UsernameID int `json:"username_id"`
	Comment string `json:"comment"`
}

type UpdateComment struct {
	Comment *string `json:"comment"`
}