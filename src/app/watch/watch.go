package watch

import (
	"errors"
	"github.com/carlosrodriguesf/dfile/src/app/sum"
	"github.com/carlosrodriguesf/dfile/src/repository/file"
	"github.com/carlosrodriguesf/dfile/src/repository/path"
	"github.com/carlosrodriguesf/dfile/src/tool/hlog"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

type (
	Options struct {
		PathRep path.Repository
		FileRep file.Repository
		Sum     sum.App
	}
	App interface {
		Watch() error
	}

	appImpl struct {
		pathRep path.Repository
		fileRep file.Repository
		appSum  sum.App
	}
)

func New(opts Options) App {
	return &appImpl{
		pathRep: opts.PathRep,
		fileRep: opts.FileRep,
		appSum:  opts.Sum,
	}
}

func (a *appImpl) Watch() error {
	allPaths, err := a.pathRep.All()
	if err != nil {
		return hlog.LogError(err)
	}
	if len(allPaths) == 0 {
		return hlog.LogError(errors.New("no paths to watch"))
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	go a.startWatch(watcher)

	for _, currentPath := range allPaths {
		err = a.addPathIntoWatcher(watcher, currentPath.Path)
		if err != nil {
			log.Printf("error: %v", err)
			return err
		}
	}

	<-make(chan bool)

	return nil
}

func (a *appImpl) onRemove(event Event) {
	a.fileRep.Remove(event.Name)
}

func (a *appImpl) onRename(event Event) {
	a.fileRep.Remove(event.Name)
}

func (a *appImpl) onCreate(event Event) {
	if event.IsDir {
		err := a.addPathIntoWatcher(event.Watcher, event.Name)
		if err != nil {
			hlog.Error(err)
		}
		return
	}

	go a.appSum.GenerateFileSum(event.Event.Name)
}

func (a *appImpl) startWatch(watcher *fsnotify.Watcher) {
	eventMap := map[fsnotify.Op]func(Event){
		fsnotify.Remove: a.onRemove,
		fsnotify.Rename: a.onRemove,
		fsnotify.Create: a.onCreate,
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			switch event.Op {
			case fsnotify.Remove, fsnotify.Rename, fsnotify.Create:
				log.Println("watcher event:", event.Op, event.Name)
				info, err := os.Stat(event.Name)
				if err != nil {
					log.Printf("error: %v", err)
					return
				}
				eventMap[event.Op](Event{
					Watcher: watcher,
					Event:   &event,
					IsDir:   info.IsDir(),
				})
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func (a *appImpl) addPathIntoWatcher(watcher *fsnotify.Watcher, path string) error {
	log.Println("watching path:", path)
	err := watcher.Add(path)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			err = a.addPathIntoWatcher(watcher, path+"/"+entry.Name())
			if err != nil {
				log.Printf("error: %v", err)
				return err
			}
		}
	}

	return nil
}
