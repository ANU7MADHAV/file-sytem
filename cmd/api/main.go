package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("err", err)
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					fmt.Println("!ok")
					break
				}
				log.Println("event", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file ", event)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					log.Println("ok not")
					return
				}
				log.Println("erro", err)
			}

		}

	}()

	err = watcher.Add("testDir/test.txt")
	if err != nil {
		log.Fatal("err", err)
	}
	<-make(chan struct{})
}
