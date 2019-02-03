package ui

import (
	"fmt"
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/errors"
	"github.com/Ayupov-Ayaz/go-web/graphQl"
	"github.com/graphql-go/graphql"
	"log"
	"net"
	"net/http"
	"time"
)

type Config struct {
	AssetsPath string
	AssetsPrefix string
}

func Start(cfg *Config, db *db.DB,  listner *net.Listener) {
	server := &http.Server{
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 16}

	http.Handle("/assets/", http.StripPrefix(cfg.AssetsPrefix, http.FileServer(http.Dir(cfg.AssetsPath))))

	schema := graphQl.NewGraphQlSchema(db)

	routes(schema)
	if err := server.Serve(*listner); err != nil {
		errors.PrintSystemErr(fmt.Sprintf("Не удалось запустить сервер: \n %s", err.Error()))
	}
}


func executorQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("ERROR |main/executorQuery: %v ", result.Errors)
	}
	return result
}

func GetUiConfig() *Config{
	return &Config{
		AssetsPrefix: "/assets/",
		AssetsPath: "./assets/",
	}
}
