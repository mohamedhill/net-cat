package netc

import (
	"fmt"
	"net"
	"time"
)

func broadcast(message string, excludeConn net.Conn, flg bool) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		if conn != excludeConn {
			if flg {
				conn.Write([]byte(message))
			}

			if !flg {
				conn.Write([]byte("\n" + message + "\n"))
			}
		}
	}
}

func propmt() {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	for conn := range clients {
		clientName, ok := clients[conn]
		if !ok {
			continue
		}

		formatted1 := fmt.Sprintf("[%s][%s]:",
			time.Now().Format("2006-01-02 15:04:05"),
			clientName)
		_, err := conn.Write([]byte(formatted1))
		if err != nil {
			fmt.Println("error print the propmt", err)
		}

	}
}
