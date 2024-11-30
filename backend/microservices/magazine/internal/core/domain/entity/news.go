package entity

import "time"

type UUID string

type News struct {
	NewsId  UUID      `json:"news_id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}
