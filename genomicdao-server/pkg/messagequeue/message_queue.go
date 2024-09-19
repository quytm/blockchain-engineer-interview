package msgqueue

// MessageQueue ...
type MessageQueue struct {
	queue chan interface{}
}

// NewMessageQueue ...
func NewMessageQueue(bufferSize int) *MessageQueue {
	return &MessageQueue{
		queue: make(chan interface{}, bufferSize),
	}
}

// Publish message to queue.
func (mq *MessageQueue) Publish(message interface{}) {
	mq.queue <- message
}

// Listen message from queue.
func (mq *MessageQueue) Listen(handleMessage func(interface{})) {
	for msg := range mq.queue {
		handleMessage(msg)
	}
}
