package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	pr, pw := io.Pipe()
	go func() {
		fmt.Fprintln(pw, "hello world")
		pw.Close()
	}()
	io.Copy(os.Stdout, pr)
}
