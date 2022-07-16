package watch

import (
	"github.com/carlosrodriguesf/dfile/src/app/path"
	"github.com/carlosrodriguesf/dfile/src/app/sum"
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/carlosrodriguesf/dfile/src/pkg/dbfile"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

type (
	App interface {
		Watch(ctx context.Context) error
	}

	appImpl struct {
		appSum  sum.App
		appPath path.App
	}
)

func New(sum sum.App) App {
	return &appImpl{
		appSum: sum,
	}
}

func (a *appImpl) Watch(ctx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	go a.startWatch(ctx, watcher)

	for _, path := range ctx.DBFile().GetPathKeys() {
		err = a.addPathIntoWatcher(ctx, watcher, path)
		if err != nil {
			log.Printf("error: %v", err)
			return err
		}
	}

	<-make(chan bool)

	return nil
}

func (a *appImpl) onRemove(ctx EventContext) {
	dbFile := ctx.DBFile()
	dbFile.DelFile(ctx.Event.Name)
	err := dbFile.Persist()
	if err != nil {
		log.Fatal(err)
	}
}

func (a *appImpl) onRename(ctx EventContext) {
	ctx.DBFile().DelFile(ctx.Event.Name)
}

func (a *appImpl) onCreate(ctx EventContext) {
	if ctx.IsDir {
		err := a.addPathIntoWatcher(ctx, ctx.Watcher, ctx.Event.Name)
		if err != nil {
			log.Printf("error: %v", err)
		}
		return
	}

	ctx.DBFile().SetFile(ctx.Event.Name, dbfile.FileEntry{
		Ready: false,
	})

	go a.appSum.GenerateFileSum(ctx, ctx.Event.Name, true)
}

func (a *appImpl) startWatch(ctx context.Context, watcher *fsnotify.Watcher) {
	eventMap := map[fsnotify.Op]func(EventContext){
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
				eventMap[event.Op](EventContext{
					Context: ctx,
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

func (a *appImpl) addPathIntoWatcher(ctx context.Context, watcher *fsnotify.Watcher, path string) error {
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
			err = a.addPathIntoWatcher(ctx, watcher, path+"/"+entry.Name())
			if err != nil {
				log.Printf("error: %v", err)
				return err
			}
		}
	}

	return nil
}
