package model

type Post struct {
	ID int64
	Title string
	Description string
	Author int64 `db:"author_id" json:"author_id"`
}