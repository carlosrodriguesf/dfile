package main

import (
	"encoding/json"
	"fmt"
	"github.com/carlosrodriguesf/dfile/src/repository"
	"github.com/carlosrodriguesf/dfile/src/tool/dbm"
	"github.com/carlosrodriguesf/dfile/src/tool/lh"
	"log"
)

func main() {
	db, err := dbm.Open("dfile.db")
	if err != nil {
		log.Fatal(err)
	}
	defer lh.LogClose(db)

	pathRepository := repository.Path(db)
	//err = pathRepository.Save(model.Path{
	//	Path:             "/test",
	//	Enabled:          true,
	//	IgnoreFolders:    []string{".debris"},
	//	AcceptExtensions: []string{"jpg", "png"},
	//})
	path, err := pathRepository.Get("/test")
	if err != nil {
		log.Fatal(err)
	}

	dt, _ := json.Marshal(path)
	fmt.Println(string(dt))

	//log.SetPrefix(fmt.Sprintf("[%d] ", os.Getppid()))
	//if err := cmd.Run(); err != nil {
	//	log.Fatal(err)
	//}
}
