package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	netc "netc/handlers"
)

func main() {
	address := ":8989" // default port
	if len(os.Args) == 2 {
		newadress, err := strconv.Atoi(os.Args[1])
		if err != nil || newadress < 1024 || newadress > 65000 {	//checking the validity of the port 1024>=port<=65000
			fmt.Println("check the validity of the port")
			return
		}
		address = ":" + strconv.Itoa(newadress)
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
