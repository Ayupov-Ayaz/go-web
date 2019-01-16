package main

import (
	"go-web/daemon"
	"go-web/db"
	"go-web/db/migrations"
	"os"
)

func main() {
	dbConfig := db.GetConfigs()

	if os.Args[1] == "--migrate" {
		migrations.StartMigration()
		return
	} else if os.Args[1] == "--drop-tables" {
		migrations.RollBack()
		return
	}


	daemonConfig := daemon.Config{
		ListenSpec: "localhost:8000",
		Db: *dbConfig,
	}

	daemon.Run(&daemonConfig)
}