package model

type Model struct {
	db
}


func New(db db) *Model{
	return &Model{
		db: db,
	}
}

func (m *Model)SelectAllUsers() ([]*User, error){

	return nil, nil
}
