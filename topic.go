package yamq

// Topic is generic interface of all message data structuress
type Topic interface {
	Send(message []byte)
	Subscribe(subscriber chan []byte)
	Unsubscribe(subscriber chan []byte)
}

type Warehouse struct {
	Queues []MessageQueue
}

// MessageQueue is structure for sending messages in FIFO mechanism.
type MessageQueue struct {
	Name string
	// TODO: Decide how data will be stored
	// Data
	Subscribers []chan []byte
}

// NewMessageQueue creates new MessageQueue with given string as a name.
func (w *Warehouse) NewMessageQueue(name string, reply *MessageQueue) error {
	// TODO: Validation for same name or perhaps disregard.
	queue := new(MessageQueue)
	queue.Name = name
	// TODO: Some goroutine to read the data.
	return nil
}

// Send sends messages to all subscribers of that MessageQueue
func (queue *MessageQueue) Send(message []byte, reply *MessageQueue) error {
	// TODO: Send message to subscribers
	return nil
}

// Subscribe adds subscribers to the list of subscribers of that MessageQueue
func (queue *MessageQueue) Subscribe(subscriber chan []byte, reply *MessageQueue) error {
	for _, v := range queue.Subscribers {
		if subscriber == v {
			return nil
		}
	}
	queue.Subscribers = append(queue.Subscribers, subscriber)
	return nil
}

// Unsubscribe removes subscriber from the list of subscribers of that MessageQueue
func (queue *MessageQueue) Unsubscribe(subscriber chan []byte, reply *MessageQueue) error {
	queue.Subscribers = removeChanFromSlice(queue.Subscribers, subscriber)
	return nil
}
