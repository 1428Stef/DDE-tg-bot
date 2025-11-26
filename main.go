package main

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var button = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Get dose!"),
	),
)

var textStart string = "Hi, press the button for a dose of English X﹏X"

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Fatalf("Error token:%v", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyMarkup = button

			if update.Message.Command() == "start" {
				msg.Text = textStart
			}

			switch update.Message.Text {
			case "Get dose!":
				msg.Text = "Get your dose: " + randomWord()
			default:
				msg.Text = "Sorry, unknown command ＞﹏＜"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Fatal(err)
			}
		}
	}
}
