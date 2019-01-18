package daemon

import (
	"fmt"
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

func Run(cfg *Config) {
	errorMessage := "DB error: in daemon/daemon.go (%s) \n %s"

	db, err := db.InitDB(cfg.Db)
	if err != nil {
		log.Printf(errorMessage, "InitDB", err.Error())
	}

	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Printf(errorMessage, "net.Listen", err)
	}
	fmt.Printf("Starting server on %s", cfg.ListenSpec )
	ui.Start(&cfg.UI, db, &l)
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <- ch
	log.Printf("Got signal: %v, exiting.", s)
}