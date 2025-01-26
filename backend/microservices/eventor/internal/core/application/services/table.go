package services

import (
	serr "github.com/Iamirup/whaler/backend/microservices/eventor/pkg/errors"
	"go.uber.org/zap"

	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/application/ports"
	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/application/services/transaction"
	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/domain/entity"
)

type TableConfigApplicationService struct {
	domainService ports.TableConfigServicePort
	logger        *zap.Logger
}

func NewTableConfigApplicationService(domainService ports.TableConfigServicePort, logger *zap.Logger) *TableConfigApplicationService {
	return &TableConfigApplicationService{
		domainService: domainService,
		logger:        logger,
	}
}

func (s *TableConfigApplicationService) SeeTable(cryptocurrency string, minAmount float64, age int, encryptedCursor string, limit int) ([]entity.TableRecord, string, *serr.ServiceError) {
	var transactions []map[string]interface{}

	switch cryptocurrency {
	case "bitcoin":
		transactions = transaction.AllBitcoinTransactions
	case "ethereum":
		transactions = transaction.AllEthereumTransactions
	case "dogecoin":
		transactions = transaction.AllDogecoinTransactions
	default:
		return nil, "", nil
	}

	records, err := s.filterAndConvertTransactions(transactions, minAmount, age)
	if err != nil {
		return nil, "", nil
	}

	return records, "", nil
}

func (s *TableConfigApplicationService) filterAndConvertTransactions(transactions []map[string]interface{}, minAmount float64, age int) ([]entity.TableRecord, error) {
	var records []entity.TableRecord

	for _, tx := range transactions {
		if !transaction.MeetsFilterCriteria(tx, minAmount, age) {
			continue
		}

		record, err := transaction.ConvertToTableRecord(tx)
		if err != nil {
			continue
		}

		records = append(records, record)
	}

	return records, nil
}
