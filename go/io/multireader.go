package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	header := strings.NewReader("<message>")
	body := strings.NewReader("hello world")
	footer := strings.NewReader("</message>")
	mr := io.MultiReader(header, body, footer)
	io.Copy(os.Stdout, mr)
}
