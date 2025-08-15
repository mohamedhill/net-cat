package netc

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)
// Declaring global variables
var (
	clients    = make(map[net.Conn]string)
	clientsMu  sync.Mutex
	messageLog []string
	logMu      sync.Mutex
)
// the main function to handle connections(Name,Limit,prompt,broadcast, connect and disconnect...)
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	disconnect := false

	name, err := getClientName(conn)
	if err != nil {
		fmt.Println("Invalid name. Disconnecting client.")
		return
	}

	clientsMu.Lock()
	if len(clients) >= 10 {
		clientsMu.Unlock()
		_, err := conn.Write([]byte("Server full. Try again later.\n"))
		if err != nil {
			fmt.Println("erorr writing to the client", err)
		}
		return
	}

	clients[conn] = name

	clientsMu.Unlock()

	sendHistory(conn)

	joinMsg := fmt.Sprintf("✅ %s has joined our chat...", name)
	broadcast(joinMsg, conn, disconnect)
	disconnect = true
	logs(joinMsg + "\n")

	reader := bufio.NewReader(conn)

	flag := true
	for {
		if flag {
			propmt()
			disconnect = false
		}
		message, err := reader.ReadString('\n')
		if err != nil {
			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()

			leaveMsg := fmt.Sprintf("🔴 %s has left our chat...", name)
			broadcast(leaveMsg, conn, disconnect)
			propmt()
			disconnect = true
			logs(leaveMsg + "\n")
			flag = true
			return
		}
		message = strings.TrimSpace(message)

		if message == "/name" {
			currentname := name
			newname, err := changeClientName(conn)
			if err != nil {
				fmt.Println("Invalid name. Disconnecting client.")
				return
			}
			clientsMu.Lock()
			clients[conn] = newname
			clientsMu.Unlock()
			name = newname
			changenameMsg := fmt.Sprintf("📝%s has change there name to: %s", currentname, newname)
			broadcast(changenameMsg, conn, disconnect)
			propmt()
			disconnect = true
			logs(changenameMsg + "\n")
			flag = false
			continue

		} else if message == "/members" {
			conn.Write([]byte("💬➡️ the chat members is\n"))
			clientsMu.Lock()
			for _, j := range clients {
				conn.Write([]byte("👤:" + j + "\n"))
			}
			clientsMu.Unlock()
			flag = true
		}

		if message == "" || !Isvalidmessage(message) {
			flag = false
			clientsMu.Lock()
			clientName, ok := clients[conn]
			clientsMu.Unlock()
			if !ok {
				return
			}

			formatted1 := fmt.Sprintf("[%s][%s]:",
				time.Now().Format("2006-01-02 15:04:05"),
				clientName)
			conn.Write([]byte(formatted1))
			continue
		}

		formatted := fmt.Sprintf("[%s][%s]: %s",
			time.Now().Format("2006-01-02 15:04:05"),
			name,
			message)

		addToHistory(formatted)
		broadcast(formatted, conn, disconnect)
		disconnect = false
		logs(formatted + "\n")
		flag = true

	}
}
