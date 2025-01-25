package services

import (
	serr "github.com/Iamirup/whaler/backend/microservices/eventor/pkg/errors"
	"go.uber.org/zap"

	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/application/ports"
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
		transactions = allBitcoinTransactions
	case "ethereum":
		transactions = allEthereumTransactions
	case "dogecoin":
		transactions = allDogecoinTransactions
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
		if !meetsFilterCriteria(tx, minAmount, age) {
			continue
		}

		record, err := convertToTableRecord(tx)
		if err != nil {
			continue
		}

		records = append(records, record)
	}

	return records, nil
}

func convertToTableRecord(tx map[string]interface{}) (entity.TableRecord, error) {
	record := entity.TableRecord{}

	if blockId, ok := tx["block_id"].(float64); ok {
		record.BlockId = int(blockId)
	}

	if outputTotal, ok := tx["output_total_usd"].(float64); ok {
		record.OutputTotalUsd = outputTotal
	}

	if hash, ok := tx["hash"].(string); ok {
		record.Hash = hash
	}

	if timeStr, ok := tx["time"].(string); ok {
		record.Time = timeStr
	}

	if _, ok := tx["res"].(string); ok {
		record.Res = "OK"
	}

	if _, ok := tx["token"].(string); ok {
		record.Token = "USD Token"
	}

	if _, ok := tx["type"].(string); ok {
		record.Type = "Transfer"
	}

	return record, nil
}

func meetsFilterCriteria(tx map[string]interface{}, minAmount float64, age int) bool {
	outputTotalUSD, ok := tx["output_total_usd"].(float64)
	if !ok || outputTotalUSD <= minAmount {
		return false
	}

	timeStr, ok := tx["time"].(string)
	if !ok || !isWithinLastAge(timeStr, age) {
		return false
	}

	return true
}
