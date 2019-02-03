package main

import (
	"github.com/Ayupov-Ayaz/go-web/configs"
	"github.com/Ayupov-Ayaz/go-web/daemon"
	"github.com/Ayupov-Ayaz/go-web/db/migrations"
	"os"
)

func processFlags() *daemon.Config {
	cfg := &daemon.Config{}
	cfg.Db = configs.GetDbConfig()

	//// Если main запущен с аргументами, то запускаем наши миграции
	if len(os.Args) > 1 {
		migrations.Migration(os.Args[1], cfg.Db)
		os.Exit(-1)
	}

	cfg.ListenSpec = "localhost:3000"
	cfg.UI = configs.GetUiConfig()
	return cfg
}


func main() {
	cfg := processFlags()
	daemon.Run(cfg)
}