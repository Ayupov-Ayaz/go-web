package errors

import (
	"log"
	"os"
)

func PrintSystemErr(message string) {
	log.Println("| SYSTEM.ERROR |" + message)
	os.Exit(-1)
}

type Error struct {
	s string
}

func NewError(message string) *Error {
	return &Error{
		s: message,
	}
}

func (e *Error)Error() string {
	return e.s
}