package models

type RefreshToken struct {
	Token   string `json:"refresh_token"`
	OwnerId uint64 `json:"id"`
}
