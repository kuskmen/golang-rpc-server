package yamqclient

import (
	"log"
	"net/rpc"
)

// Client Generic interface of message pattern client.
type Client interface {
	NewClient(addr string)
	Query(name string)
	Publish(name string, value []byte)
	Consume(name string)
}

//ClientRpc This struct contain RPC call data structure include connection
type ClientRpc struct {
	conn *rpc.Client
}

//NewClient Init related RPC call and dial to server by address<addr>
func NewClient(addr string) *ClientRpc {

	var err error

	c := new(ClientRpc)
	c.conn, err = rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return c
}

//Query  Declare Queue, will create new queue if it does not exist
func (c *ClientRpc) Query(name string) {
	// TODO
}

//Publish Publish data to the Queue
func (c *ClientRpc) Publish(name string, value []byte) {
	// TODO
}

//Consume Consume data from Queue until now, if there still not get any info return empty result
func (c *ClientRpc) Consume(name string) {
	// TODO
}
