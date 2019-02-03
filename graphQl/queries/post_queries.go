package queries

import (
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/graphQl/types"
	"github.com/graphql-go/graphql"
)

func GetPostById(db *db.DB) *graphql.Field {
	return &graphql.Field{
		Type:        types.GetPostType(),
		Description: "Get post by id",
		// Какие аргументы необходимы для получения поста
		Args: graphql.FieldConfigArgument{
			// В данном случае нам необходи только id
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		// как мы будем получать пост
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			// если это не дефолтное значение
			if ok {
				// тянем из бд продукт
				return db.SelectPostById(int64(id))
			}
			// id не передан
			return nil, nil
		},
	}
}

func GetPostByTitle(db *db.DB) *graphql.Field {
	return &graphql.Field{
		Type: types.GetPostType(),
		Description: "Get post by name",
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			if title, ok := p.Args["title"].(string); ok {
				return db.SelectPostByTitle(title)
			}
			return nil, nil
		},
	}
}
