package notice

var GlobalJobPool *JobPool
var GlobalChannel Channel

func InitNotification(token, chatID string) {
	sender := NewBot(token)
	jobPool := NewJobPool(sender, 5, -1, -1)
	jobPool.Start()
	GlobalJobPool = jobPool
	channel := Channel{
		ChatID: chatID,
	}
	GlobalChannel = channel
}
