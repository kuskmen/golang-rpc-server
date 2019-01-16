package golangrpcserver

import (
	"log"
	"net"
	"net/http"
)

// NewServer creates and starts new rpc server on given port
func NewServer(port string) {
	forever := make(chan int)

	// TODO
	// rpc.Register(queue)
	// rpc.HandleHTTP()

	listener, e := net.Listen("tcp", port)
	log.Println("start to listen:", listener)

	if e != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(listener, nil)
	go Read()

	<-forever
}

// Read Infinite loop where messages are read.
func Read() {

	// TODO
	// Read for incomming messages
}
