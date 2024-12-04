package entity

import "time"

type UUID string

type Article struct {
	ArticleId      UUID      `json:"article_id"`
	UrlPath        string    `json:"url_path"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	AuthorId       UUID      `json:"author_id"`
	AuthorUsername string    `json:"author_username"`
	Likes          int       `json:"likes"`
	Date           time.Time `json:"date"`
}
