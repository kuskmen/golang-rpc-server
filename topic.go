package golangrpcserver

// Topic is generic interface of all message data structuress
type Topic interface {
	Send(message []byte)
	Subscribe(subscriber chan []byte)
	Unsubscribe(subscriber chan []byte)
}

// MessageQueue is structure for sending messages in FIFO mechanism.
type MessageQueue struct {
	Name string
	// TODO: Decide how data will be stored
	// Data
	Subscribers []chan []byte
}

// NewMessageQueue creates new MessageQueue with given string as a name.
func NewMessageQueue(name string) *MessageQueue {
	// TODO: Validation for same name or perhaps disregard.
	queue := new(MessageQueue)
	queue.Name = name
	// TODO: Some goroutine to read the data.
	return queue
}

// Send sends messages to all subscribers of that MessageQueue
func (queue *MessageQueue) Send(message []byte) {
	// TODO: Send message to subscribers
}

// Subscribe adds subscribers to the list of subscribers of that MessageQueue
func (queue *MessageQueue) Subscribe(subscriber chan []byte) {
	for _, v := range queue.Subscribers {
		if subscriber == v {
			return
		}
	}
	queue.Subscribers = append(queue.Subscribers, subscriber)
}

// Unsubscribe removes subscriber from the list of subscribers of that MessageQueue
func (queue *MessageQueue) Unsubscribe(subscriber chan []byte) {
	queue.Subscribers = removeChanFromSlice(queue.Subscribers, subscriber)
}
