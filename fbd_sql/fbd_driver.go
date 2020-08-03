package main

import (
	"fmt"

	"github.com/nakagami/firebirdsql"
)

func main() {
	dsn := "user:password@servername/foo/bar.fdb"
	events := []string{"my_event", "order_created"}
	fbEvent, _ := firebirdsql.NewFBEvent(dsn)
	defer fbEvent.Close()
	sbr, _ := fbEvent.Subscribe(events, func(event firebirdsql.Event) { //or use SubscribeChan
		fmt.Printf("event: %s, count: %d, id: %d, remote id:%d \n", event.Name, event.Count, event.ID, event.RemoteID)
	})
	defer sbr.Unsubscribe()
	go func() {
		fbEvent.PostEvent(events[0])
		fbEvent.PostEvent(events[1])
	}()
	<-make(chan struct{}) //wait
}
