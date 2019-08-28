package main

import (
	"net"
	"fmt"
	"time"
	"log"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	
	for i := 0; i < 10000; i++ { 
		start := time.Now()
		
		fmt.Fprintf(conn, "hello server!\n")
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		
		elapsed := time.Since(start)
		fmt.Printf("%s\n", elapsed)
			
		log.Print("Message from server: " + string(buffer))
	}
	conn.Close()
}
