package main

import (
	"log"

	"github.com/praveenvoonna/go-cricket-analytics-api/csv"
	"github.com/praveenvoonna/go-cricket-analytics-api/handler"
)

func main() {
	if err := loadData(); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	router := handler.SetupHandlers()

	log.Println("Server started at :8080")
	router.Run(":8080")
}

func loadData() error {
	filePath := "data/ODI Data.csv"
	_, err := csv.ReadCSVLoadMaps(filePath)
	if err != nil {
		return err
	}

	return nil
}
