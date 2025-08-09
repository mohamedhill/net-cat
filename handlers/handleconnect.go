package netc

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)


var (
	clients    = make(map[net.Conn]string)
	clientsMu  sync.Mutex
	messageLog []string
	logMu      sync.Mutex
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	name, err := getClientName(conn)
	if err != nil {
		fmt.Println("Invalid name. Disconnecting client.")
		return
	}

	clientsMu.Lock()
	if len(clients) >= 10 {
		clientsMu.Unlock()
		conn.Write([]byte("Server full. Try again later.\n"))
		return
	}

	

	clients[conn] = name 

	clientsMu.Unlock()

	sendHistory(conn)
	//propmt()

	joinMsg := fmt.Sprintf("✅ %s has joined our chat...", name)
	broadcast2(joinMsg, conn)


	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "✅ User connected:", name)

	reader := bufio.NewReader(conn)

	for {
		propmt()
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("🔴 Client %s disconnected: %v\n", name, err)

			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()

			leaveMsg := fmt.Sprintf("🔴 %s has left our chat...", name)
			broadcast2(leaveMsg, conn)
			//propmt()
			return
		}

		message = strings.TrimSpace(message)
		formatted := fmt.Sprintf("[%s][%s]: %s",
			time.Now().Format("2006-01-02 15:04:05"),
			name,
			message)

		addToHistory(formatted)
		broadcast(formatted, conn,name)
		//propmt()
		/* broadcast2(formatted1,conn) */
	}
}
