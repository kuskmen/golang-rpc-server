package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const maxNumberOfQueues = 65535

type YamqServer struct {
	workQueues []string
}

type RegisterWorkQueueArgs struct {
	Name string
	Port string
}

var yamqServer = new(YamqServer)

// NewServer creates and starts new rpc server on given port
func NewServer(port string) {
	forever := make(chan int)

	err := rpc.Register(yamqServer)
	if err != nil {
		log.Fatal("Error registering yamq server.", err)
	}

	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", port)
	log.Println("start to listen:", listener)

	if e != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(listener, nil)

	<-forever
}

func (ymServer *YamqServer) RegisterWorkQueue(args RegisterWorkQueueArgs, reply bool) error {

	//err := rpc.RegisterName(args.Name, new(WorkQueue))
	//if err != nil {
	//return err
	//}

	ymServer.workQueues = append(ymServer.workQueues, args.Name)

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", args.Port)

	if err != nil {
		return err
	}

	go http.Serve(listener, nil)

	return nil
}
