package mutation

import (
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/errors"
	"github.com/Ayupov-Ayaz/go-web/graphQl/types"
	"github.com/Ayupov-Ayaz/go-web/model"
	"github.com/graphql-go/graphql"
	"strconv"
)

// Создание post
func CreatePost(db *db.DB) *graphql.Field {
	return &graphql.Field{
		Type: types.GetPostType(),
		Description: "Create new Post",
		Args: graphql.FieldConfigArgument{
			// перечисляем наши поля
			"title" 	  : types.GqNewNonNullString(),
			"description" : types.GqNewNonNullString(),
			"author" 	  : types.GqNewNonNullString(),

		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			authorId := p.Args["author"].(int)
			post := &model.Post{
				Title: p.Args["title"].(string),
				Description: p.Args["description"].(string),
				Author: int64(authorId),
			}
			return db.InsertPost(post)
		},
	}
}

func UpdatePost(db *db.DB) *graphql.Field {
	return &graphql.Field{
		Type: types.GetPostType(),
		Description: "Update post",
		Args: graphql.FieldConfigArgument {
			// перечисляем наши поля
			"title" 	  : types.GqNewNullString(),
			"description" : types.GqNewNullString(),
			"author" 	  : types.GqNewNullString(),
			"id"		  : types.GqNewNonNullString(),
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			title, titleOk := p.Args["title"].(string)
			description, descriptionOk := p.Args["description"].(string)
			author, parseErr := strconv.ParseInt(p.Args["author"].(string), 10, 64)
			post := &model.Post{}
			empty := true
			if titleOk {
				post.Title = title
				empty = false
			}
			if descriptionOk {
				post.Description = description
				empty = false
			}
			if parseErr == nil && author != 0 {
				post.Author = author
				empty = false
			}
			if empty {
				err := errors.NewError("No parameter passed")
				return nil, err
			}
			if id, parseErr := strconv.ParseInt(p.Args["id"].(string), 10, 64); id != 0 &&
				parseErr == nil {
				post.ID = id
				return db.UpdatePost(post)
			} else {
				return nil, errors.NewError("Id is not Int!")
			}
		},
	}
}


