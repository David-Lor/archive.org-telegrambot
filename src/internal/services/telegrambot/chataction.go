package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
	"time"
)

type ChatAction struct {
	bot       *TelegramBot
	chatID    int64
	action    string
	stop      bool
	onceStart sync.Once
}

func (action *ChatAction) Start() {
	action.onceStart.Do(func() {
		go action.worker()
	})
}

func (action *ChatAction) Stop() {
	action.stop = true
}

func (action *ChatAction) worker() {
	chatAction := tgbotapi.NewChatAction(action.chatID, action.action)
	log.Printf("Start ChatAction %s for %d", action.action, action.chatID)
	for {
		if action.stop {
			break
		}

		_, _ = action.bot.bot.Send(chatAction)
		time.Sleep(ChatActionPeriodMillis * time.Millisecond)
	}
	log.Printf("Stop ChatAction %s for %d", action.action, action.chatID)
}

func NewChatAction(bot *TelegramBot, chatID int64, action string) *ChatAction {
	return &ChatAction{
		bot:    bot,
		chatID: chatID,
		action: action,
		stop:   false,
	}
}
