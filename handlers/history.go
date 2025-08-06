package netc


func addToHistory(msg string) {
	logMu.Lock()
	defer logMu.Unlock()
	messageLog = append(messageLog, msg)
}