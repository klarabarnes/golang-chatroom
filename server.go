// Header to update later
package main

import (
	"fmt"
	"net"
	"bufio"
)

/* type Server 
	clients - array of all clients in the server
	joiner - connection channel that lets people join the server
	sending - data being sent from the server to the cleints
	receiving - data coming into the server from the clients
*/
type Server struct {
	clients		[]*Client 
	joiner		chan net.Conn
	sending		chan string
	receiving	chan string
}

/* func Broadcast
	Sends a message to all clients in the server
*/
func (server *Server) broadcast(message string) {
	for _, client := range server.cleints {
		client.sending <- message
	}
}

/* func Join
	Accepts new clinets to the server 
*/
func (server *Server) join(connection net.Conn, new Client) {
	new.setConnection(connection)
	server.clients = append(server.clients, new)
	go func () {
		for {
			server.receiving <- <-client.receiving
		}
	}()
}

/* func Listen
	Listens for data from the clients to process 
*/
func (server *Server) listen() {
	go func() {
		for {
			select {
			case message := <- server.receiving:
				server.broadcast(message)
			case connection := <- server.joiner:
				server.join(connection)
			}
		}
	}()
}

/* func NewServer
	Initializes the new server
*/
func newserver() *Server {
	server := &Server {
		clients: make([]*Client,0),
		joiner:	make(chan net.Conn),
		receiving: make(chan string),
		outgoing: make(chan string),
	}

	// Initialize handler and return
	server.listen()
	return server
}