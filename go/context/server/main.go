package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("start")
	defer log.Println("end")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-r.Context().Done():
		log.Println(r.Context().Err())
		http.Error(w, r.Context().Err().Error(), http.StatusInternalServerError)
	}
}
