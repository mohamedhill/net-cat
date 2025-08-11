package netc

import (
	"log"
	"os"
)

func logs(text string) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		
	}
	defer file.Close()

	if _, err := file.WriteString(text); err != nil {
		log.Fatal(err)
	}
}
