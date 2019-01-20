package ui

import (
	"fmt"
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/graphQl"
	"github.com/Ayupov-Ayaz/go-web/model"
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

	m := model.New(db)
	schema := graphQl.NewGraphQlSchema(db)

	routes(m, schema)
	server.Serve(*listner)
}


func executorQuery(query string, schema graphql.Schema) *graphql.Result {
	fmt.Println(query, schema)
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("ERROR |main/executorQuery: %v ", result.Errors)
	}
	return result
}