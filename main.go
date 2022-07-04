// Comments to update later
package main

import (
	"fmt"
	"net"
	"bufio"
)


/*
 * STRUCTS
 */
type Client struct {
	name		string // yo name
	incoming	chan string // incoming strings from the server
	outgoing 	chan string // what you are sending to the server
	reader		*bufio.Reader // reader connected to server
	writer		*bufio.Writer // writer connected to server
}

type Server struct {
	clients		[]*Client // all the clients in the server
	joiner		chan net.Conn // connection channel to let people join the server
	incoming	chan string // data that is coming into the server
	outgoing	chan string // what the server is sending to the clients
}

/*
 * CLIENT FUNCTIONS
 * newclient - creates a new client to add to server
 * listen - captures the clients input and output to update screen 
 * read - reads in client input
 * write - sends out client input 
 */
func newclient(connection net.Conn) *Client {
	// Setting up reader and writer)
	this_reader := bufio.NewReader(connection)
	this_writer := bufio.NewWriter(connection)
	
	// Setup new client
	client := &Client {
		name: "anon", 
		incoming: make(chan string),
		outgoing: make(chan string), 
		reader: this_reader, 
		writer: this_writer,
	}

	// Initialize handler and return
	client.listen()
	return client
}

func (client *Client) listen() {
	go client.read()
	go client.write()
}

func (client *Client) read() {
	for {
		line, _ := client.reader.ReadString('\n')
		client.incoming <- line
	}
}

func (client *Client) write() {
	for data := range client.outgoing {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}

/*
 * SERVER FUNCTIONS
 * broadcast - prints out data to all the clients
 * join - creates new clients who join the server
 * listen - listens for incoming data (parses as messages or command) or connection with join
 * newserver - initializes the server and it's variables
 */
func (server *Server) broadcast(data string) {
	for _, client := range server.clients {
		client.outgoing <- data
	}
}

func (server *Server) join(connection net.Conn) {
	client := newclient(connection)
	server.clients = append(server.clients, client)
	go func (){
		for {
			server.incoming <- <-client.incoming 
		}
	}()
}

 func (server *Server) listen() {
	go func() {
		for {
			select {
			case data := <-server.incoming:
				server.broadcast(data)
			case conn := <-server.joiner:
				server.join(conn)
			}
		}
	} ()
}

func newserver() *Server {
	
	// Setup new server
	server := &Server {
		clients: make([]*Client, 0),
		joiner: make(chan net.Conn), 
		incoming: make(chan string),
		outgoing: make(chan string),
	}

	// Intialize handler and return
	server.listen()
	return server
}


/*
 * MAIN
 */
func main() {

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