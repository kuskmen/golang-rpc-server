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
type YamqClient struct {
	conn *rpc.Client
}

//NewClient Init related RPC call and dial to server by address<addr>
func NewClient(addr string) *YamqClient {

	var err error

	c := new(YamqClient)
	c.conn, err = rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return c
}

//Query  Query Queue, will create new queue if it does not exist
func (c *YamqClient) Query(name string) {
	var err error
	args := &QueryArgs{QueueName: name}
	var reply int
	err = c.conn.Call("WorkQueue.QueueDeclare", args, &reply)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	log.Println("WorkQueue: ", name)
}

//Publish Publish data to the Queue
func (c *YamqClient) Publish(name string, value []byte) {
	// TODO
}

//Subscribe Subscribes to a Queue
func (c *YamqClient) Subscribe(name string) {
	// TODO
}

//Consume Consume data from Queue until now, if there still not get any info return empty result
func (c *YamqClient) Consume(name string) {
	// TODO
}
