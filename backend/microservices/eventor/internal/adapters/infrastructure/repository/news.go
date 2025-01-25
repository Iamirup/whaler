package repository

import (
	"github.com/Iamirup/whaler/backend/microservices/eventor/internal/core/domain/entity"
)

const QueryCreateTableConfig = `
INSERT INTO tableConfig(title, content) VALUES($1, $2)
RETURNING id;`

func (r *tableConfigRepository) UpdateTableConfig(tableConfig *entity.TableConfig) error {

	// if len(tableConfig.Title) == 0 || len(tableConfig.Content) == 0 {
	// 	return errors.New("insufficient information for tableConfig")
	// }

	// in := []any{tableConfig.Title, tableConfig.Content}
	// out := []any{&tableConfig.TableConfigId}
	// if err := r.rdbms.QueryRow(QueryCreateTableConfig, in, out); err != nil {
	// 	r.logger.Error("Error inserting tableConfig", zap.Error(err))
	// 	return err
	// }

	return nil
}

const QueryGetTableConfig = `
SELECT *
FROM tableConfig
WHERE date > $1
ORDER BY date
FETCH NEXT $2 ROWS ONLY;`

func (r *tableConfigRepository) GetTableConfig(encryptedCursor string, limit int) ([]entity.TableConfig, string, error) {
	// var date time.Time

	// if limit < r.config.Limit.Min {
	// 	limit = r.config.Limit.Min
	// } else if limit > r.config.Limit.Max {
	// 	limit = r.config.Limit.Max
	// }

	// // decrypt cursor
	// if len(encryptedCursor) != 0 {
	// 	cursor, err := crypto.Decrypt(encryptedCursor, r.config.CursorSecret)
	// 	if err != nil {
	// 		return nil, "", err
	// 	}

	// 	date, err = time.Parse(time.RFC3339Nano, cursor)
	// 	if err != nil {
	// 		return nil, "", err
	// 	}
	// } else {
	// 	date = time.Unix(0, 0)
	// }

	// tableConfig := make([]entity.TableConfig, limit)
	// out := make([][]any, limit)

	// for index := 0; index < limit; index++ {
	// 	out[index] = []any{
	// 		&tableConfig[index].TableConfigId,
	// 		&tableConfig[index].Title,
	// 		&tableConfig[index].Content,
	// 		&tableConfig[index].Date,
	// 	}
	// }

	// in := []any{date, limit}
	// if err := r.rdbms.Query(QueryGetTableConfig, in, out); err != nil {
	// 	r.logger.Error("Error query tableConfig", zap.Error(err))
	// 	return nil, "", err
	// }

	// if len(tableConfig) == 0 {
	// 	return tableConfig, "", nil
	// }

	// var lastTableConfig entity.TableConfig

	// for index := limit - 1; index >= 0; index-- {
	// 	if tableConfig[index].TableConfigId != "" {
	// 		lastTableConfig = tableConfig[index]
	// 		break
	// 	} else {
	// 		tableConfig = tableConfig[:index]
	// 	}
	// }

	// if lastTableConfig.TableConfigId == "" {
	// 	return tableConfig, "", nil
	// }

	// cursor := lastTableConfig.Date.Format(time.RFC3339Nano)

	// // encrypt cursor
	// encryptedCursor, err := crypto.Encrypt(cursor, r.config.CursorSecret)
	// if err != nil {
	// 	return nil, "", err
	// }

	// return tableConfig, encryptedCursor, nil
	return nil, "", nil
}
