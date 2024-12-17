package entity

type TopAuthor struct {
	AuthorId       UUID   `json:"author_id"`
	AuthorUsername string `json:"author_username"`
	Likes          int    `json:"likes"`
}
