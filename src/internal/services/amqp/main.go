package amqp

import (
	"github.com/David-Lor/archive.org-telegrambot/internal/settings"
	"github.com/gammazero/workerpool"
	"github.com/streadway/amqp"
)

type ConsumerCallback func(amqp.Delivery) error

type Consumer struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	consumerTag  string
	queue        string
	callback     ConsumerCallback
	deliveryChan <-chan amqp.Delivery
	closeChan    chan *amqp.Error
	handlerChan  chan error
	workPool     *workerpool.WorkerPool
}

func NewConsumer(settings settings.AMQPSettings) (consumer *Consumer, err error) {
	consumer = &Consumer{
		consumerTag: settings.ConsumerTag,
		queue:       settings.Queue,
		workPool:    workerpool.New(settings.ConsecutiveRequestsLimit),
	}

	consumer.connection, err = amqp.Dial(settings.URI)
	if err != nil {
		return
	}

	consumer.closeChan = consumer.connection.NotifyClose(make(chan *amqp.Error))
	consumer.channel, err = consumer.connection.Channel()
	if err != nil {
		consumer.Close()
	}
	return
}

func (consumer *Consumer) SetCallback(callback ConsumerCallback) {
	consumer.callback = callback
}

func (consumer *Consumer) Close() {
	if consumer.channel != nil {
		_ = consumer.channel.Cancel(consumer.consumerTag, true)
	}
	if consumer.connection != nil {
		_ = consumer.connection.Close()
	}
}

func (consumer *Consumer) WaitUntilClosed() error {
	return <-consumer.closeChan
}
