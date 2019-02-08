package main

import (
	"log"

	"github.com/rjeczalik/notify"
)

func main() {

	// Make the channel buffered to ensure no event is dropped. Notify will drop
	// an event if the receiver is not able to keep up the sending pace.
	c := make(chan notify.EventInfo, 1)

	// Set up a watchpoint listening on events within current working directory.
	// Dispatch each create and remove events separately to c.
	if err := notify.Watch("C:/Users/ikiru.yoshizaki/go/src/github.com/guitarrapc/go-lab/ioLab/filesystem_watcher/logfiles/hogemoge", c, notify.Create); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	// Block until an event is received.
	for {
		select {
		case ei := <-c:
			log.Println("Got event:", ei)
		}
	}
}
