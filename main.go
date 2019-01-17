package main

import (
	"flag"
	"go-web/daemon"
	"go-web/db"
	"go-web/db/migrations"
	"log"
	"net/http"
	"os"
)
var assetsPath string

func processFlags() * daemon.Config {
	cfg := &daemon.Config{}
	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
	cfg.Db = &db.Config{ Driver: "mysql", User: "tommy", Password: "43", Host: "127.0.0.1", Database: "go_lang", Port: 3306}
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
	//// Если main запущен с аргументами, то запускаем наши миграции
	if len(os.Args) > 1 {
		migrations.Migration(os.Args[1], cfg)
		os.Exit(-1)
	}

	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main: \n %s", err.Error())
	}
}