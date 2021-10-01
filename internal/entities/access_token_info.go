package entities

import (
	"time"
)

type AccessTokenInfo struct {
	UserId          string
	ExpiredDuration time.Duration
}
