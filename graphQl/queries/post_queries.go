package queries

import (
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/graphQl/types"
	"github.com/graphql-go/graphql"
)

func GetPostById(db *db.DB) *graphql.Field {
	return &graphql.Field{
		Type:        types.GetQueryPostType(),
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

func GetPostByName(db *db.DB) *graphql.Field {
	return &graphql.Field{
		// TODO: реализовать
	}
}
