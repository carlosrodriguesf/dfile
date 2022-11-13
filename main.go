package main

import (
	"github.com/carlosrodrigues/dfile/cmd"
	"log"
)

func main() {
	if err := cmd.Command().Execute(); err != nil {
		log.Fatal(err)
	}
}
