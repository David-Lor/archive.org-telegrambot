package telegrambot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// isCommand verifies wether the given text has the given command
func isCommand(text string, cmd string) bool {
	// TODO Advanced checking: text="/startblabla" should not trigger cmd="/start"
	return strings.HasPrefix(text, cmd)
}

func extractURL(text string) (url string) {
	if strings.HasPrefix(text, "http://") || strings.HasPrefix(text, "https://") {
		url = text
	}
	return
}

func newReply(replyToMsg *tgbotapi.Message, text string) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(replyToMsg.Chat.ID, text)
	msg.ReplyToMessageID = replyToMsg.MessageID
	return &msg
}
