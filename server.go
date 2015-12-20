package main

import (
	"bufio"
	"fmt"
	"net"
)

var port = 8081

func main() {
	// fmt.Println("Launching discovery server...")
	// go discoveryServer()

	fmt.Println("Launching tcp server...")
	tcpConnect()
}

func tcpConnect() {
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		readr := bufio.NewReader(conn)
		//writr := bufio.NewWriter(conn)
		//rwer := bufio.NewReadWriter(readr, writr)

		msg, err := readr.ReadString('\n')
		if err != nil {
			conn.Close()
			ln.Close()
			break
		}
		// output message received
		fmt.Print("Message Received:", string(msg))
		// sample process for string received
		//newmessage := strings.ToUpper(msg)
		// send new string back to client

		//fmt.Fprintf(conn, msg+"\n")
		// conn.Write([]byte("hi \n"))
		bufio.NewWriter(conn).Write([]byte("hi \n"))

		// _, err = bufio.NewWriter(conn).WriteString(fmt.Sprintln("hi"))
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}
}

// func discoveryServer() {
// 	serv := portManager.New(8100, 8110)
// 	i, _ := serv.TakePort()
// 	fmt.Println(i)
// 	http.HandleFunc("/discover", discoveryHandler)
// 	http.ListenAndServe(":8080", nil)
// }
//
// func discoveryHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("hit!")
// 	fmt.Fprint(w, port)
// 	port++
// }
