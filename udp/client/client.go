package main

import (
	"log"
	"net"
	"fmt"
	"time"
)

func main() {
	address, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	conn, _ := net.DialUDP("udp", nil, address)
	
	for i := 0; i < 10000; i++ {
		start := time.Now()
		
		fmt.Fprintf(conn, "hello server!\n")
		buffer := make([]byte, 1024)
		conn.ReadFromUDP(buffer)

		elapsed := time.Since(start)
		fmt.Printf("%s\n", elapsed)		
		
		log.Print("Message from server: " + string(buffer) + "\n")
	}
	conn.Close()
}
