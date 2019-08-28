package main

import (
	"net"
)

func handleConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		conn.Write([]byte("hello client!\n"))
	}
}

func main() {
    l, _ := net.Listen("tcp4", ":8080")
    
    for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConnection(c)
    }
}
