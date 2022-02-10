package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type message struct {
	Filename string `json:filename`
	Data     []byte `json:data`
}

func (m *message) String() string {
	return fmt.Sprintf("{\n filename: %v,\n dataLength: %v \n}\n", m.Filename, len(m.Data))
}

const delimiter byte = 254 // ■

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
		// msg, err := bufio.NewReader(c).ReadBytes(delimiter)
		msg, err := ioutil.ReadAll(c)
		if err != nil {
			// printErr(err)
			c.Close()
			return
		}

		var myMessage message
		err = json.Unmarshal(msg, &myMessage)
		if err != nil {
			printErr(fmt.Errorf("JSON: %v", err))
			return
		}

		// if strings.ToUpper(msg) == "EXIT" {
		// 	c.Close()
		// 	return
		// }
		
		// toWrite := msg[: len(msg) - 1]
		// toWrite = append(toWrite, '\n')
		// fmt.Println(toWrite)
		fmt.Println(myMessage)
		os.WriteFile("../file/"+myMessage.Filename, myMessage.Data, 0644)
		// return
	}
}