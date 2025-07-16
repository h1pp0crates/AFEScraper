package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var tgToken, chatId, auchanURL, fozzy, epicentrk string

func SendTgMessage(auchanPrice, fozzyPrice, epicentrkPrice string) {
	importEnvs()

	u := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tgToken)
	now := time.Now().Format("mondey, 02.01.2006")

	text := fmt.Sprintf(`
	<b> Вітаю!</b>
	<i>Сьогодні %s</i>
	- <a href="%s">Ашан</a> - %s 
	- <a href="%s">Fozzy</a> - %s 
	- <a href="%s">Епіцентр</a> - %s 
	`, now, auchanURL, auchanPrice, fozzy, fozzyPrice, epicentrk, epicentrkPrice)

	b, err := json.Marshal(map[string]string{
		"chat_id":    chatId,
		"text":       text,
		"parse_mode": "HTML",
	})
	if err != nil {
		log.Fatalln("Json marshal error:\n", err)
	}

	res, err := http.Post(u, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatalln("Post request error:\n", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalln("Telegram API error:\n", res.Status)
	}

	fmt.Println("Message was succesfully sent")
}

func importEnvs() {
	tgToken = os.Getenv("tgToken")
	chatId = os.Getenv("chatId")
	auchanURL = os.Getenv("auchanURL")
	fozzy = os.Getenv("fozzy")
	epicentrk = os.Getenv("epicentrk")

	if tgToken == "" || chatId == "" || auchanURL == "" || fozzy == "" || epicentrk == "" {
		log.Fatalln("One or more required environment variables for function sendTgMessage are missing")
	}
}
