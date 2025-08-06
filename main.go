package main

import (
	"fmt"
	"net"
	"os"
	"netc/handlers"
	
)

func main() {
	address := ":8989"
	if len(os.Args) == 2 {
		address = ":" + os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	ln, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()
	fmt.Printf("Listening on port %s\n", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go netc.HandleConnection(conn)
	}
}
