package services

import (
	"net/http"

	serr "github.com/Iamirup/whaler/backend/microservices/support/pkg/errors"
	"go.uber.org/zap"

	api "github.com/Iamirup/whaler/backend/microservices/support/internal/adapters/interfaces/rest/dto"

	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/support/internal/core/domain/entity"
)

type TicketApplicationService struct {
	domainService ports.TicketServicePort
	logger        *zap.Logger
}

func NewTicketApplicationService(domainService ports.TicketServicePort, logger *zap.Logger) *TicketApplicationService {
	return &TicketApplicationService{
		domainService: domainService,
		logger:        logger,
	}
}

func (s *TicketApplicationService) NewTicket(request *api.NewTicketRequest, userId entity.uuid, username string) (entity.uuid, *serr.ServiceError) {
	if err := request.Validate(); err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		return "", &serr.ServiceError{Message: err.Error(), StatusCode: http.StatusBadRequest}
	}

	return s.domainService.NewTicket(request.Title, request.Content, userId, username)
}

func (s *TicketApplicationService) MyTickets(userId entity.uuid) ([]entity.Ticket, *serr.ServiceError) {
	return s.domainService.MyTickets(userId)
}

func (s *TicketApplicationService) ReplyToTicket(request *api.ReplyToTicketRequest) *serr.ServiceError {
	if err := request.Validate(); err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		return "", &serr.ServiceError{Message: err.Error(), StatusCode: http.StatusBadRequest}
	}

	return s.domainService.ReplyToTicket(request.TicketId, request.ReplyText)
}

func (s *TicketApplicationService) AllTicket() ([]entity.Ticket, *serr.ServiceError) {
	return s.domainService.AllTicket()
}
