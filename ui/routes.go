package ui

import (
	"github.com/Ayupov-Ayaz/go-web/model"
	"github.com/graphql-go/graphql"
	"net/http"

)

func routes(m *model.Model, schema graphql.Schema) {
	http.Handle("/", IndexHandler(m))
	//Страница регистрации
	http.Handle("/sign/up", ShowRegisterFormHandler())
	// Отправка формы регистрации post
	http.Handle("/registration", RegisterHandler(m))
	// graphQl схема
	http.Handle("/graphql", GraphqlHandler(schema))
}