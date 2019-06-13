package main

import (
	// "fmt"
	"log"
	"os"
	// "net/http"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"encoding/json"
)

func main() {
	log.Printf("Starting duckhunt\n")

	data := []byte(`{"ok":true,"result":[{"update_id":418076391,	"message":{"message_id":27,"from":{"id":861117934,"is_bot":false,"first_name":"Bart","last_name":"van Hunnik","language_code":"nl"},"chat":{"id":861117934,"first_name":"Bart","last_name":"van Hunnik","type":"private"},"date":1559764150,"photo":[{"file_id":"AgADBAAD2K4xGyDowVMP8q7bFCvnog_XLBsABL1hgo6sL4HMFHIEAAEC","file_size":2176,"width":90,"height":90},{"file_id":"AgADBAAD2K4xGyDowVMP8q7bFCvnog_XLBsABN8jTyQ9EQ4LFXIEAAEC","file_size":21651,"width":320,"height":320},{"file_id":"AgADBAAD2K4xGyDowVMP8q7bFCvnog_XLBsABHz2BLw8MzwpFnIEAAEC","file_size":60221,"width":640,"height":640}]}}]}`)
	
	type TelegPhoto struct {
		File_id string `json:"file_id"`
		File_size int `json:"file_size"`
		Width int `json:"width"`
		Height int `json:"height"`
	}

	type TelegFrom struct {
		Id int `json:"id"`
		Is_bot bool `json:"is_bot"`
		First_name string `json:"first_name"`
		Last_name string `json:"last_name"`
		Language_code string `json:"language_code"`
	}

	type TelegChat struct {
		Id int `json:"id"`
		Is_bot bool `json:"is_bot"`
		First_name string `json:"first_name"`
		Last_name string `json:"last_name"`
		Type string `json:"type"`
	}
	type TelegMessage struct {
		Message_id int `json:"message_id"`
		From TelegFrom `json:"from"`
		Chat TelegChat `json:"chat"`
		Photo []TelegPhoto `json:"photo"`
	}
	type TelegUpdate struct {
		UpdateId int `json:"update_id"`
		Message TelegMessage `json:"message"`
	}
	type TelegReturnMessage struct {
		Result []TelegUpdate `json:"result"`
	}


	var returnMessage TelegReturnMessage
	if err := json.Unmarshal(data, &returnMessage); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", returnMessage)

	// var dat map[string]interface{}
	// if err := json.Unmarshal(byt, &dat); err != nil {
	// 	panic(err)
	// }
	// log.Println(dat["result"])

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_KEY"))
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		log.Printf("New message from %s %s", update.Message.From.FirstName, update.Message.From.LastName)
		log.Printf("%s", update.Message)
		//SendWithExistingPhoto(update.Message.Chat.ID, update.Message.photo[0].file_id)
		
		//msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, update.Message.Chat.big_file_id)
		//msg.Caption = "here you go"

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		//_, err := bot.Send(msg)

		if err != nil {
			log.Fatal(err)
		} else {log.Printf("Success")}
	}
}

// func SendWithExistingPhoto(bot, ChatID, ExistingPhotoFileID) {

// 	msg := tgbotapi.NewPhotoShare(bot, ChatID, ExistingPhotoFileID)
// 	msg.Caption = "Lekker hoor"
// 	_, err := bot.Send(msg)

// 	if err != nil {
// 		t.Error(err)
// 		t.Fail()
// 	}
// }