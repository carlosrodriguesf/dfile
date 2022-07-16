package lh

import (
	"io"
	"log"
)

func LogError(err error) error {
	log.Printf("error: %v", err)
	return err
}

func LogClose(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Printf("error: %v", err)
	}
}
