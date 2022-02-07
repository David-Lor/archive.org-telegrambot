package telegrambot

import (
	"github.com/David-Lor/archive.org-telegrambot/internal/services/archiveorg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/David-Lor/archive.org-telegrambot/internal/settings"
)

type TelegramBot struct {
	bot              *tgbotapi.BotAPI
	archiveorgClient *archiveorg.Client
}

func NewTelegramBot(config settings.TelegramSettings, archiveorgClient *archiveorg.Client) (telegramBot *TelegramBot, err error) {
	bot, err := tgbotapi.NewBotAPI(config.Bot.Token)
	if err != nil {
		return
	}

	telegramBot = &TelegramBot{
		bot:              bot,
		archiveorgClient: archiveorgClient,
	}
	return
}
