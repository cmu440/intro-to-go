package main

import (
	"fmt"
	"net"
	"os"
)

type myConn struct {
	conn net.Conn
	prefix string
}

func handleConn(mconn *myConn) {
	fmt.Println("Reading once from connection")
	
	var buf [1024]byte
	n, err := mconn.conn.Read(buf[:])
	if err != nil {
		fmt.Println("Error on read: ", err)
		os.Exit(-1)
	}
	
	fmt.Println(mconn.prefix, ":", string(buf[0:n]))
	mconn.conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":15440")
	if err != nil {
		fmt.Println("Error on listen: ", err)
		os.Exit(-1)
	}
	connNumber := 0
	for {
		fmt.Println("Waiting for a connection via Accept")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error on accept: ", err)
			os.Exit(-1)
		}
		mconn := &myConn{
			conn:conn, 
		        prefix: fmt.Sprintf("%d says", connNumber),
		}
		go handleConn(mconn)
		connNumber++
	}
	fmt.Println("Exiting")
}
