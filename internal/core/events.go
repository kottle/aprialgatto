package core

import "github.com/aprialgatto/internal/utils/events"

var VerifyAuthToken = events.Topic("VerifyAuthToken")

type VerifyAuthTokenEvt struct {
	TokenID string
}

func NewVerifyAuthToken(tokenid string) *VerifyAuthTokenEvt {
	return &VerifyAuthTokenEvt{TokenID: tokenid}
}
