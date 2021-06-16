package telegram

import (
	"github.com/l-orlov/events-bot/internal/service"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

type Bot struct {
	log *logrus.Logger
	svc *service.Service
	bot *tb.Bot
}

func NewBot(
	log *logrus.Logger,
	svc *service.Service,
	telegramToken string,
) (*Bot, error) {
	bot, err := tb.NewBot(tb.Settings{
		Token:  telegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	b := &Bot{
		log: log,
		svc: svc,
		bot: bot,
	}
	b.InitHandlers()

	return b, err
}

func (b *Bot) Start() {
	b.bot.Start()
}

func (b *Bot) Stop() {
	b.bot.Stop()
}
