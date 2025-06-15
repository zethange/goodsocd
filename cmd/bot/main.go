package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zethange/goodsocd/internal/app/bot"
	"github.com/zethange/goodsocd/internal/domain/counter"
	"github.com/zethange/goodsocd/internal/infrastructure/db/counter_db"
	"xorm.io/xorm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	pathToDb := os.Getenv("PATH_TO_DB")
	if pathToDb == "" {
		pathToDb = "./counter.db"
	}

	engine, err := xorm.NewEngine("sqlite3", pathToDb)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer engine.Close()

	if err := engine.Sync2(new(counter_db.CommandCounterDB)); err != nil {
		log.Fatalf("Database sync failed: %v", err)
	}

	counterRepo := counter_db.NewXORMRepository(engine)
	counterService := counter.NewService(counterRepo)
	appService := bot.NewAppService(counterService)
	handlers := bot.NewHandlers(appService)

	botToken := os.Getenv("TELEGO_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGO_BOT_TOKEN environment variable not set")
	}

	bot, err := bot.NewTelegramBot(botToken, handlers)
	if err != nil {
		log.Fatalf("Bot initialization failed: %v", err)
	}

	bot.Start()
}
