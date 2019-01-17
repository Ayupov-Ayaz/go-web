package commands

import (
	"go-web/db"
	"go-web/model"
)

func GetUsers(db *db.DB) ([]*model.User, error) {
	query := `select id, login, email from users`
	users := make([]*model.User, 0)
	err := db.Select(users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
