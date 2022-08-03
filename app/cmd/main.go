package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BotToken"))
	if err != nil {
		log.Panicln(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	var ucfg = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updatesChan := bot.GetUpdatesChan(ucfg)
	for update := range updatesChan {
		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "party":
			poll := tgbotapi.SendPollConfig{
				BaseChat: tgbotapi.BaseChat{
					ChatID: update.Message.Chat.ID,
				},
				Question:             "Тусим в пятницу?",
				Options:              []string{"Конечно!", "Неее", "Я глупый и не могу определиться"},
				CorrectOptionID:      1,
				Explanation:          "",
				ExplanationParseMode: "",
				ExplanationEntities:  nil,
				OpenPeriod:           0,
				CloseDate:            0,
				IsClosed:             false,
			}
			_, err = bot.Send(poll)
			if err != nil {
				log.Panicln(err)
			}
		}
	}
}
