package ports

import (
	serr "github.com/Iamirup/whaler/backend/microservices/support/pkg/errors"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"
)

type (
	// TicketPersistencePort defines the methods for interacting with ticket data
	TicketPersistencePort interface {
		// CreateTicket adds a new ticket to the database
		CreateTicket(ticket *entity.Ticket) (entity.uuid, error)

		// GetMyTickets retrieves someone's tickets by its username and password
		GetMyTickets(userId entity.uuid) ([]entity.Ticket, string, error)

		// CreateReplyForTicket sends a reply to a ticket
		CreateReplyForTicket(ticketId entity.uuid, replyText string) error

		// CheckIfIsReplyForTheTicket checks if the ticket has already replied or not
		CheckIfIsReplyForTheTicket(ticketId entity.uuid) error

		// GetAllTickets retrieves all tickets in the system
		GetAllTickets() ([]entity.Ticket, string, error)
	}

	// TicketServicePort defines the methods for interacting with ticket services
	TicketServicePort interface {
		NewTicket(title, content string, userId entity.uuid, username string) (entity.uuid, *serr.ServiceError)
		MyTickets(userId entity.uuid) ([]entity.Ticket, *serr.ServiceError)
		ReplyToTicket(ticketId entity.uuid, replyText string) *serr.ServiceError
		AllTicket() ([]entity.Ticket, *serr.ServiceError)
	}
)
