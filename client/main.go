package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type message struct {
	Filename string `json:filename`
	Data     []byte `json:data`
}

const delimiter byte = 254 // ■

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

	path := strings.Join(args[1:], " ")
	fileContent, err := os.ReadFile(path)
	if err != nil {
		printErr(err)
		return
	}
	// fileContent = append(fileContent, delimiter)
	message := message{
		Filename: filepath.Base(path),
		Data: fileContent,
	}
	fmt.Println(fileContent)
	// fmt.Println(string(fileContent))
	// text := fmt.Sprintf("%v says: %v\n", conn.LocalAddr(), msg)
	jsonString, err := json.Marshal(message)
	if err != nil {
		printErr(err)
	}
	conn.Write([]byte(jsonString))
}

func printErr(err error) {
	fmt.Printf("Error: %v", err)
}
