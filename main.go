package main

import (
	"go-web/daemon"
	"go-web/db"
	"go-web/db/migrations"
	"os"
)

func main() {
	dbConfig := db.GetConfigs()

	// Если main запущен с аргументами, то это миграция
	if len(os.Args) > 1 {
		migrations.Migration(os.Args[1])
		return
	}

	daemonConfig := daemon.Config{
		ListenSpec: "localhost:8000",
		Db: *dbConfig,
	}

	daemon.Run(&daemonConfig)
}