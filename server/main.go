package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		printErr(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			printErr(err)
			continue
		}

		go handleConn(conn)
	}
}

func printErr(err error) {
	fmt.Printf("Error: %v", err)
}

func handleConn(c net.Conn) {
	for {
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			// printErr(err)
			c.Close()
			return
		}

		if strings.ToUpper(msg) == "EXIT" {
			c.Close()
			return
		}

		fmt.Println(msg)
	}
}