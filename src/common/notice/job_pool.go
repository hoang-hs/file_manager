package notice

import (
	"log"
	"net/http"
	"time"
)

type JobPool struct {
	Sender                Sender
	Queue                 chan Package
	GapTimeRequest        time.Duration
	GapTimeTooManyRequest time.Duration
}

func NewJobPool(b Sender, sizeQueue int, gap, gap429 time.Duration) *JobPool {
	queue := make(chan Package, sizeQueue)
	if gap < 0 {
		gap = GapTimeRequest
	}
	if gap429 < 0 {
		gap429 = GapTimeTooManyRequest
	}
	return &JobPool{
		Sender:                b,
		Queue:                 queue,
		GapTimeRequest:        gap,
		GapTimeTooManyRequest: gap429,
	}
}

func (jp *JobPool) AddPackage(p Package) {
	jp.Queue <- p
}

func (jp *JobPool) Start() {
	go func() {
		for {
			limitRequest := true
			p, ok := <-jp.Queue
			if !ok {
				return
			}
			for limitRequest {
				resp := jp.Sender.Send(p)
				switch {
				case resp.Code == http.StatusOK:
					limitRequest = false
				case resp.Code == http.StatusTooManyRequests:
					limitRequest = true
					time.Sleep(jp.GapTimeTooManyRequest)
				default:
					log.Printf("Send message to telegram failed: code-%d, error-%s\n", resp.Code, resp.Err.Error())
					limitRequest = false
				}
			}
			// wait to avoid limit between two request
			time.Sleep(jp.GapTimeRequest)
		}
	}()
}
