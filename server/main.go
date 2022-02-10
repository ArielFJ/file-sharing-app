package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var conns = make(map[net.Conn]bool)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		conns[c] = true
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			// if err == io.EOF {
			// 	return
			// }
			// fmt.Println("err:", err)
			delete(conns, c)
			c.Close()
			return
		}
		// if strings.TrimSpace(string(netData)) == "STOP" {
		//         fmt.Println("Exiting TCP server!")
		//         return
		// }
		if strings.TrimSpace(netData) == "ADMIN" {
			fmt.Println("Es el admin")
			c.Write([]byte("EL ADMINNNNN\n"))
			return
		}

		fmt.Printf("-> %v: %v", c.RemoteAddr(), string(netData))
		t := time.Now()
		for conn := range conns {
			if conn != c {
				msg := fmt.Sprintf("%v, %v\n", t.Format(time.Kitchen), netData)
				conn.Write([]byte(msg))
			}
		}
	}
}
