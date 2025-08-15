package netc
// a function to store the history messages
func addToHistory(msg string) {
	logMu.Lock()
	defer logMu.Unlock()
	messageLog = append(messageLog, msg)
}
