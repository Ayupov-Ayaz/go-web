package main

import (
	"go-web/daemon"
	"go-web/db"
)

func main() {
	dbConfig := db.Config{
		Driver: "mysql",
		User: "tommy",
		Password: "43",
		Host: "127.0.0.1",
		Database: "go_lang",
		Port: 3306,
	}

	daemonConfig := daemon.Config{
		ListenSpec: "localhost:8000",
		Db: dbConfig,
	}

	daemon.Run(&daemonConfig)
}