package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"

	"github.com/David-Lor/archive.org-telegrambot/pkg/utils"
)

func (bot *TelegramBot) HandlerEntrypoint(payload []byte) error {
	// TODO Panic handler
	request := &http.Request{
		Method: http.MethodPost,
		Body:   utils.BytesToIOReadCloser(payload),
	}

	update, err := bot.bot.HandleUpdate(request)
	if err != nil {
		log.Printf("Error parsing update: %v", err)
		return nil // avoid re-enqueue
	}
	if update.UpdateID == 0 {
		log.Printf("Invalid update payload received: %s", string(payload))
		return nil // avoid re-enqueue
	}

	log.Printf("RX Telegram Update: %+v", update)
	if update.Message != nil {
		text := update.Message.Text
		if isCommand(text, CommandStart) {
			return bot.commandBasicReply(update, CommandStartReply)
		}
		if isCommand(text, CommandHelp) {
			return bot.commandBasicReply(update, CommandHelpReply)
		}
		if url := extractURL(text); url != "" {
			return bot.commandArchiveURL(update, url)
		}
	}

	return nil
}

func (bot *TelegramBot) commandBasicReply(update *tgbotapi.Update, replyText string) error {
	msg := newReply(update.Message, replyText)
	_, err := bot.bot.Send(msg)
	return err
}

func (bot *TelegramBot) commandArchiveURL(update *tgbotapi.Update, url string) error {
	msg := newReply(update.Message, CommandArchiveStartReply)
	_, err := bot.bot.Send(msg)
	if err != nil {
		return err
	}

	typingChatAction := NewChatAction(bot, update.Message.Chat.ID, tgbotapi.ChatTyping)
	typingChatAction.Start()
	defer typingChatAction.Stop()

	var replyMsg *tgbotapi.MessageConfig
	resultURL, err := bot.archiveorgClient.ArchiveURL(url)
	if err == nil && resultURL != "" {
		replyMsg = newReply(update.Message, CommandArchiveCompletedReply+resultURL)
	} else {
		replyMsg = newReply(update.Message, CommandArchiveErrorReply)
	}

	typingChatAction.Stop()
	_, err = bot.bot.Send(replyMsg)
	return err
}
