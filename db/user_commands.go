package db

import (
	"github.com/Ayupov-Ayaz/go-web/model"
)

func  (db *DB) SelectAllUsers() (*[]model.User, error) {
	users := &[]model.User{}
	query := "select id, login, email, age from users"
	err := db.Select(users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *DB) SelectUserById(id int64) (*model.User, error){
	user := &model.User{}
	query := "select id, login, email, age from users limit ?"
	err := db.Get(user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
