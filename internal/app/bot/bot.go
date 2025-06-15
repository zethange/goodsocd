package bot

import (
	"context"
	"log"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
)

type TelegramBot struct {
	bot      *telego.Bot
	handlers *Handlers
}

func NewTelegramBot(token string, handlers *Handlers) (*TelegramBot, error) {
	b, err := telego.NewBot(token, telego.WithDefaultLogger(false, true))
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		bot:      b,
		handlers: handlers,
	}, nil
}

func (t *TelegramBot) Start() {
	updates, err := t.bot.UpdatesViaLongPolling(context.Background(), &telego.GetUpdatesParams{})
	if err != nil {
		log.Fatalf("Failed to get updates: %v", err)
	}

	bh, err := telegohandler.NewBotHandler(t.bot, updates)
	if err != nil {
		log.Fatalf("Failed to create bot handler: %v", err)
	}
	defer bh.Stop()

	t.handlers.Register(bh)

	log.Println("Bot started and handlers registered")
	bh.Start()
}
