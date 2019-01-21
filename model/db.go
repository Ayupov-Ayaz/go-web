package model
/**
 	All queries for all types
 */
type db interface {
	// user
	SelectAllUsers() (*[]User, error)
	SelectUserById(id int64) (*User, error)
	InsertUser(user * User) (lastId int64, err error)

	// post
	InsertPost(post *Post) (lastId int64, err error)
	SelectPostById(id int64) (*Post, error)
}