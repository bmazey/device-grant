package device

import "device-grant/pkg/oauth"

type Granter struct {
	Issuer oauth.SimpleIssuer
}

func NewGranter(issuer oauth.SimpleIssuer) Granter {
	return Granter{
		Issuer: issuer,
	}
}
