package rmq

import "fmt"

type Cleaner struct {
	connection *redisConnection
}

func NewCleaner(connection *redisConnection) *Cleaner {
	return &Cleaner{connection: connection}
}

func (cleaner *Cleaner) Clean() error {
	connectionNames, err := cleaner.connection.GetConnections()
	if err != nil {
		return err
	}

	for _, connectionName := range connectionNames {
		connection := cleaner.connection.hijackConnection(connectionName)
		if connection.Check() {
			continue // skip active connections!
		}

		if err := cleaner.CleanConnection(connection); err != nil {
			return err
		}
	}

	return nil
}

func (cleaner *Cleaner) CleanConnection(connection *redisConnection) error {
	queueNames, err := connection.GetConsumingQueues()
	if err != nil {
		return err
	}

	for _, queueName := range queueNames {
		queue, ok := connection.OpenQueue(queueName).(*redisQueue)
		if !ok {
			return fmt.Errorf("rmq cleaner failed to open queue %s", queueName)
		}

		cleaner.CleanQueue(queue)
	}

	if err := connection.Close(); err != nil {
		return fmt.Errorf("rmq cleaner failed to close connection %s: %s", connection, err)
	}

	if err := connection.CloseAllQueuesInConnection(); err != nil {
		return fmt.Errorf("rmq cleaner failed to close all queues %s %s", connection, err)
	}

	// log.Printf("rmq cleaner cleaned connection %s", connection)
	return nil
}

func (cleaner *Cleaner) CleanQueue(queue *redisQueue) {
	returned, _ := queue.ReturnAllUnacked()
	queue.CloseInConnection()
	_ = returned
	// log.Printf("rmq cleaner cleaned queue %s %d", queue, returned)
}
