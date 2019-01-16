package daemon

import (
	"fmt"
	"go-web/db"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	ListenSpec string
	Db db.Config
}

func Run(cfg *Config) error {
	log.Printf("Starting, HTTP on %s\n", cfg.ListenSpec)

	db, err := db.DbConnection(&cfg.Db)
	if err != nil {
		log.Printf("Error initialization database %v\n", err)
		return err
	}

	// TODO:
	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}

	fmt.Println(db, l)
	waitForSignal()
	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <- ch
	log.Printf("Got signal: %v, exiting.", s)
}