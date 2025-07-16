package main

import (
	"fmt"
	"log"
	"os"
	"parser/sites"

	"github.com/joho/godotenv"
)

var tgToken, chatId, auchan, epicentrk, fozzy string

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Fatalln(e)
	}
	tgToken = os.Getenv("tgToken")
	chatId = os.Getenv("chatId")
	fozzy = os.Getenv("fozzy")
	auchan = os.Getenv("auchan")
	epicentrk = os.Getenv("epicentrk")
}

func main() {
	fmt.Println("In fozzy:")
	sites.FozzyScrape(fozzy)
	fmt.Println("\nIn auchane")
	sites.AuchanScrape(auchan)
	fmt.Println("\nIn epicentrk:")
	sites.EpicentrkScrape(epicentrk)
}
