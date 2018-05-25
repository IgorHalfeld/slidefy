package helpers

import (
	"log"
	"os"
)

// ErrorHandler abstract error messages
func ErrorHandler(err error, message string) {
	if err != nil {
		log.Fatalf(message, ": %v", err)
		os.Exit(1)
	}
}
