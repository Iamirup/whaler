package models

type RefreshToken struct {
	Token   string `json:"refresh_token"`
	OwnerId string `json:"id"`
}
