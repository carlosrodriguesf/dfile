package calculator

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

type (
	Calculator interface {
		Calculate(file string) (string, error)
	}

	calculator struct {
		output string
	}
)

func New() Calculator {
	return &calculator{}
}

func (c calculator) Calculate(file string) (string, error) {
	log.Println("sum: ", file)

	f, err := os.Open(file)
	if err != nil {
		log.Printf("error: %v", err)
		return "", err
	}

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Printf("error: %v", err)
		return "", err
	}

	if err = f.Close(); err != nil {
		log.Printf("error: %v", err)
		return "", err
	}

	sum := fmt.Sprintf("%x", h.Sum(nil))
	log.Println("sum result: ", file, " : ", sum)

	return sum, nil
}
