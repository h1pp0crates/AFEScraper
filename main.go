package main

import (
	"log"
	"os"
	"parser/sites"
	"parser/telegram"

	"github.com/joho/godotenv"
)

var auchanURLKey, epicentrk, fozzy string

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Fatalln(e)
	}

	fozzy = os.Getenv("fozzy")
	auchanURLKey = os.Getenv("auchan")
	epicentrk = os.Getenv("epicentrk")

	if fozzy == "" || auchanURLKey == "" || epicentrk == "" {
		log.Fatalln("One or more required environment variables for main function are missing")
	}
}

func main() {
	log.Println("Starting...")

	auchanPrice := sites.AuchanScrape(auchanURLKey)
	fozzyPrice := sites.FozzyScrape(fozzy)
	epicentrkPrice := sites.EpicentrkScrape(epicentrk)

	telegram.SendTgMessage(auchanPrice, fozzyPrice, epicentrkPrice)
}
