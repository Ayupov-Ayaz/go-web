package model

type Model struct {
	db
}


func New(db db) *Model{
	return &Model{
		db: db,
	}
}

func (m *Model) GetUsers() ([]*User, error) {
	return db.SelectAllUsers(m.db)
}

func (m *Model) GetUserById(id int64) (*User, error) {
	return db.SelectUserById(m.db, id)
}