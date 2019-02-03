package db

import (
	"github.com/Ayupov-Ayaz/go-web/model"
	"log"
)

func (db *DB) SelectPostById(id int64) (*model.Post, error) {
	post := &model.Post{}
	query := "select * from posts where id = ?"
	if err := db.Get(post, query, id); err != nil {
		log.Println("ERROR| SelectPostById: \n %v", err.Error())
		return nil, err
	}
	return post, nil
}
// localhost:port/graphql?query=mutation{create_post(title:"title",description:"description",author:1){id,title}}
func (db *DB) InsertPost(post *model.Post) (user *model.Post, err error) {
	query := "insert into posts (title, description, author_id) values (:title, :description, :author_id)"
	result, err := db.NamedExec(query, post)
	if err != nil {
		return nil, err
	}
	if lastId, err := result.LastInsertId(); err != nil {
		return nil, err
	} else {
		lastId = int64(lastId)
		return db.SelectPostById(lastId)
	}
}

func (db *DB) SelectPostByTitle(title string) (*model.Post, error) {
	post := &model.Post{}
	query := "SELECT * FROM posts where title = ?"
	if err := db.Get(post, query, title); err != nil {
		return nil, err
	}
	return post, nil
}