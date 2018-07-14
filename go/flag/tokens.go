package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type tokens []string

var _ flag.Value = &tokens{}

func (t *tokens) String() string {
	return strings.Join([]string(*t), ",")
}

func (t *tokens) Set(s string) error {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return errors.New("no tokens specified")
	}
	for _, token := range strings.Split(s, ",") {
		*t = append(*t, token)
	}
	return nil
}

func main() {
	var t tokens
	flag.Var(&t, "tokens", "a list of comma separated tokens")
	flag.Parse()
	fmt.Println(t)
}
