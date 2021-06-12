package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/l-orlov/events-bot/internal/service"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	log *logrus.Logger
	svc *service.Service
	bot *tgbotapi.BotAPI
}

func NewBot(
	log *logrus.Logger,
	svc *service.Service,
	bot *tgbotapi.BotAPI,
) *Bot {
	return &Bot{
		log: log,
		svc: svc,
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

		b.log.Debugf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Handle commands
		if update.Message.IsCommand() {
			// ToDo: add handling commands

			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hello")
		btn := tgbotapi.KeyboardButton{
			RequestLocation: true,
			Text:            "Gimme where u live!!",
		}
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn})
		_, err = b.bot.Send(msg)
		if err != nil {
			b.log.Errorf("failed to send message: %v", err)
		}

		// Handle regular messages
		// ToDo: add handling messages
		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hello")
		//_, _ = b.bot.Send(msg)
	}

	return nil
}
