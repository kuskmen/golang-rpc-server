package yamqserver

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	yamq "github.com/kuskmen/yamq"
)

// NewServer creates and starts new rpc server on given port
func NewServer(port string) {
	forever := make(chan int)

	warehouse := new(yamq.Warehouse)

	err := rpc.Register(warehouse)
	if err != nil {
		log.Fatal("Error registering warehouse.", err)
	}

	rpc.HandleHTTP()

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
