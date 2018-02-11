package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":1097")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = fmt.Fprintln(conn, "hello world")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.(*net.TCPConn).CloseWrite()
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
