package repository

import (
	"errors"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/magazine/pkg/crypto"
	"go.uber.org/zap"
)

const QueryCreateNews = `
INSERT INTO news(title, content) VALUES($1, $2)
RETURNING id;`

func (r *newsRepository) CreateNews(news *entity.News) error {

	if len(news.Title) == 0 || len(news.Content) == 0 {
		return errors.New("insufficient information for news")
	}

	in := []any{news.Title, news.Content}
	out := []any{&news.NewsId}
	if err := r.rdbms.QueryRow(QueryCreateNews, in, out); err != nil {
		r.logger.Error("Error inserting news", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetNews = `
SELECT *
FROM news
WHERE date > $1
ORDER BY date
FETCH NEXT $2 ROWS ONLY;`

func (r *newsRepository) GetNews(encryptedCursor string, limit int) ([]entity.News, string, error) {
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
			return nil, "", err
		}

		date, err = time.Parse(time.RFC3339Nano, cursor)
		if err != nil {
			return nil, "", err
		}
	} else {
		date = time.Unix(0, 0)
	}

	news := make([]entity.News, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&news[index].NewsId,
			&news[index].Title,
			&news[index].Content,
			&news[index].Date,
		}
	}

	in := []any{date, limit}
	if err := r.rdbms.Query(QueryGetNews, in, out); err != nil {
		r.logger.Error("Error query news", zap.Error(err))
		return nil, "", err
	}

	if len(news) == 0 {
		return news, "", nil
	}

	var lastNews entity.News

	for index := limit - 1; index >= 0; index-- {
		if news[index].NewsId != "" {
			lastNews = news[index]
			break
		} else {
			news = news[:index]
		}
	}

	if lastNews.NewsId == "" {
		return news, "", nil
	}

	cursor := lastNews.Date.Format(time.RFC3339Nano)

	// encrypt cursor
	encryptedCursor, err := crypto.Encrypt(cursor, r.config.CursorSecret)
	if err != nil {
		return nil, "", err
	}

	return news, encryptedCursor, nil
}
