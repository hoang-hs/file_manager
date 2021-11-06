package notice

type Channel struct {
	ChatID string
}

func NewChannel(chatID string) *Channel {
	return &Channel{
		ChatID: chatID,
	}
}
