package entity

type uuid string

type Ticket struct {
	TicketId  uuid   `json:"ticket_id"`
	UserId    uuid   `json:"user_id"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Date      string `json:"date"`
	IsDone    bool   `json:"is_done"`
	ReplyText string `json:"reply_text"`
	ReplyDate string `json:"reply_date"`
}
