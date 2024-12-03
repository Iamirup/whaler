package services

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/support/pkg/token"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"

	serr "github.com/Iamirup/whaler/backend/microservices/support/pkg/errors"
	"go.uber.org/zap"
)

type TicketService struct {
	ticketPersistencePort ports.TicketPersistencePort
	logger                *zap.Logger
	token                 token.Token
}

func NewTicketService(
	ticketPersistencePort ports.TicketPersistencePort,
	logger *zap.Logger, token token.Token) ports.TicketServicePort {

	return &TicketService{
		ticketPersistencePort: ticketPersistencePort,
		logger:                logger,
		token:                 token,
	}
}

func (s *TicketService) NewTicket(title, content string, userId entity.UUID, username string) (entity.UUID, *serr.ServiceError) {

	ticketEntity := &entity.Ticket{
		UserId:   userId,
		Username: username,
		Title:    title,
		Content:  content,
	}

	err := s.ticketPersistencePort.CreateTicket(ticketEntity)
	if err != nil {
		s.logger.Error("Error happened while creating the ticket", zap.Error(err))
		return "", &serr.ServiceError{Message: "Error happened while creating the ticket", StatusCode: http.StatusInternalServerError}
	} else if ticketEntity.TicketId == "" {
		s.logger.Error("Error invalid ticket id created", zap.Any("ticket", ticketEntity))
		return "", &serr.ServiceError{Message: "Error invalid ticket id created", StatusCode: http.StatusInternalServerError}
	}

	return ticketEntity.TicketId, nil
}

func (s *TicketService) MyTickets(userId entity.UUID, encryptedCursor string, limit int) ([]entity.Ticket, string, *serr.ServiceError) {

	tickets, newEncryptedCursor, err := s.ticketPersistencePort.GetMyTickets(userId, encryptedCursor, limit)
	if err != nil {
		s.logger.Error("Something went wrong in retrieving tickets", zap.Error(err))
		return nil, "", &serr.ServiceError{Message: "Something went wrong in retrieving tickets", StatusCode: http.StatusInternalServerError}
	}

	return tickets, newEncryptedCursor, nil
}

func (s *TicketService) ReplyToTicket(ticketId entity.UUID, replyText string) *serr.ServiceError {

	if err := s.ticketPersistencePort.CheckIfIsReplyForTheTicket(ticketId); err != nil {
		s.logger.Error("The ticket is already replied by an admin", zap.Error(err))
		return &serr.ServiceError{Message: "The ticket is already replied by an admin", StatusCode: http.StatusConflict}
	}

	if err := s.ticketPersistencePort.CreateReplyForTicket(ticketId, replyText); err != nil {
		s.logger.Error("Error creating a reply to the ticket", zap.Error(err))
		return &serr.ServiceError{Message: "Error creating a reply to the ticket", StatusCode: http.StatusInternalServerError}
	}

	return nil
}

func (s *TicketService) AllTicket(encryptedCursor string, limit int) ([]entity.Ticket, string, *serr.ServiceError) {

	tickets, newEncryptedCursor, err := s.ticketPersistencePort.GetAllTickets(encryptedCursor, limit)
	if err != nil {
		s.logger.Error("Something went wrong in retrieving tickets", zap.Error(err))
		return nil, "", &serr.ServiceError{Message: "Something went wrong in retrieving tickets", StatusCode: http.StatusInternalServerError}
	}

	return tickets, newEncryptedCursor, nil
}
