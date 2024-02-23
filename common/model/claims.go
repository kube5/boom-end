package model

import (
	"time"
)

type TokenDetails struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	AccessUuid   string        `json:"access_uuid"`
	RefreshUuid  string        `json:"refresh_uuid"`
	AtExpires    time.Duration `json:"at_expires"`
	RtExpires    time.Duration `json:"rt_expires"`
}

type AccessDetails struct {
	AccessUuid string
	UserId     string
}
