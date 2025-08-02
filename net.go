package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	Conn net.Conn
	Name string
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error", err)
	}
	defer ln.Close()
	fmt.Println("the server is running")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error", err)
			continue
		}
		go handleconnection(conn)
	}
}

func handleconnection(conn net.Conn) {
	defer conn.Close()

	name, err := getClientName(conn)
	if err != nil {
		fmt.Println("Invalid name, disconnecting client.")
		return
	}

	fmt.Println(time.Now(), "User name is:", name)

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected: %v\n", name, err)
			return
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue 
		}

		formatted := fmt.Sprintf("[%s][%s]: %s",
			time.Now().Format("2006-01-02 15:04:05"),
			name, message,
		)

		fmt.Println(formatted)

		
		_, err = conn.Write([]byte(formatted + "\n"))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
	}
}

const Name = "[ENTER YOUR NAME]:"
func getClientName(conn net.Conn) (string, error) {
	conn.Write([]byte(Name))
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	name = strings.TrimSpace(name)
	if name == "" {
		conn.Write([]byte("Name cannot be empty. Goodbye!\n"))
		conn.Close()
		return "", fmt.Errorf("empty name")
	}

	return name, nil
}
