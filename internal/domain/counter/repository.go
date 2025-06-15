package counter

type Repository interface {
	FindByChatAndUser(chatID, userID int64) (*CommandCounter, error)
	Save(counter *CommandCounter) error
	GetTopByChat(chatID int64, limit int) ([]*CommandCounter, error)
}
