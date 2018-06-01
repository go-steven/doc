package handler

import (
	"log"
	"os"
)

var logger *log.Logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

func SetLogger(l *log.Logger) {
	logger = l
}
