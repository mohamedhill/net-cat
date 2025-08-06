package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Client struct {
	Conn net.Conn
	Name string
}

var (
	clients    = make(map[net.Conn]Client)
	clientsMu  sync.Mutex
	messageLog []string      
	logMu      sync.Mutex    
)

func main() {
   adrr := ""
	if len(os.Args) == 1 {
		adrr = ":8989"
	}else if len(os.Args)==2{
		adrr = ":"+os.Args[1]
		
	}else{
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	//args := os.Args[1:]
	ln, err := net.Listen("tcp", adrr)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer ln.Close()
	fmt.Printf("Listening on the port %s",adrr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	name, err := getClientName(conn)
	if err != nil {
		fmt.Println("Invalid name. Disconnecting client.")
		return
	}

	client := Client{Conn: conn, Name: name}


	clientsMu.Lock()
	clients[conn] = client
	clientsMu.Unlock()

	sendHistory(conn)


	joinMsg := fmt.Sprintf("%s has joined our chat...", name)
	broadcast(joinMsg, conn)
	addToHistory(joinMsg)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "User connected:", name)

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected: %v\n", name, err)

			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()

			leaveMsg := fmt.Sprintf("%s has left our chat...", name)
			broadcast(leaveMsg, conn)
			addToHistory(leaveMsg)
			return
		}

		message = strings.TrimSpace(message)
		formatted := fmt.Sprintf("[%s][%s]: %s",
			time.Now().Format("2006-01-02 15:04:05"),
			name,
			message)

		addToHistory(formatted)
		broadcast(formatted, nil)
	}
}


func sendHistory(conn net.Conn) {
	logMu.Lock()
	defer logMu.Unlock()
	for _, msg := range messageLog {
		conn.Write([]byte(msg + "\n"))
	}
}

func addToHistory(msg string) {
	logMu.Lock()
	defer logMu.Unlock()
	messageLog = append(messageLog, msg)
}


func broadcast(message string, excludeConn net.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		if conn != excludeConn {
			conn.Write([]byte(message + "\n"))
		}
	}
}


func peng()[]byte{
hellomsg ,err:= os.ReadFile("peng.txt")
if err != nil{
	return []byte("erorr read peng")
}
return hellomsg
}

func getClientName(conn net.Conn) (string, error) {
	hellomsg := peng()
	conn.Write(hellomsg)
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	name = strings.TrimSpace(name)
	
	return name, nil
}
