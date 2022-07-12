package logger

import (
	"io"
	"os"
	"strings"
)

type logger struct {
	file    *os.File
	verbose bool
}

func New(outputFile string, verbose bool) (io.WriteCloser, error) {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0700)
	if err != nil {
		if err == os.ErrNotExist || strings.HasSuffix(err.Error(), "no such file or directory") {
			file, err = os.Create(outputFile)
		} else {
			return nil, err
		}
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
