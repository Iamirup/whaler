package token

import "time"

type Config struct {
	PublicPem             string        `koanf:"public_pem"`
	AccessTokenExpiration time.Duration `koanf:"access_token_expiration"`
}
