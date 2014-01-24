package main

import (
	"fmt"
	"net"
	"os"
)

type ReaderCloser interface {
	Read(p []byte) (n int, err error)
	Close() error
}

func handleConn(conn ReaderCloser) { // conn net.Conn) {
	fmt.Println("Reading once from connection")
	
	var buf [1024]byte
	n, _ := conn.Read(buf[:])
	fmt.Println("Client sent: ", string(buf[0:n]))
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":15440")
	if err != nil {
		fmt.Println("Couldn't listen: ", err)
		os.Exit(-1)
	}
	for {
		fmt.Println("Waiting for inbound connection")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Couldn't accept: ", err)
			os.Exit(-1)
		}
		go handleConn(conn) 
	}
	fmt.Println("All done")
	
}
