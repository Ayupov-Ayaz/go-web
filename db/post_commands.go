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

func (db *DB) InsertPost(post *model.Post) (lastId int64, err error) {
	query := "insert into posts (title, description, author_id) values (:title, :description, :author_id)"
	result, err := db.NamedExec(query, post)
	if err != nil {
		return 0, err
	}
	if lastId, err := result.LastInsertId(); err != nil {
		return 0, err
	} else {
		return lastId, nil
	}
}
