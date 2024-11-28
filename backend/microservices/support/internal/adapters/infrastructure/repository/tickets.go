package repository

import (
	"errors"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/support/pkg/crypto"
	"go.uber.org/zap"
)

const QueryCreateTicket = `
INSERT INTO tickets(owner_id, username, title, content) VALUES($1, $2, $3, $4)
RETURNING id;`

func (r *ticketRepository) CreateTicket(ticket *entity.Ticket) error {

	if len(ticket.Title) == 0 || len(ticket.Content) == 0 || len(ticket.Username) == 0 || len(ticket.UserId) == 0 {
		return errors.New("insufficient information for ticket")
	}

	in := []any{ticket.UserId, ticket.Username, ticket.Title, ticket.Content}
	out := []any{&ticket.TicketId}
	if err := r.rdbms.QueryRow(QueryCreateTicket, in, out); err != nil {
		r.logger.Error("Error inserting ticket", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetMyTickets = `
SELECT *
FROM tickets
WHERE 
    owner_id=$1 AND
    date > $2
ORDER BY date
FETCH NEXT $3 ROWS ONLY;`

func (r *ticketRepository) GetMyTickets(userId entity.UUID, encryptedCursor string, limit int) ([]entity.Ticket, string, error) {
	var date time.Time

	if limit < r.config.Limit.Min {
		limit = r.config.Limit.Min
	} else if limit > r.config.Limit.Max {
		limit = r.config.Limit.Max
	}

	// decrypt cursor
	if len(encryptedCursor) != 0 {
		cursor, err := crypto.Decrypt(encryptedCursor, r.config.CursorSecret)
		if err != nil {
			panic(err)
		}

		date, err = time.Parse(time.RFC3339, cursor)
		if err != nil {
			panic(err)
		}
	} else {
		date = time.Unix(0, 0)
	}

	tickets := make([]entity.Ticket, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&tickets[index].TicketId,
			&tickets[index].UserId,
			&tickets[index].Username,
			&tickets[index].Title,
			&tickets[index].Content,
			&tickets[index].Date,
			&tickets[index].IsDone,
			&tickets[index].ReplyText,
			&tickets[index].ReplyDate,
		}
	}

	in := []any{userId, date, limit}
	if err := r.rdbms.Query(QueryGetMyTickets, in, out); err != nil {
		r.logger.Error("Error query tickets", zap.Error(err))
		return nil, "", err
	}

	if len(tickets) == 0 {
		return tickets, "", nil
	}

	var lastTicket entity.Ticket

	for index := limit - 1; index >= 0; index-- {
		if tickets[index].TicketId != "" {
			lastTicket = tickets[index]
			break
		} else {
			tickets = tickets[:index]
		}
	}

	if lastTicket.TicketId == "" {
		return tickets, "", nil
	}

	cursor := lastTicket.Date.Format(time.RFC3339)

	// encrypt cursor
	encryptedCursor, err := crypto.Encrypt(cursor, r.config.CursorSecret)
	if err != nil {
		panic(err)
	}

	return tickets, encryptedCursor, nil
}

const QueryCreateReplyForTicket = `
UPDATE tickets
SET is_done=true, reply_text=$1, reply_date=$2
WHERE id=$3`

func (r *ticketRepository) CreateReplyForTicket(ticketId entity.UUID, replyText string) error {

	in := []interface{}{replyText, time.Now(), ticketId}
	if err := r.rdbms.Execute(QueryCreateReplyForTicket, in); err != nil {
		r.logger.Error("Error inserting new reply", zap.Error(err))
		return err
	}

	return nil
}

const QueryCheckIfIsReplyForTheTicket = `
SELECT id
FROM tickets
WHERE id=$1 and is_done=false`

func (r *ticketRepository) CheckIfIsReplyForTheTicket(ticketId entity.UUID) error {

	ticket := &entity.Ticket{TicketId: ticketId}

	in := []interface{}{ticketId}
	out := []interface{}{&ticket.TicketId}
	if err := r.rdbms.QueryRow(QueryCheckIfIsReplyForTheTicket, in, out); err != nil {
		r.logger.Error("Error finding ticket by ticketId", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetAllTickets = `
SELECT *
FROM tickets
WHERE date > $1
ORDER BY date
FETCH NEXT $2 ROWS ONLY;`

func (r *ticketRepository) GetAllTickets(encryptedCursor string, limit int) ([]entity.Ticket, string, error) {

	var date time.Time

	if limit < r.config.Limit.Min {
		limit = r.config.Limit.Min
	} else if limit > r.config.Limit.Max {
		limit = r.config.Limit.Max
	}

	// decrypt cursor
	if len(encryptedCursor) != 0 {
		cursor, err := crypto.Decrypt(encryptedCursor, r.config.CursorSecret)
		if err != nil {
			panic(err)
		}

		date, err = time.Parse(time.RFC3339, cursor)
		if err != nil {
			panic(err)
		}
	} else {
		date = time.Unix(0, 0)
	}

	tickets := make([]entity.Ticket, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&tickets[index].TicketId,
			&tickets[index].UserId,
			&tickets[index].Username,
			&tickets[index].Title,
			&tickets[index].Content,
			&tickets[index].Date,
			&tickets[index].IsDone,
			&tickets[index].ReplyText,
			&tickets[index].ReplyDate,
		}
	}

	in := []any{date, limit}
	if err := r.rdbms.Query(QueryGetMyTickets, in, out); err != nil {
		r.logger.Error("Error query tickets", zap.Error(err))
		return nil, "", err
	}

	if len(tickets) == 0 {
		return tickets, "", nil
	}

	var lastTicket entity.Ticket

	for index := limit - 1; index >= 0; index-- {
		if tickets[index].TicketId != "" {
			lastTicket = tickets[index]
			break
		} else {
			tickets = tickets[:index]
		}
	}

	if lastTicket.TicketId == "" {
		return tickets, "", nil
	}

	cursor := lastTicket.Date.Format(time.RFC3339)

	// encrypt cursor
	encryptedCursor, err := crypto.Encrypt(cursor, r.config.CursorSecret)
	if err != nil {
		panic(err)
	}

	return tickets, encryptedCursor, nil
}
