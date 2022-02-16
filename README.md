# Archive.org Telegram Bot

A Telegram bot that saves links sent to it on web.archive.org.

This bow follows an architecture consisting on:

- A Webhook updates receiver
- An AMQP broker (like RabbitMQ), where updates received are enqueued
- One or more workers (current repository), that consume updates from AMQP and process them
- Optionally, a self-hosted Telegram Bot API
