package netc
import "net"


func broadcast(message string, excludeConn net.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		if conn != excludeConn {
			conn.Write([]byte(message + "\n"))
		}
	}
}
