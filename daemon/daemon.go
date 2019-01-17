package daemon

import (
	"github.com/Ayupov-Ayaz/go-web/db"
	"github.com/Ayupov-Ayaz/go-web/ui"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	ListenSpec string
	Db *db.Config
	UI ui.Config
}

func Run(cfg *Config) error {
	log.Printf("Starting, HTTP on %s\n", cfg.ListenSpec)

	db, err := db.InitDB(cfg.Db)

	if err != nil {
		log.Printf("Error initialization database %v\n", err)
		return err
	}

	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}
	ui.Start(&cfg.UI, db, &l)
	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <- ch
	log.Printf("Got signal: %v, exiting.", s)
}