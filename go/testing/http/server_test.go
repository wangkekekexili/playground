package http

import (
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	helloHandler(recorder, nil)
	got := recorder.Body.String()
	if got != "hello" {
		t.Fatalf("got %v", got)
	}
}
