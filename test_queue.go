package rmq

import "time"

type TestQueue struct {
	name           string
	LastDeliveries []string
}

func NewTestQueue(name string) *TestQueue {
	queue := &TestQueue{name: name}
	queue.Reset()
	return queue
}

func (queue *TestQueue) String() string {
	return queue.name
}

func (queue *TestQueue) Publish(payload string) error {
	queue.LastDeliveries = append(queue.LastDeliveries, payload)
	return nil
}

func (queue *TestQueue) PublishBytes(payload []byte) error {
	return queue.Publish(string(payload))
}

func (queue *TestQueue) SetPushQueue(pushQueue Queue) {
}

func (queue *TestQueue) StartConsuming(prefetchLimit int, pollDuration time.Duration) error {
	return nil
}

func (queue *TestQueue) StopConsuming() error {
	return nil
}

func (queue *TestQueue) AddConsumer(tag string, consumer Consumer) (string, error) {
	return "", nil
}

func (queue *TestQueue) AddConsumerFunc(tag string, consumerFunc ConsumerFunc) (string, error) {
	return "", nil
}

func (queue *TestQueue) AddBatchConsumer(tag string, batchSize int, consumer BatchConsumer) (string, error) {
	return "", nil
}

func (queue *TestQueue) AddBatchConsumerWithTimeout(tag string, batchSize int, timeout time.Duration, consumer BatchConsumer) (string, error) {
	return "", nil
}

func (queue *TestQueue) ReturnRejected(count int) (int, error) {
	return 0, nil
}

func (queue *TestQueue) ReturnAllRejected() (int, error) {
	return 0, nil
}

func (queue *TestQueue) PurgeReady() (int, error) {
	return 0, nil
}

func (queue *TestQueue) PurgeRejected() (int, error) {
	return 0, nil
}

func (queue *TestQueue) Close() error {
	return nil
}

func (queue *TestQueue) Reset() {
	queue.LastDeliveries = []string{}
}
