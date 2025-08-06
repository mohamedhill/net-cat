package netc

import (
	"bufio"
	"net"
	"strings"
)

func getClientName(conn net.Conn) (string, error) {
	hellomsg := peng()
	conn.Write(hellomsg)

	reader := bufio.NewReader(conn)

	for {
		name, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		name = strings.TrimSpace(name)

		if name == "" {
			conn.Write([]byte("Name cannot be empty.\n"))
			conn.Write([]byte("[ENTER YOUR NAME]:"))
			continue
		}

		return name, nil
	}
}
