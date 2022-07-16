package main

import (
	"fmt"
	"github.com/carlosrodriguesf/dfile/src/cmd"
	"log"
	"os"
)

func main() {
	log.SetPrefix(fmt.Sprintf("[%d] ", os.Getppid()))
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
