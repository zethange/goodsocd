package bot

import "github.com/zethange/goodsocd/internal/domain/counter"

type AppService struct {
	counterService *counter.Service
}

func NewAppService(counterService *counter.Service) *AppService {
	return &AppService{
		counterService: counterService,
	}
}

func (s *AppService) RegisterCommand(chatID, userID int64, username, fullname string) error {
	return s.counterService.RegisterCommand(chatID, userID, username, fullname)
}

func (s *AppService) GetTopUsers(chatID int64) ([]*counter.CommandCounter, error) {
	return s.counterService.GetTopUsers(chatID, 10)
}
