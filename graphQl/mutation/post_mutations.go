package mutation

import (
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/graphQl/types"
	"github.com/Ayupov-Ayaz/go-web/model"
	"github.com/graphql-go/graphql"
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
			"author" 	  : types.GqNewNonNullInt(),

		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			post := &model.Post{
				Title: p.Args["title"].(string),
				Description: p.Args["description"].(string),
				Author: p.Args["author_id"].(int64),
			}
			return db.InsertPost(post)
		},
	}
}


