package token

import (
	"time"
)

type AuthorizationCode struct {
	Value     string
	ExpiredAt time.Time
}
