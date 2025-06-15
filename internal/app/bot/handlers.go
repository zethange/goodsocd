package bot

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/rb-go/plural-ru"
)

type Handlers struct {
	service *AppService
}

func NewHandlers(service *AppService) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) Register(bh *telegohandler.BotHandler) {
	bh.Handle(h.HandleCommand1600x720, telegohandler.CommandEqual("1600x720"))
	bh.Handle(h.HandleCommand1600x720Top, telegohandler.CommandEqual("1600x720_top"))
}

func (h *Handlers) HandleCommand1600x720(ctx *telegohandler.Context, update telego.Update) error {
	msg := update.Message
	chatID := msg.Chat.ID
	userID := msg.From.ID
	username := ""
	if msg.From.Username != "" {
		username = "@" + msg.From.Username
	}
	if err := h.service.RegisterCommand(chatID, userID, username, msg.From.FirstName+" "+msg.From.LastName); err != nil {
		log.Printf("Error registering command: %v", err)
	}

	return nil
}

func (h *Handlers) HandleCommand1600x720Top(ctx *telegohandler.Context, update telego.Update) error {
	msg := update.Message
	chatID := msg.Chat.ID

	top, err := h.service.GetTopUsers(chatID)
	if err != nil {
		log.Printf("Error getting top users: %v", err)
		_, _ = ctx.Bot().SendMessage(context.Background(), &telego.SendMessageParams{
			ChatID: telego.ChatID{ID: chatID},
			Text:   "❌ Ошибка при получении данных",
		})
		return nil
	}

	if len(top) == 0 {
		_, _ = ctx.Bot().SendMessage(context.Background(), &telego.SendMessageParams{
			ChatID: telego.ChatID{ID: chatID},
			Text:   "В этом чате еще никто не /1600x720",
		})
		return nil
	}

	response := "🏆 Топ пользователей по /1600x720:\n\n"
	for i, user := range top {
		name := user.FullName
		link := user.UserName

		if name == "" {
			name = link[1:]
		}

		if strings.HasPrefix(link, "@") {
			link = "tg://resolve?domain=" + link[1:]
		} else {
			link = "tg://user?id=" + strconv.Itoa(int(user.UserID))
		}

		response += fmt.Sprintf("%d. <a href=\"%s\">%s</a> — %d %s\n", i+1, link, name, user.Count, plural.Noun(user.Count, "раз", "раза", "раз"))
	}

	_, _ = ctx.Bot().SendMessage(context.Background(), &telego.SendMessageParams{
		ChatID:          telego.ChatID{ID: chatID},
		Text:            response,
		ReplyParameters: &telego.ReplyParameters{MessageID: msg.MessageID},
		ParseMode:       "HTML",
	})

	return nil
}
