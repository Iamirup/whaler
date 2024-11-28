package ports

import (
	serr "github.com/Iamirup/whaler/backend/microservices/support/pkg/errors"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"
)

type (
	// TicketPersistencePort defines the methods for interacting with ticket data
	TicketPersistencePort interface {
		// CreateTicket adds a new ticket to the database
		CreateTicket(ticket *entity.Ticket) error

		// GetMyTickets retrieves someone's tickets by its username and password
		GetMyTickets(userId entity.UUID, encryptedCursor string, limit int) ([]entity.Ticket, string, error)

		// CreateReplyForTicket sends a reply to a ticket
		CreateReplyForTicket(ticketId entity.UUID, replyText string) error

		// CheckIfIsReplyForTheTicket checks if the ticket has already replied or not
		CheckIfIsReplyForTheTicket(ticketId entity.UUID) error

		// GetAllTickets retrieves all tickets in the system
		GetAllTickets(encryptedCursor string, limit int) ([]entity.Ticket, string, error)
	}

	// TicketServicePort defines the methods for interacting with ticket services
	TicketServicePort interface {
		NewTicket(title, content string, userId entity.UUID, username string) (entity.UUID, *serr.ServiceError)
		MyTickets(userId entity.UUID, encryptedCursor string, limit int) ([]entity.Ticket, string, *serr.ServiceError)
		ReplyToTicket(ticketId entity.UUID, replyText string) *serr.ServiceError
		AllTicket(encryptedCursor string, limit int) ([]entity.Ticket, string, *serr.ServiceError)
	}
)
