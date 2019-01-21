package model

type Post struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Author int64 `db:"author_id" json:"author_id"`
}