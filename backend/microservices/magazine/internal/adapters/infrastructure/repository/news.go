package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/magazine/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/magazine/pkg/crypto"
	"go.uber.org/zap"
)

const QueryCreateNews = `
INSERT INTO news(owner_id, username, title, content) VALUES($1, $2, $3, $4)
RETURNING id;`

func (r *newsRepository) CreateNews(news *entity.News) error {

	if len(news.Title) == 0 || len(news.Content) == 0 || len(news.Username) == 0 || len(news.UserId) == 0 {
		return errors.New("insufficient information for news")
	}

	in := []any{news.UserId, news.Username, news.Title, news.Content}
	out := []any{&news.NewsId}
	if err := r.rdbms.QueryRow(QueryCreateNews, in, out); err != nil {
		r.logger.Error("Error inserting news", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetMyNews = `
SELECT *
FROM news
WHERE 
    owner_id=$1 AND
    date > $2
ORDER BY date
FETCH NEXT $3 ROWS ONLY;`

func (r *newsRepository) GetMyNews(userId entity.UUID, encryptedCursor string, limit int) ([]entity.News, string, error) {
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

		date, err = time.Parse(time.RFC3339Nano, cursor)
		if err != nil {
			panic(err)
		}
	} else {
		date = time.Unix(0, 0)
	}

	news := make([]entity.News, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&news[index].NewsId,
			&news[index].UserId,
			&news[index].Username,
			&news[index].Title,
			&news[index].Content,
			&news[index].Date,
			&news[index].IsDone,
			&news[index].ReplyText,
			&news[index].ReplyDate,
		}
	}

	in := []any{userId, date, limit}
	if err := r.rdbms.Query(QueryGetMyNews, in, out); err != nil {
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
		panic(err)
	}

	return news, encryptedCursor, nil
}

const QueryCreateReplyForNews = `
UPDATE news
SET is_done=true, reply_text=$1, reply_date=$2
WHERE id=$3`

func (r *newsRepository) CreateReplyForNews(newsId entity.UUID, replyText string) error {

	in := []interface{}{replyText, time.Now(), newsId}
	if err := r.rdbms.Execute(QueryCreateReplyForNews, in); err != nil {
		r.logger.Error("Error inserting new reply", zap.Error(err))
		return err
	}

	return nil
}

const QueryCheckIfIsReplyForTheNews = `
SELECT id
FROM news
WHERE id=$1 and is_done=false`

func (r *newsRepository) CheckIfIsReplyForTheNews(newsId entity.UUID) error {

	news := &entity.News{NewsId: newsId}

	in := []interface{}{newsId}
	out := []interface{}{&news.NewsId}
	if err := r.rdbms.QueryRow(QueryCheckIfIsReplyForTheNews, in, out); err != nil {
		r.logger.Error("Error finding news by newsId", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetAllNews = `
SELECT *
FROM news
WHERE date > $1
ORDER BY date
FETCH NEXT $2 ROWS ONLY;`

func (r *newsRepository) GetAllNews(encryptedCursor string, limit int) ([]entity.News, string, error) {

	var date time.Time

	fmt.Println("encryptedCursor: ", encryptedCursor)
	fmt.Println("limit: ", limit)

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

		fmt.Println("cursor: ", cursor)

		date, err = time.Parse(time.RFC3339Nano, cursor)
		if err != nil {
			panic(err)
		}
	} else {
		date = time.Unix(0, 0)
	}

	fmt.Println("date: ", date)

	news := make([]entity.News, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&news[index].NewsId,
			&news[index].UserId,
			&news[index].Username,
			&news[index].Title,
			&news[index].Content,
			&news[index].Date,
			&news[index].IsDone,
			&news[index].ReplyText,
			&news[index].ReplyDate,
		}
	}

	in := []any{date, limit}
	if err := r.rdbms.Query(QueryGetAllNews, in, out); err != nil {
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

	fmt.Println("lastNews: ", lastNews)

	cursor := lastNews.Date.Format(time.RFC3339Nano)

	fmt.Println("cursor: ", cursor)

	// encrypt cursor
	encryptedCursor, err := crypto.Encrypt(cursor, r.config.CursorSecret)
	if err != nil {
		panic(err)
	}

	return news, encryptedCursor, nil
}
