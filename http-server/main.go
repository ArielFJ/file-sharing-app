package main

import (
	"bufio"
	"encoding/json"
	"file-sharing-app/http-server/models"
	"fmt"
	"log"
	"net"
	"net/http"
)

const tcpPort = ":8000"

func main() {
	port := ":8001"

	http.HandleFunc("/", greet)
	http.HandleFunc("/data", getData("data", []models.ChannelData{}))
	http.HandleFunc("/clients", getData("clients", []string{}))

	fmt.Println("Listening on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

// func getData(w http.ResponseWriter, r *http.Request) {
// 	// Create connection with TCP
// 	conn, err := net.Dial("tcp", tcpPort)
// 	if err != nil {
// 		sendErr(w, "DIAL", err)
// 		return
// 	}

// 	// Request channel data to TCP server
// 	req := models.NewRequest("data", "")
// 	reqBytes, err := json.Marshal(req)
// 	if err != nil {
// 		sendErr(w, "REQ", err)
// 		return
// 	}
// 	conn.Write(append(reqBytes, '\n'))

// 	// Read response data from TCP server
// 	netData, err := bufio.NewReader(conn).ReadString('\n')
// 	if err != nil {
// 		sendErr(w, "READ", err)
// 		return
// 	}

// 	var res []models.ChannelData
// 	err = json.Unmarshal([]byte(netData), &res)
// 	if err != nil {
// 		sendErr(w, "RES", err)
// 		return
// 	}

// 	// Send the received data to HTTP client
// 	bytes, err := json.Marshal(res)
// 	if err != nil {
// 		sendErr(w, "JSON", err)
// 		return
// 	}

// 	sendJsonSuccessfulResponse(w, bytes)
// }

func getData(endpoint string, object interface{}) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// Create connection with TCP
		conn, err := net.Dial("tcp", tcpPort)
		if err != nil {
			sendErr(w, "DIAL", err)
			return
		}
		defer conn.Close()

		// Request data to TCP server
		req := models.NewRequest(endpoint, "")
		reqBytes, err := json.Marshal(req)
		if err != nil {
			sendErr(w, "REQ", err)
			return
		}
		conn.Write(append(reqBytes, '\n'))

		// Read response data from TCP server
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			sendErr(w, "READ", err)
			return
		}

		err = json.Unmarshal([]byte(netData), &object)
		if err != nil {
			sendErr(w, "RES", err)
			return
		}

		// Send the received data to HTTP client
		bytes, err := json.Marshal(object)
		if err != nil {
			sendErr(w, "JSON", err)
			return
		}

		sendJsonSuccessfulResponse(w, bytes)
	}
}

func sendErr(w http.ResponseWriter, prefix string, err error) {
	fmt.Fprint(w, "Error "+prefix+": "+err.Error())
}

func sendJsonSuccessfulResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
