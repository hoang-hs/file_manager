package notice

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
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
	message := fmt.Sprintf("[%s] have message: \n %s", template.Job, template.Message)
	url := bot.BaseURL + SendMessage
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET") //default, with other method(POST, DELETE,...), pls creat new method
	query := req.URI().QueryArgs()
	query.Add("chat_id", teleChan.ChatID)
	query.Add("text", message)
	log.Printf("uri send to telegram: [%s]", req.URI().String())
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
