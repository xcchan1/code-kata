package logger

import (
	"fmt"
	"log"
)

func Error(message string) {
	log.Println(fmt.Sprintf("ERROR: %s", message))
}

func Info(message string) {
	log.Println(fmt.Sprintf("INFO: %s", message))
}
