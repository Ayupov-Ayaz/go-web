package ui

import (
	"github.com/graphql-go/graphql"
	"net/http"
)

func routes(schema graphql.Schema) {
	// graphQl схема
	http.Handle("/graphql", GraphqlHandler(schema))
}