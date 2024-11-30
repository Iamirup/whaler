package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/crypto"
	"go.uber.org/zap"
)

const QueryCreateArticle = `
INSERT INTO articles(owner_id, username, title, content) VALUES($1, $2, $3, $4)
RETURNING id;`

func (r *articleRepository) CreateArticle(article *entity.Article) error {

	if len(article.Title) == 0 || len(article.Content) == 0 || len(article.Username) == 0 || len(article.UserId) == 0 {
		return errors.New("insufficient information for article")
	}

	in := []any{article.UserId, article.Username, article.Title, article.Content}
	out := []any{&article.ArticleId}
	if err := r.rdbms.QueryRow(QueryCreateArticle, in, out); err != nil {
		r.logger.Error("Error inserting article", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetMyArticles = `
SELECT *
FROM articles
WHERE 
    owner_id=$1 AND
    date > $2
ORDER BY date
FETCH NEXT $3 ROWS ONLY;`

func (r *articleRepository) GetMyArticles(userId entity.UUID, encryptedCursor string, limit int) ([]entity.Article, string, error) {
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

	articles := make([]entity.Article, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&articles[index].ArticleId,
			&articles[index].UserId,
			&articles[index].Username,
			&articles[index].Title,
			&articles[index].Content,
			&articles[index].Date,
			&articles[index].IsDone,
			&articles[index].ReplyText,
			&articles[index].ReplyDate,
		}
	}

	in := []any{userId, date, limit}
	if err := r.rdbms.Query(QueryGetMyArticles, in, out); err != nil {
		r.logger.Error("Error query articles", zap.Error(err))
		return nil, "", err
	}

	if len(articles) == 0 {
		return articles, "", nil
	}

	var lastArticle entity.Article

	for index := limit - 1; index >= 0; index-- {
		if articles[index].ArticleId != "" {
			lastArticle = articles[index]
			break
		} else {
			articles = articles[:index]
		}
	}

	if lastArticle.ArticleId == "" {
		return articles, "", nil
	}

	cursor := lastArticle.Date.Format(time.RFC3339Nano)

	// encrypt cursor
	encryptedCursor, err := crypto.Encrypt(cursor, r.config.CursorSecret)
	if err != nil {
		return nil, "", err
	}

	return articles, encryptedCursor, nil
}

const QueryCreateReplyForArticle = `
UPDATE articles
SET is_done=true, reply_text=$1, reply_date=$2
WHERE id=$3`

func (r *articleRepository) CreateReplyForArticle(articleId entity.UUID, replyText string) error {

	in := []interface{}{replyText, time.Now(), articleId}
	if err := r.rdbms.Execute(QueryCreateReplyForArticle, in); err != nil {
		r.logger.Error("Error inserting new reply", zap.Error(err))
		return err
	}

	return nil
}

const QueryCheckIfIsReplyForTheArticle = `
SELECT id
FROM articles
WHERE id=$1 and is_done=false`

func (r *articleRepository) CheckIfIsReplyForTheArticle(articleId entity.UUID) error {

	article := &entity.Article{ArticleId: articleId}

	in := []interface{}{articleId}
	out := []interface{}{&article.ArticleId}
	if err := r.rdbms.QueryRow(QueryCheckIfIsReplyForTheArticle, in, out); err != nil {
		r.logger.Error("Error finding article by articleId", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetAllArticles = `
SELECT *
FROM articles
WHERE date > $1
ORDER BY date
FETCH NEXT $2 ROWS ONLY;`

func (r *articleRepository) GetAllArticles(encryptedCursor string, limit int) ([]entity.Article, string, error) {

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
			return nil, "", err
		}

		fmt.Println("cursor: ", cursor)

		date, err = time.Parse(time.RFC3339Nano, cursor)
		if err != nil {
			return nil, "", err
		}
	} else {
		date = time.Unix(0, 0)
	}

	fmt.Println("date: ", date)

	articles := make([]entity.Article, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&articles[index].ArticleId,
			&articles[index].UserId,
			&articles[index].Username,
			&articles[index].Title,
			&articles[index].Content,
			&articles[index].Date,
			&articles[index].IsDone,
			&articles[index].ReplyText,
			&articles[index].ReplyDate,
		}
	}

	in := []any{date, limit}
	if err := r.rdbms.Query(QueryGetAllArticles, in, out); err != nil {
		r.logger.Error("Error query articles", zap.Error(err))
		return nil, "", err
	}

	if len(articles) == 0 {
		return articles, "", nil
	}

	var lastArticle entity.Article

	for index := limit - 1; index >= 0; index-- {
		if articles[index].ArticleId != "" {
			lastArticle = articles[index]
			break
		} else {
			articles = articles[:index]
		}
	}

	if lastArticle.ArticleId == "" {
		return articles, "", nil
	}

	fmt.Println("lastArticle: ", lastArticle)

	cursor := lastArticle.Date.Format(time.RFC3339Nano)

	fmt.Println("cursor: ", cursor)

	// encrypt cursor
	encryptedCursor, err := crypto.Encrypt(cursor, r.config.CursorSecret)
	if err != nil {
		return nil, "", err
	}

	return articles, encryptedCursor, nil
}
