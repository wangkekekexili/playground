package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buf bytes.Buffer
	w := io.MultiWriter(os.Stdout, os.Stderr, &buf)
	fmt.Fprintln(w, "hello world")
	fmt.Println("from buffer:", buf.String())
}
