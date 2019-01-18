package model

type db interface {
	SelectAllUsers() (*[]User, error)
	SelectUserById(id int64) (*User, error)
}