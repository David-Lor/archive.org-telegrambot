package main

import (
	"github.com/David-Lor/archive.org-telegrambot/internal/services/amqp"
	"github.com/David-Lor/archive.org-telegrambot/internal/services/archiveorg"
	"github.com/David-Lor/archive.org-telegrambot/internal/services/telegrambot"
	"github.com/David-Lor/archive.org-telegrambot/internal/settings"
	amqplib "github.com/streadway/amqp"
)

func main() {
	config, err := settings.LoadSettings("")
	if err != nil {
		panic(err)
	}

	archiveorgClient := archiveorg.NewArchiveOrgClient(config.Archiveorg)
	bot, err := telegrambot.NewTelegramBot(config.Telegram, archiveorgClient)
	if err != nil {
		panic(err)
	}

	amqpClient, err := amqp.NewConsumer(config.AMQP)
	if err != nil {
		panic(err)
	}

	amqpClient.SetCallback(func(msg amqplib.Delivery) error {
		return bot.HandlerEntrypoint(msg.Body)
	})

	err = amqpClient.Start()
	if err != nil {
		panic(err)
	}

	_ = amqpClient.WaitUntilClosed() // if amqp connection closed by any reason, stop app (with failure)
}
