package counter

type CommandCounter struct {
	ID       int64
	ChatID   int64
	UserID   int64
	UserName string
	FullName string
	Count    int
}

func NewCommandCounter(chatID, userID int64, username string, fullname string) *CommandCounter {
	return &CommandCounter{
		ChatID:   chatID,
		UserID:   userID,
		UserName: username,
		FullName: fullname,
		Count:    1,
	}
}

func (c *CommandCounter) Increment() {
	c.Count++
}

func (c *CommandCounter) UpdateUsername(username string) {
	c.UserName = username
}

func (c *CommandCounter) UpdateFullname(fullname string) {
	c.FullName = fullname
}
