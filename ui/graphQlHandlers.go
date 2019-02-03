package ui

import (
	"encoding/json"
	"github.com/Ayupov-Ayaz/go-web/errors"
	"github.com/graphql-go/graphql"
	"net/http"
)

func GraphqlHandler(schema graphql.Schema) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		query  := r.FormValue("query")
		gqSchema := executorQuery(query, schema)
		err := json.NewEncoder(w).Encode(gqSchema)
		if err != nil {
			errors.PrintSystemErr("Не удалось преобразовать graphQl схему")
		}
	})
}