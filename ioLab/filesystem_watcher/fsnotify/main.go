package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer watcher.Close()
	path := "C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc/go-lab/ioLab/filesystem_watcher/fsnotify/logs"
	err = watcher.Add(path)
	if err != nil {
		log.Fatalf("Watcher.Watch() failed: %s", err)
	}

	go func() {
		for err := range watcher.Errors {
			log.Fatalf("error received: %s", err)
		}
	}()

	go func() {
		for {
			select {
			case e := <-watcher.Events:
				if e.Op&fsnotify.Create == fsnotify.Create {
					log.Printf("create event detected. %s, %s", e.Name, e.String())
				} else if e.Op&fsnotify.Rename == fsnotify.Rename {
					log.Printf("rename event detected. %s, %s", e.Name, e.String())
				} else if e.Op&fsnotify.Remove == fsnotify.Remove {
					log.Printf("remove event detected. %s, %s", e.Name, e.String())
				} else if e.Op&fsnotify.Write == fsnotify.Write {
					log.Printf("write event detected. %s, %s", e.Name, e.String())
				}
			}
		}
	}()

	f, err := os.Create(filepath.Join(path, "hogemoge.log"))
	f.Close()

	end := make(chan string, 1)
	<-end
}
