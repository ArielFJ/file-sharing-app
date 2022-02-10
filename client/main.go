package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var args = os.Args
	if len(args) == 1 {
		fmt.Printf("MSG should be passed as argument")
		return
	}

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		printErr(err)
	}
	defer conn.Close()
	msg := strings.Join(args[1:], " ")
	text := fmt.Sprintf("%v says: %v\n", conn.LocalAddr(), msg)
	conn.Write([]byte(text))
}

func printErr(err error) {
	fmt.Printf("Error: %v", err)
}
