package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	diskqueue "github.com/kkdai/diskqueue"
)

type Warehouse struct {
	sync.RWMutex

	//To quick refer if queue exist
	queueMap map[string]int

	//map to store "queue -> chan List" for publish
	subscriptions map[string]*WorkQueue
}

func NewWarehouse() *Warehouse {
	warehouse := &Warehouse{
		queueMap:      make(map[string]int),
		subscriptions: make(map[string]*WorkQueue),
	}

	return warehouse
}

// WorkQueue is structure for sending messages in FIFO mechanism.
type WorkQueue struct {
	Name        string
	dataQueue   diskqueue.WorkQueue
	subscribers []chan []byte
	exit        chan int
}

// NewWorkQueue creates new MessageQueue with given string as a name.
func (w *Warehouse) NewWorkQueue(name string, reply *WorkQueue) error {
	if _, ok := w.queueMap[name]; ok {
		return nil
	}

	dir, err := ioutil.TempDir("Diskqueue", fmt.Sprintf("%s - %v", reply.Name, time.Now().UTC()))
	if err != nil {
		log.Println("can't create temporary directory for diskqueue")
		return err
	}

	reply = &WorkQueue{
		Name:      name,
		dataQueue: diskqueue.NewDiskqueue(name, dir),
		exit:      make(chan int),
	}

	go reply.work()

	return nil
}

// Send sends messages to all subscribers of that WorkQueue
func (wq *WorkQueue) Send(message []byte, reply *WorkQueue) error {
	var err = wq.dataQueue.Put(message)
	if err != nil {
		return errors.New("diskQueue failed to send data to subscribers")
	}

	return nil
}

// Subscribe adds subscribers to the list of subscribers of that WorkQueue
func (wq *WorkQueue) Subscribe(subscriber chan []byte, reply *WorkQueue) error {
	for _, v := range wq.subscribers {
		if subscriber == v {
			return nil
		}
	}
	wq.subscribers = append(wq.subscribers, subscriber)
	return nil
}

// Unsubscribe removes subscriber from the list of subscribers of that WorkQueue
func (wq *WorkQueue) Unsubscribe(subscriber chan []byte, reply *WorkQueue) error {
	wq.subscribers = removeChanFromSlice(wq.subscribers, subscriber)
	return nil
}

func (wq *WorkQueue) work() {
	var dataRead []byte
	for {
		dataRead = nil
		select {
		case dataRead = <-wq.dataQueue.ReadChan():
			log.Println("TOPIC:Got data ", string(dataRead))
			//Got data start to put on all chan
			for _, targetChan := range wq.subscribers {
				targetChan <- dataRead
			}
		case <-wq.exit:
			//exist
			wq.dataQueue.Empty()
			wq.dataQueue.Close()
		}
	}
}
