package workspace

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Server() {
	bot, err := tgbotapi.NewBotAPI("5585348727:AAGJCM7EWOQJTBFxXe5BsWmUpSQOWiQ_XqU")
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true
	var responce string
	var fileName string
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if len(strings.Fields(update.Message.Text)) != 1 {
				responce = Responce(update.Message.Text)
			} else {
				switch update.Message.Text {
				case "Kms":
					fileName = "Kms.exe"

				case "/a":
					go bot.Send(Video(update.Message.From.ID, update.Message.From.UserName, "a"))

				case "/t":
					fileName = "t.jpeg"
					go bot.Send(Video(update.Message.From.ID, update.Message.From.UserName, "t"))
				case "/p":
					fileName = "p.jpeg"
				case "/s":
					fileName = "s.jpeg"
				case "/nt":
					fileName = "nt.jpeg"
				case "/f1":
					fileName = "f1.jpeg"
				}
				photo := tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FilePath("PhotoProblem/"+fileName))

				bot.Send(photo)
				fileName = "" // чтобы обнулять имя фото чтобы недопустить дублирования
				responce = Command(update.Message.Text)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, responce)

			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

//
