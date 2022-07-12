package main

import (
	"github.com/carlosrodriguesf/dfile/src/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
