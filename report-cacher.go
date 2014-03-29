package main

import (
	"github.com/jfmarket/report-cacher/download"
	"log"
)

func main() {
	log.Println("Starting...")
	downloader, err := download.New("https://jonesboroughfarmersmkt.shopkeepapp.com", "chad@snapstudent.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	err = downloader.GetSoldItemsReport("files/sold_items.csv", "2014-02-28", "2014-03-29")
	if err != nil {
		log.Fatalln(err)
	}
}
