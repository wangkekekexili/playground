package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	fmt.Printf("%+v", err())
}

func err() error {
	return errors.New("this is an error")
}
