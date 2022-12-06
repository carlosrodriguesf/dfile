package sumcalc

import (
	"crypto/sha256"
	"fmt"
	"github.com/carlosrodriguesf/dfile/pkg/tool/stacktrace"
	"io"
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
	f, err := os.Open(file)
	if err != nil {
		return "", stacktrace.WrapError(err)
	}

	h := sha256.New()
	if _, err = io.Copy(h, f); err != nil {
		return "", stacktrace.WrapError(err)
	}

	if err = f.Close(); err != nil {
		return "", stacktrace.WrapError(err)
	}

	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum, nil
}
