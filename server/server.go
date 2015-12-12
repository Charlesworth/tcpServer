package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Launching discovery server...")
	go discoveryServer()

	fmt.Println("Launching tcp server...")
	for {
		tcpConnect()
	}
}

func tcpConnect() {
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Close()
			ln.Close()
			break
		}
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))

	}
}

var port = 8081

func discoveryServer() {
	http.HandleFunc("/discover", discoveryHandler)
	http.ListenAndServe(":8080", nil)
}

func discoveryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hit!")
	fmt.Fprint(w, port)
	port++
}
