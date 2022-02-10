package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	fmt.Println(arguments)
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	readonly := false

	if len(arguments) == 3 && arguments[2] == "-readonly" {
		fmt.Println("You are a readonly user")
		readonly = true
	}

	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	// c.Write([]byte("New Client with name\n"))

	// if len(arguments) == 3 {
	// 	c.Write([]byte("ADMIN\n"))

	// 	message, err := bufio.NewReader(c).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(message)
	// 	// c.Close()
	// 	return
	// }

	for {
		if readonly {
			message, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message)
			continue
		}
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		
		if strings.ToUpper(strings.TrimSpace(string(text))) == "EXIT" {
			fmt.Println("TCP client exiting...")
			fmt.Fprintf(c, "%v has disconnected\n", c.LocalAddr())
			return
		}

		fmt.Fprintf(c, text+"\n")
	}
}
