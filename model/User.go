package model

type User struct {
	ID int `db:"id"`
	Login string `db:"login"`
	Password string `db:"password"`
	Email string `db:"email"`
}
