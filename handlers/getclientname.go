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
		booln := isnameexist(name)
		bool2 := validname(name)

		if !booln {
			conn.Write([]byte("this name is exist .\n"))
			conn.Write([]byte("[ENTER YOUR NAME]:"))
			continue
		}
		if !bool2 {
			conn.Write([]byte("this is not valid name.\n"))
			conn.Write([]byte("[ENTER YOUR NAME]:"))
			continue
		}

		if name == "" {
			conn.Write([]byte("Name cannot be empty.\n"))
			conn.Write([]byte("[ENTER YOUR NAME]:"))
			continue
		}

		return name, nil
	}
}

func isnameexist(name string) bool {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	for _, k := range clients {
		if name == k {
			return false
		}
	}
	return true
}

func validname(name string) bool {
	if len(name) == 0 || len(name) > 20 {
		return false
	}
	for _, i := range name {
		if i < 32 || i > 126 {
			return false
		}
	}

	return true
}

func changeClientName(conn net.Conn) (string, error) {
	conn.Write([]byte("[ENTER YOUR NEW NAME]:"))
	reader := bufio.NewReader(conn)

	for {
		name, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		name = strings.TrimSpace(name)
		booln := isnameexist(name)
		bool2 := validname(name)

		if !booln {
			conn.Write([]byte("this name is exist .\n"))
			conn.Write([]byte("[ENTER YOUR NEW NAME]:"))
			continue
		}
		if !bool2 {
			conn.Write([]byte("this is not valid name.\n"))
			conn.Write([]byte("[ENTER YOUR NAME]:"))
			continue
		}

		if name == "" {
			conn.Write([]byte("Name cannot be empty.\n"))
			conn.Write([]byte("[ENTER YOUR NEW NAME]:"))
			continue
		}

		return name, nil
	}
}

func Isvalidmessage(msg string) bool {
	/* if len(msg) < 0 {
		return false
	} */

	for _, i := range msg {
		if i < 32 || i > 126 {
			return false
		}
	}
	return true
}
