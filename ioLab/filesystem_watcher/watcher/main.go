package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

func main() {
	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)

	w.FilterOps(watcher.Create)

	// Only files that match the regular expression during file listings
	// will be watched.
	r := regexp.MustCompile("^.*\\.log$")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add("C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc/go-lab/ioLab/filesystem_watcher/logfiles/hogemoge"); err != nil {
		log.Fatalln(err)
	}

	// Trigger 2 events after watcher started.
	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Create, nil)
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
