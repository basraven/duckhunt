package telegram


import (
	"log"
	"os"
	"strconv"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)


func Send(message string) {
	log.Printf("SEND: %s bij deze key \n\r", message)	
	
	ChatId, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_KEY"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	msg := tgbotapi.NewMessage(int64(ChatId), message)
	bot.Send(msg)
}