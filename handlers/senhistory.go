package netc
import "net"

func sendHistory(conn net.Conn) {
	logMu.Lock()
	defer logMu.Unlock()
	for _, msg := range messageLog {
		conn.Write([]byte(msg + "\n"))
	}
}