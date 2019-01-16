package model

import (
	"go-web/comands"
	"go-web/db"
)

// структура которая будет иметь подключение к бд и будет выступать в качестве CRUD
// для наших структур хранящихся в бд
type Model struct {
	db.DB // расширяем бд
}

func New(db db.DB) *Model{
	return &Model{
		db,
	}
}

func SelectAllUsers() []*User{
	userCommands := comands.UserCommands{}
	users, err := userCommands.GetUsers()
	if err != nil {
		panic(err)
	}
	return users
}
