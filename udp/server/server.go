package main

import (
	"net"
)

func handleConnection(conn net.PacketConn, addr net.Addr, buf []byte) {
	buf = []byte("hello Client!\n")
	conn.WriteTo(buf, addr)
}

func main() {
	l, _ := net.ListenPacket("udp", ":8080")
	for {
		buf := make([]byte, 1024)
		n, addr, err := l.ReadFrom(buf)
		if err != nil {
			continue
		}
		go handleConnection(l, addr, buf[:n])
	}
}
