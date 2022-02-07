package amqp

import (
	"log"
)

func (consumer *Consumer) Start() (err error) {
	consumer.deliveryChan, err = consumer.channel.Consume(
		consumer.queue,       // queue name
		consumer.consumerTag, // consumer tag
		false,                // autoAck
		false,                // exclusive
		false,                // noLocal
		false,                // noWait
		nil,                  // arguments
	)
	if err != nil {
		return err
	}

	log.Print("AMQP consumer started")
	go consumer.consumerWorker()
	return
}

func (consumer *Consumer) consumerWorker() {
	log.Print("AMQP consumerWorker: working")
	for delivery := range consumer.deliveryChan {
		var err error
		log.Printf("AMQP RX: %v", string(delivery.Body))

		if consumer.callback != nil {
			// TODO Problem with this pool: messages are received by consumer but unacked? so they could not be acquired by other consumers (except for round-robin nature of rabbitmq?)
			consumer.workPool.SubmitWait(func() {
				err = consumer.callback(delivery)
			})
		}

		if err == nil {
			_ = delivery.Ack(false)
		} else {
			_ = delivery.Nack(false, true)
		}

	}

	log.Print("AMQP consumerWorker: stopped")
}
