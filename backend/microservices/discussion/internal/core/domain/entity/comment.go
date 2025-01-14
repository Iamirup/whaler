package entity

import "time"

type UUID string

type Comment struct {
	CommentId int64     `json:"comment_id"`
	TopicId   UUID      `json:"user_id"`
	Currency  string    `json:"currency"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	Date      time.Time `json:"date"`
}
