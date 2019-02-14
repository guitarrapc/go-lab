package main

import (
	"log"
	"path/filepath"
	"regexp"

	"github.com/rjeczalik/notify"
)

func main() {

	// Make the channel buffered to ensure no event is dropped. Notify will drop
	// an event if the receiver is not able to keep up the sending pace.
	c := make(chan notify.EventInfo, 1)
	defer close(c)

	// Set up a watchpoint listening on events within current working directory.
	// Dispatch each create and remove events separately to c.
	// notify.Create event will dispatch via both folder and file.
	// notify.FileNotifyChangeXxxxx to restrict event only file.
	dirs := make([]string, 2)
	dirs[0] = "C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc/go-lab/ioLab/filesystem_watcher/logfiles/hogemoge"
	dirs[1] = "C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc/go-lab/ioLab/filesystem_watcher/logfiles/_hogemoge"
	for _, dir := range dirs {
		if err := notify.Watch(dir, c, notify.FileNotifyChangeFileName); err != nil {
			log.Fatal(err)
		}
		defer notify.Stop(c)
	}

	pattern := regexp.MustCompile(".*log")

	// Block until an event is received.
	for {
		select {
		case ei := <-c:
			if pattern.MatchString(filepath.Base(ei.Path())) {
				log.Println("Got event:", ei)

				// file event have many subevents.
				// switch to use appropriate event.
				// match event: https://github.com/rjeczalik/notify/issues/10#issuecomment-66179535
				switch ei.Event() {
				case notify.FileActionAdded:
					log.Println("added!!!!!", filepath.Base(ei.Path()))
				case notify.FileActionModified:
					log.Println("modified!!!!!", filepath.Base(ei.Path()))
				case notify.FileActionRemoved:
					log.Println("removed!!!!!", filepath.Base(ei.Path()))
				case notify.FileActionRenamedNewName:
					log.Println("renamedNewName!!!!!", filepath.Base(ei.Path()))
				case notify.FileActionRenamedOldName:
					log.Println("renamedOldName!!!!!", filepath.Base(ei.Path()))
				}
			}
		}
	}
}

func isFileActionAdded(e notify.Event) bool {
	return e == notify.FileActionAdded
}
