package entity

import "time"

type UUID string

type Ticket struct {
	TicketId  UUID      `json:"ticket_id"`
	UserId    UUID      `json:"user_id"`
	Username  string    `json:"username"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Date      time.Time `json:"date"`
	IsDone    bool      `json:"is_done"`
	ReplyText string    `json:"reply_text"`
	ReplyDate time.Time `json:"reply_date"`
}
