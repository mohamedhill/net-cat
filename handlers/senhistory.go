package netc
import "net"
// sending history for new joigned clients
func sendHistory(conn net.Conn) {
	logMu.Lock()
	defer logMu.Unlock()
	for _, msg := range messageLog {
		conn.Write([]byte(msg + "\n"))
	}
}