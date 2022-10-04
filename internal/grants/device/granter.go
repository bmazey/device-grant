package device

import (
	"device-grant/internal/data"
	"device-grant/pkg/oauth"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

const RUNES = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Granter struct {
	Issuer             oauth.SimpleIssuer
	ClientStore        data.ClientStore
	InteractionStore   data.InteractionStore
	TrustedDeviceStore data.TrustedDeviceStore
	CodeTTL            time.Duration
	UserCodeLength     int
}

func NewGranter(issuer oauth.SimpleIssuer, ttl time.Duration, length int) Granter {
	return Granter{
		Issuer:             issuer,
		ClientStore:        data.NewClientStore(),
		InteractionStore:   data.NewInteractionStore(),
		TrustedDeviceStore: data.NewTrustedDeviceStore(),
		CodeTTL:            ttl,
		UserCodeLength:     length,
	}
}

func (g *Granter) CreateInteraction(clientID string) {
	expires := time.Unix(time.Now().Unix()+int64(g.CodeTTL.Seconds()), 0)

	i := data.Interaction{
		ClientID:   clientID,
		DeviceCode: g.generateDeviceCode(),
		UserCode:   g.generateUserCode(),
		Expires:    expires,
	}

	g.InteractionStore.Add(i)
}

// AuthorizeDevice checks for an unexpired interaction by userCode and if one exists, trusts device
func (g *Granter) AuthorizeDevice(userCode string) error {
	// start by looking for a pre-existing interaction
	i, err := g.InteractionStore.Retrieve(userCode)
	if err != nil {
		return err
	}

	// interaction exists, clean it up after
	interaction := i.(data.Interaction)
	defer g.InteractionStore.Delete(interaction)

	// create device based on original interaction data
	d := data.TrustedDevice{
		Code:  interaction.DeviceCode,
		Owner: interaction.ClientID,
	}

	// add device to trusted store
	g.TrustedDeviceStore.AddDevice(d)

	return nil
}

func (g *Granter) generateDeviceCode() string {
	// we're just going to do a UUID here for the sake of simplicity
	return uuid.New().String()
}

func (g *Granter) generateUserCode() string {
	// this code has to be simple enough for a human to interact with
	runes := []rune(RUNES)

	code := make([]rune, g.UserCodeLength)
	for i := range code {
		code[i] = runes[rand.Intn(len(runes))]
	}
	return string(code)
}
