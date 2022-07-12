package path

import (
	"github.com/carlosrodriguesf/dfile/src/pkg/context"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

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

func (a *appImpl) startWatch(ctx context.Context, watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			switch event.Op {
			case fsnotify.Create, fsnotify.Remove:
				log.Println("watcher event:", event.Op, event.Name)

				if event.Op == fsnotify.Create {
					info, err := os.Stat(event.Name)
					if err != nil {
						log.Printf("error: %v", err)
						continue
					}

					if info.IsDir() {
						err = a.addPathIntoWatcher(ctx, watcher, event.Name)
						if err != nil {
							log.Printf("error: %v", err)
						}
						continue
					}
				}

				err := a.Sync(ctx)
				if err != nil {
					log.Printf("error: %v", err)
					continue
				}

				err = a.sum.Generate(ctx)
				if err != nil {
					log.Printf("error: %v", err)
				}
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
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	log.Println("watching path:", path)
	err = watcher.Add(path)
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
