package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 1024)
	_, err := c.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	s := fmt.Sprintf("+PONG\r\n")
	c.Write([]byte(s))
}

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error to listen 8080")
		os.Exit(1)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
