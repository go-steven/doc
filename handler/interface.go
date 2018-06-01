package handler

import (
	"log"
)

var logger *log.Logger

func SetLogger(l *log.Logger) {
	logger = l
}
