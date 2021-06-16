package telegram

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	msgError          = "К сожалению что-то пошло не так. Попробуй снова путём ввода команды /start"
	msgIntro          = "Я бот, который помогает выбрать подходящее культурное мероприятие. Используй следующие кнопки:\nПрофиль - настроить свой профиль;\nМероприятия - получить информацию о культурных мероприятиях"
	msgChooseLocation = "Выбери город"
)

func (b *Bot) InitHandlers() {
	b.bot.Handle("/start", b.IntroHandler())
	b.bot.Handle(&InlineBtnProfile, b.ProfileHandler())
}

func (b *Bot) IntroHandler() func(*tb.Message) {
	funcName := "IntroHandler"

	return func(m *tb.Message) {
		_, err := b.bot.Send(m.Sender, msgIntro, MakeReplyMarkup(InlineBtnProfile, InlineBtnEvents), tb.ModeHTML)
		if err != nil {
			b.log.Errorf("%s: failed to send message: %v", funcName, err.Error())
			b.bot.Send(m.Sender, msgError)
			return
		}
	}
}

func (b *Bot) ProfileHandler() func(c *tb.Callback) {
	funcName := "ProfileHandler"

	return func(c *tb.Callback) {
		_, err := b.bot.Send(c.Sender, msgChooseLocation, MakeReplyMarkup(InlineBtnLocationMsk, InlineBtnLocationKzn), tb.ModeHTML)
		if err != nil {
			b.log.Errorf("%s: failed to send message: %v", funcName, err.Error())
			b.bot.Send(c.Sender, msgError)
			return
		}
	}
}
