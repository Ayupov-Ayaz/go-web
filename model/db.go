package model

type db interface {
	SelectAllUsers() ([]*User, error)
}