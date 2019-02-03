package errors

import (
	"log"
	"os"
)

func PrintSystemErr(message string) {
	log.Println("| SYSTEM.ERROR |" + message)
	os.Exit(-1)
}