package device

import (
	"device-grant/pkg/oauth"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

const RUNES = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Granter struct {
	Issuer         oauth.SimpleIssuer
	CodeTTL        time.Duration
	UserCodeLength int
}

func NewGranter(issuer oauth.SimpleIssuer, ttl time.Duration, length int) Granter {
	return Granter{
		Issuer:         issuer,
		CodeTTL:        ttl,
		UserCodeLength: length,
	}
}

func (g *Granter) GenerateDeviceCode() string {
	// we're just going to do a UUID here for the sake of simplicity
	return uuid.New().String()
}

func (g *Granter) GenerateUserCode() string {
	// this code has to be simple enough for a human to interact with
	runes := []rune(RUNES)

	code := make([]rune, g.UserCodeLength)
	for i := range code {
		code[i] = runes[rand.Intn(len(runes))]
	}
	return string(code)
}
