package logger

import (
	"io"
	"os"
)

type logger struct {
	file    *os.File
	verbose bool
}

func New(outputFile string, verbose bool) (io.WriteCloser, error) {
	file, err := os.Create(outputFile)
	if err != nil {
		return nil, err
	}

	return &logger{
		file:    file,
		verbose: verbose,
	}, nil
}

func (c logger) Write(p []byte) (n int, err error) {
	if c.verbose {
		n, err = os.Stdout.Write(p)
		if err != nil {
			return n, err
		}
	}
	return c.file.Write(p)
}

func (c logger) Close() error {
	return c.file.Close()
}
