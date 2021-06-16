package telegram

import tb "gopkg.in/tucnak/telebot.v2"

var (
	InlineBtnProfile = tb.InlineButton{
		Unique: "profile",
		Text:   "Профиль",
	}
	InlineBtnEvents = tb.InlineButton{
		Unique: "events",
		Text:   "Мероприятия",
	}
	InlineBtnLocationMsk = tb.InlineButton{
		Unique: "locationMsk",
		Text:   "Москва",
	}
	InlineBtnLocationKzn = tb.InlineButton{
		Unique: "locationKzn",
		Text:   "Казань",
	}
)

func getKeysForButtons(buttons ...tb.InlineButton) (ret [][]tb.InlineButton) {
	for _, button := range buttons {
		ret = append(ret, []tb.InlineButton{button})
	}
	return
}

func MakeReplyMarkup(buttons ...tb.InlineButton) *tb.ReplyMarkup {
	return &tb.ReplyMarkup{
		InlineKeyboard: getKeysForButtons(buttons...),
	}
}
