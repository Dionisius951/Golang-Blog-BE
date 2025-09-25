package models

type Post struct {
	Id int `json:"id,omiempty"`
	Title string `json:"title"`
	Content string `json:"content"`
	UsernameId int `json:"username_id"`
}