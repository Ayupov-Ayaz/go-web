package model

type Post struct {
	ID int64
	Title string
	Description string
	Author int64 `json:"author_id"`
}