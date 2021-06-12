package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
	}
}

func (b *Bot) Start() error {
	// Set update interval
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Get updates
	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			// ToDo: add handling commands

			continue
		}

		// Handle regular messages
		// ToDo: add handling messages
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hello")
		_, _ = b.bot.Send(msg)
	}

	return nil
}
