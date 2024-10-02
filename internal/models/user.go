package models

type User struct {
	Id         uint64     `json:"Id"`
	Username   string     `json:"username"`
	Password   string     `json:"password,omitempty"`
	UserConfig UserConfig `json:"user_config,omitempty"`
	CreatedAt  string     `json:"created_at"`
}

func (c User) Marshal() *User {
	return &User{
		Id:         c.Id,
		Username:   c.Username,
		UserConfig: c.UserConfig,
		CreatedAt:  c.CreatedAt,
	}
}
