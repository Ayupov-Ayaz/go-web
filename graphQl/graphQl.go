package graphQl

import (
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/graphQl/mutation"
	"github.com/Ayupov-Ayaz/go-web/graphQl/queries"
	"github.com/graphql-go/graphql"
)

// Все возможные вариации для мутаций
func NewGraphQlSchema(db *db.DB) graphql.Schema {
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: getQueryType(db),
		Mutation: getMutationType(db),
	})
	return schema
}

// Все возможные вариации запросов Select
func getQueryType(db *db.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"post_by_id"	: queries.GetPostById(db),
			"post_by_name"  : queries.GetPostByName(db),
		},
	})
}

// Все возможные мутации
func getMutationType(db *db.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutations",
		Fields: graphql.Fields{
			"create_post" : mutation.CreatePost(db),
		},
	})
}


