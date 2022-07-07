// Header to update later
package main

import (
	"fmt"
	"net"
	"bufio"
)

/* type Client
	name - yo name you need to set
	receiving - what we are sending to the server
	sending - incoming messages from the server	
	reader - reader connected to the server channel
	writer - writer connected to the server channel

	TODO: 
	id - numeric ID assigned by server
	name - based on some credentials??? 
*/
type Client struct {
	name		string
	receiving	chan string
	sending		chan string
	reader		*bufio.Reader
	writer		*bufio.Writer
}

/* func NewClient 
	Creates a base client with no connection
*/
func newclient() *Client {
	// Ask user to input a name
	fmt.Printf("Please enter your name: ")
	fmt.Scanf("%s", &(client.name))

	// Setup new client
	client := &Client {
		receiving: make(chan string),
		sending: make(chan string),
	}

	return client
}

/* func NewConnection 
	Connects the reader and writer to the server connection
*/
func (client *Client) newconnection(connection net.Conn) {
	client.reader := bufio.NewReader(connection)
	client.writer:= bufio.NewWriter(connection)

	client.listen()
}

/* func Listen
	Captures clients input and output on screen
*/
func (client *Client) listen() {
	go client.read()
	go client.write()
}

/* func Read
	Reads client input into the reader
*/
func (client *Client) read() {
	for {
		line, _ := client.reader.ReadString('\n')
		client.receiving <-line
	}
}

/* func Write
	Sents the clients output
*/
func (client *Client) write() {
	for message := range client.sending {
		client.writer.WriteString(message)
		client.writer.Flush()
	}
}