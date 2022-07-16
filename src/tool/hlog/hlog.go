package hlog

import (
	"io"
	"log"
)

func LogError(err error) error {
	Error(err)
	return err
}

func LogClose(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Printf("error: %v", err)
	}
}

func Error(err error) {
	log.Printf("error: %v", err)
}
