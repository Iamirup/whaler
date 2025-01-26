package transaction

import (
	"log"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/domain/entity"
)

func isWithinLastAge(timeStr string, age int) bool {
	const layout = "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	return time.Since(t) <= time.Duration(age)*time.Second
}

func ConvertToTableRecord(tx map[string]interface{}) (entity.TableRecord, error) {
	record := entity.TableRecord{}

	if blockId, ok := tx["block_id"].(float64); ok {
		record.BlockId = int(blockId)
	}

	if outputTotal, ok := tx["output_total_usd"].(float64); ok {
		record.OutputTotalUsd = outputTotal
	} else if valueUsd, ok := tx["value_usd"].(float64); ok {
		record.OutputTotalUsd = valueUsd
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

func MeetsFilterCriteria(tx map[string]interface{}, minAmount float64, age int) bool {
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
