package main

import (
	"log"
	"os"
	"path/filepath"

	"golang.org/x/exp/winfsnotify"
)

func main() {
	watcher, err := winfsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		return
	}
	watcher.Close()
	path := "C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc/go-lab/ioLab/filesystem_watcher/winfsnotify/logs"
	err = watcher.AddWatch(path, winfsnotify.FS_CREATE)
	if err != nil {
		log.Fatalf("Watcher.Watch() failed: %s", err)
	}

	go func() {
		for err := range watcher.Error {
			log.Fatalf("error received: %s", err)
		}
	}()

	go func() {
		for {
			select {
			case e := <-watcher.Event:
				log.Printf("create event detected. %s, %s", e.Name, e.String())
			}
		}
	}()

	f, err := os.Create(filepath.Join(path, "hogemoge.log"))
	f.Close()

	end := make(chan string, 1)
	<-end
}
