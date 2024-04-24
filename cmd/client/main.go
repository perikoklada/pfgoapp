package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"goapp/pkg/util"

	"github.com/gorilla/websocket"
)

func main() {

	// check arguments
	if len(os.Args) < 2 {
		log.Fatal("error: missing number of sessions to create.\nusage: client SESSION_NUM")
	}

	var sessionNum uint64
	var err error
	if sessionNum, err = strconv.ParseUint(os.Args[1], 10, 16); err != nil {
		log.Fatal("error: number of sessions to create should be an integer.\nusage: client SESSION_NUM")
	}

	// Register signal handlers for exiting
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)

	url := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/goapp/ws"}
	originHeaderValue := "http://" + url.Host

	for i := sessionNum; i > 0; i-- {

		go func(id uint64) {

			// connect
			c, _, err := websocket.DefaultDialer.Dial(url.String(), http.Header{util.OriginHeaderName: []string{originHeaderValue}})
			if err != nil {
				log.Printf("error connecting to web socket: err: %s, goroutine:%d ", err, id)
			}
			defer c.Close()

			// start reading
			for {
				mt, message, err := c.ReadMessage()
				if err != nil {
					log.Println("error reading:", err)
					return
				}
				log.Printf("goroutine: %d, recv: %s, type: %d", id, message, mt)
			}

		}(i)

		time.Sleep(100 * time.Millisecond)
	}

	<-exitChannel

}
