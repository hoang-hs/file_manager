package notice

import "time"

const (
	BaseAPIEndpointURL    = "https://api.telegram.org/"
	BaseBotApiEndpointURL = BaseAPIEndpointURL + "bot"
)

const (
	SendMessage = "SendMessage"
)

const (
	GapTimeRequest        = 3 * time.Second
	GapTimeTooManyRequest = 1 * time.Minute
)
