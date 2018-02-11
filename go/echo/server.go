package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1097")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go echo(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()
	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Println(err)
	}
}
