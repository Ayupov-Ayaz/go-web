package main

import (
	"flag"
	"github.com/Ayupov-Ayaz/go-web/daemon"
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/db/migrations"
	"net/http"
	"os"
)
var assetsPath string

func processFlags() * daemon.Config {
	cfg := &daemon.Config{}
	cfg.Db = &db.Config{ Driver: "mysql", User: "tommy", Password: "43", Host: "127.0.0.1", Database: "go_lang", Port: 3306}
	//// Если main запущен с аргументами, то запускаем наши миграции
	if len(os.Args) > 1 {
		migrations.Migration(os.Args[1], cfg.Db)
		os.Exit(-1)
	}
	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
	flag.StringVar(&assetsPath, "assets-path", "assets", "Path to static dir")
	flag.Parse()
	return cfg
}

func setupHttpAssets(cfg *daemon.Config) {
	cfg.UI.Assets = http.Dir(assetsPath)
}

func main() {
	cfg := processFlags()
	setupHttpAssets(cfg)
	daemon.Run(cfg)
}