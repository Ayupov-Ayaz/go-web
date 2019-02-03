package ui

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"net/http"
)

func GraphqlHandler(schema graphql.Schema) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		query  := r.URL.Query().Get("query")
		gqSchema := executorQuery(query, schema)
		json.NewEncoder(w).Encode(gqSchema)
	})
}