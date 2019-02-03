package db

import (
	"fmt"
	"github.com/Ayupov-Ayaz/go-web/model"
	"strconv"
	"strings"
)

func (db *DB) SelectPostById(id int64) (*model.Post, error) {
	post := &model.Post{}
	query := "select * from posts where id = ?"
	if err := db.Get(post, query, id); err != nil {
		return nil, err
	}
	return post, nil
}
// localhost:port/graphql?query=mutation{create_post(title:"title",description:"description",author:1){id,title}}
func (db *DB) InsertPost(post *model.Post) (*model.Post, error) {
	query := "insert into posts (title, description, author_id) values (:title, :description, :author_id)"
	result, err := db.NamedExec(query, post)
	if err != nil {
		return nil, err
	}
	if lastId, err := result.LastInsertId(); err != nil {
		return nil, err
	} else {
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

func (db *DB) UpdatePost(post *model.Post) (*model.Post, error) {
	postParams := make([]string, 0)
	if post.Title != "" {
		postParams = append(postParams, " title = :title ")
	}
	if post.Description != "" {
		postParams = append(postParams," description = :description ")
	}
	if post.Author != 0 {
		postParams = append(postParams," author_id = :author_id " )
	}
	query := "UPDATE posts SET " + strings.Join(postParams, ", ") + " where id = " + strconv.FormatInt(post.ID,
		10)
	fmt.Println(query)
	_, err := db.NamedExec(query, post)
	if err != nil { return nil, err }
	return db.SelectPostById(post.ID)
}
