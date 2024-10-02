package token

import "time"

type Config struct {
	PrivatePem             string        `koanf:"private_pem"`
	PublicPem              string        `koanf:"public_pem"`
	AccessTokenExpiration  time.Duration `koanf:"access_token_expiration"`
	RefreshTokenExpiration time.Duration `koanf:"refresh_token_expiration"`
}
