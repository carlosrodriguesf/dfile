package watch

import (
	"github.com/fsnotify/fsnotify"
)

type Event struct {
	*fsnotify.Event
	Watcher *fsnotify.Watcher
	IsDir   bool
}
