package telegram

import (
	"AFEScraper/internal"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendTgMessage(cfg internal.AppConfig, auchanPrice, fozzyPrice, epicentrkPrice string) {
	u := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.TgToken)

	text := fmt.Sprintf(`
	<b> Вітаю!</b>
	<i>Сьогодні %s</i>
	- <a href="%s">Ашан</a> - %s 
	- <a href="%s">Fozzy</a> - %s 
	- <a href="%s">Епіцентр</a> - %s 
	`, internal.DateNow(), cfg.AuchanURL, auchanPrice, cfg.FozzyURL, fozzyPrice, cfg.EpicentrkURL, epicentrkPrice)

	b, err := json.Marshal(map[string]string{
		"chat_id":    cfg.ChatId,
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

	log.Println("Message was succesfully sent")
}
