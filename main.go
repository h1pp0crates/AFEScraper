// AFEScraper: Collects prices from Auchan, Fozzy, and Epicentrk and sends them to Telegram daily.

package main

import (
	"AFEScraper/internal"
	"AFEScraper/sites"
	"AFEScraper/telegram"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	cfg := internal.AppConfig{
		TgToken:      os.Getenv("TgToken"),
		ChatId:       os.Getenv("ChatId"),
		FozzyURL:     os.Getenv("FozzyURL"),
		AuchanURLKey: os.Getenv("AuchanURLKey"),
		AuchanURL:    os.Getenv("AuchanURL"),
		EpicentrkURL: os.Getenv("EpicentrkURL"),
	}

	if cfg.TgToken == "" || cfg.ChatId == "" || cfg.FozzyURL == "" || cfg.AuchanURLKey == "" || cfg.EpicentrkURL == "" {
		log.Fatalln("One or more required environment variables for main function are missing")
	}

	c := cron.New()
	c.AddFunc("0 0 11 * * *", func() {
		auchanPrice := sites.AuchanScrape(cfg.AuchanURLKey)
		fozzyPrice := sites.FozzyScrape(cfg.FozzyURL)
		epicentrkPrice := sites.EpicentrkScrape(cfg.EpicentrkURL)

		telegram.SendTgMessage(cfg, auchanPrice, fozzyPrice, epicentrkPrice)
	})
	c.Run()

}
