package ui

import (
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/model"
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

	routes(m)
	server.Serve(*listner)
}

func routes(m *model.Model) {
	http.Handle("/", IndexHandler(m))
	http.Handle("/sign/up", ShowRegisterFormHandler())
	http.Handle("/registration", RegisterHandler(m))
}
