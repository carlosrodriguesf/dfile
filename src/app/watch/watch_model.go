package watch

import (
	"github.com/carlosrodriguesf/dfile/src/tool/context"
	"github.com/fsnotify/fsnotify"
)

type EventContext struct {
	context.Context
	Watcher *fsnotify.Watcher
	Event   *fsnotify.Event
	IsDir   bool
}
