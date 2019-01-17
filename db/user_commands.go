package db

import (
	"go-web/model"
)

// Получение всех пользователей
func  (db *DB) SelectAllUsers() ([]*model.User, error) {
	users := make([]*model.User, 0)
	query := "select id, login, email from users"
	if err := db.Select(users, query); err != nil {
		return nil, err
	}
	return users, nil
}