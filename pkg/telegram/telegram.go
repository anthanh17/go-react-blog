package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/anthanh17/go-react-blog/pkg/config"
)

type sendMessageReqBody struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func sendMessageTelegram(
	chatID int64,
	botId int64,
	token string,
	message string) error {
	// Create request body struct
	repBody := &sendMessageReqBody{
		ChatId: chatID,
		Text:   message,
	}
	// Create the json body
	repBytes, err := json.Marshal(repBody)
	if err != nil {
		return err
	}
	// Send a post request with your token
	urlEndpoint := fmt.Sprintf("https://api.telegram.org/bot%d:%s/sendMessage", botId, token)
	res, err := http.Post(urlEndpoint, "application/json", bytes.NewBuffer(repBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}
	return nil
}

func send(message string) error {
	return sendMessageTelegram(
		config.TelegramCfg.ChatId,
		config.TelegramCfg.BotId,
		config.TelegramCfg.Token,
		message)
}

func SendMessage(message string, flag int) {
	var emoji string

	switch flag {
	case 0:
		emoji = "❎"
	case 1:
		emoji = "✅"
	case 2:
		emoji = "❌"
	default:
		log.Println("unknown level %d", flag)
		return
	}

	go func() {
		err := send(fmt.Sprintf("%s - %s", emoji, message))
		if err != nil {
			log.Println("Can not send message telegram: %s", err.Error())
		}
	}()
}
