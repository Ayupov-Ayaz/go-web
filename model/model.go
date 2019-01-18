package model

type Model struct {
	db
}


func New(db db) *Model{
	return &Model{
		db: db,
	}
}
