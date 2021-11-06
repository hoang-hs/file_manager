package notice

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Bot struct {
	Token      string
	BaseURL    string
	httpClient *http.Client
}

func NewBot(token string) *Bot {
	baseURL := BaseBotApiEndpointURL + token + "/"
	return &Bot{
		Token:      token,
		BaseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (bot *Bot) Send(p Package) Response {
	return bot.SendMessage(p.Channel, p.Template)
}

func (bot *Bot) SendMessage(teleChan Channel, template Template) Response {
	var message string
	switch template.ParseMode {
	case MarkdownV2:
		// TODO preceding special character with '/'
		message = fmt.Sprintf("%s%s%s %s %s", "*", template.Job, "*", "%0A", template.Message)
	case HTML:
		// TODO maybe preceding too :((
		message = fmt.Sprintf("<strong>%s</strong> %s <pre>%s</pre>", template.Job, "\n", template.Message)
	}
	x := bot.BaseURL + SendMessage
	req, _ := http.NewRequest("GET", x, nil)
	q := req.URL.Query()
	q.Add("chat_id", teleChan.ChatID)
	q.Add("text", message)
	q.Add("parse_mode", template.ParseMode)
	req.URL.RawQuery = q.Encode()
	response, err := bot.httpClient.Do(req)
	if err != nil {
		return Response{
			Code: 400,
			Err:  err,
		}
	}
	defer func() {
		if response != nil && response.Body != nil {
			_ = response.Body.Close()
		}
	}()
	if response.StatusCode != 200 {
		err = errors.New("response error with status code " + strconv.Itoa(response.StatusCode))
	}
	return Response{
		Code: response.StatusCode,
		Err:  err,
	}
}
