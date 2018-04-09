package qwik

import (
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
)

type FSEvent struct {
	Path      string
	Op        uint32
	Info      os.FileInfo
	Timestamp time.Time
}

type FileWatcher interface {
	Close() error
	RegisterPath(string) error
	DeregisterPath(string) error
}

type fileWatcher struct {
	watcher *fsnotify.Watcher
	handler func(FSEvent)
}

func NewFileWatcher(handler func(FSEvent)) (FileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		return fileWatcher{}, err
	}

	fileWatcher := fileWatcher{
		watcher: watcher,
		handler: handler,
	}

	go func() {
		for {
			select {
			case evt := <-watcher.Events:
				var info os.FileInfo
				if evt.Op&fsnotify.Remove == 0 {
					info, err = os.Lstat(evt.Name)

					if err != nil {
						log.Error().Msgf("error getting fileinfo: %v", err)
					}
				}
				fileWatcher.handler(FSEvent{
					Path:      evt.Name,
					Op:        uint32(evt.Op),
					Timestamp: time.Now(),
					Info:      info,
				})

			case err := <-watcher.Errors:
				log.Error().Msgf("watcher error: %v", err)
			}
		}
	}()

	return fileWatcher, nil
}

func (w fileWatcher) Close() error {
	return w.watcher.Close()
}

func (w fileWatcher) RegisterPath(path string) error {
	return w.watcher.Add(path)
}

func (w fileWatcher) DeregisterPath(path string) error {
	return w.watcher.Remove(path)
}
