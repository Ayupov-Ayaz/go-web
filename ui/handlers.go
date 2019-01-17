package ui

import (
	"github.com/Ayupov-Ayaz/go-web/model"
	"net/http"
)

func IndexHandler(m *model.Model) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		})
}