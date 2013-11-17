package onChange

import (
	"github.com/howeyc/fsnotify"
	. "github.com/azer/debug"
)

func OnChange (filename string, onChange func()) {
	Debug("Watching %s", filename)

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		Debug("Failed to setup fsnotify")
		return
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				Debug("%s has been updated. Event: %s", filename, ev)
				onChange()
			case erv := <-watcher.Error:
				Debug("Failed to monitor changes. Error: %s", erv)
			}
		}
	}()

	err = watcher.Watch(filename)

	if err != nil {
		Debug("Failed to monitor changes on %s", filename)
	}
}
