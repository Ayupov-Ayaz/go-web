package types

import (
	"github.com/graphql-go/graphql"
)

func GetQueryPostType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "post",
		Fields: graphql.Fields{
			"id"		  : GqInt(),
			"title"		  : GqString(),
			"description" : GqString(),
			"author"	  : GqInt(),
		},
	})
}
