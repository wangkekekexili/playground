package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer conn.Close()

			for {
				mType, msg, err := conn.ReadMessage()
				if err != nil {
					log.Println(err)
					return
				}
				err = conn.WriteMessage(mType, msg)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}()
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
