package netc

import (
	"fmt"
	"net"
	"time"
)

func broadcast(message string, excludeConn net.Conn, name string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		if conn != excludeConn {
			conn.Write([]byte("\n" + message + "\n"))
		
			/* formatted1 := fmt.Sprintf("[%s][%s]:",
				time.Now().Format("2006-01-02 15:04:05"),
				clients[conn])
			conn.Write([]byte("\n" + formatted1))
			fmt.Println(formatted1)
 */
		}
	}
}

func broadcast2(message string, excludeConn net.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		if conn != excludeConn {
			conn.Write([]byte("\n"+message+"\n"))
		}
	}
}

func propmt() {
	for conn := range clients {

		formatted1 := fmt.Sprintf("[%s][%s]:",
			time.Now().Format("2006-01-02 15:04:05"),
			clients[conn])
		conn.Write([]byte(formatted1))
		//fmt.Println(formatted1)

	}
}
