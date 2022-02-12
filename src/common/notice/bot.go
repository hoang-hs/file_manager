package notice

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

const requestTimeOut = 10 * time.Second

type Bot struct {
	Token      string
	BaseURL    string
	httpClient *fasthttp.Client
}

func NewBot(token string) *Bot {
	baseURL := BaseBotApiEndpointURL + token + "/"
	return &Bot{
		Token:      token,
		BaseURL:    baseURL,
		httpClient: &fasthttp.Client{},
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
	url := bot.BaseURL + SendMessage
	//api := fmt.Sprintf("%s?chat_id=%s&text=%s&parse_mode=%s", url, teleChan.ChatID, message, template.ParseMode)
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET") //default, with other method(POST, DELETE,...), pls creat new method
	query := req.URI().QueryArgs()
	query.Add("chat_id", teleChan.ChatID)
	query.Add("text", message)
	query.Add("parse_mode", template.ParseMode)
	req.SetRequestURI(url)
	response := fasthttp.AcquireResponse()
	err := bot.httpClient.DoTimeout(req, response, requestTimeOut)
	defer req.ConnectionClose()

	if err != nil {
		return Response{
			Code: 400,
			Err:  err,
		}
	}
	if response.StatusCode() != 200 {
		err = errors.New("response error with status code " + strconv.Itoa(response.StatusCode()))
	}
	return Response{
		Code: response.StatusCode(),
		Err:  err,
	}
}
