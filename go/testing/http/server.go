package http

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "hello")
}

var _ http.HandlerFunc = helloHandler
