package comands

import (
	"go-web/model"
)

type UserCommands struct {
	model.Model
}

func New() *UserCommands {
	return &UserCommands{}
}

func (uc *UserCommands) GetUsers() ([]*model.User, error) {
	query := `select id, login, email from users`
	users := make([]*model.User, 0)
	err := uc.Model.DB.Select(users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}


