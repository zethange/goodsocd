package counter

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) RegisterCommand(chatID, userID int64, username string, fullname string) error {
	counter, err := s.repo.FindByChatAndUser(chatID, userID)
	if err != nil {
		return err
	}

	if counter == nil {
		counter = NewCommandCounter(chatID, userID, username, fullname)
	} else {
		counter.Increment()
		counter.UpdateUsername(username)
	}

	return s.repo.Save(counter)
}

func (s *Service) GetTopUsers(chatID int64, limit int) ([]*CommandCounter, error) {
	return s.repo.GetTopByChat(chatID, limit)
}
