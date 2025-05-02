package mylogger

import (
	"log"
)

func LogInfo(msg string) {
	log.Printf("[INFO] %s", msg)
}
