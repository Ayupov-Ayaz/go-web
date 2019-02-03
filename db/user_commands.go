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
	query := "select id, login, email, age from users where id = ?"
	err := db.Get(user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) InsertUser(user *model.User) (lastId int64, err error) {
	query := "insert into users (login, password, email, age) values (:login, :password, :email, :age)"
	result, err := db.NamedExec(query, user)
	if err != nil {
		return 0, err
	}
	if lastId, err := result.LastInsertId(); err != nil {
		return 0, err
	} else {
		return lastId, nil
	}
}