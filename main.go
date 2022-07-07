// Comments to update later
package main

import (
	"fmt"
	"net"
	"bufio"
)


/*
 * MAIN
 */
func main() {

	// Create a Client

	// Creater a Server

	server := newserver()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Printf("Unable to start server: %s", err.Error())
	}

	defer listener.Close()
	fmt.Printf("Server started on :8888")

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %s", err.Error())
			continue
		}

		server.joiner <- connection
	}
}