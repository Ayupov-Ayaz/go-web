package main

import (
	"github.com/Ayupov-Ayaz/go-web/daemon"
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/db/migrations"
	"github.com/Ayupov-Ayaz/go-web/ui"
	"os"
)

func processFlags() *daemon.Config {
	cfg := &daemon.Config{}
	cfg.Db = db.GetDbConfig()

	//// Если main запущен с аргументами, то запускаем наши миграции
	if len(os.Args) > 1 {
		migrations.Migration(os.Args[1], cfg.Db)
		os.Exit(-1)
	}

	cfg.ListenSpec = "localhost:3000"
	cfg.UI = ui.GetUiConfig()
	return cfg
}


func main() {
	cfg := processFlags()
	daemon.Run(cfg)
}