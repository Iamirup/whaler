package repository

import (
	"errors"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/crypto"
	"go.uber.org/zap"
)

const QueryCreateArticle = `
INSERT INTO articles(url_path, title, content, author_id, author_username) VALUES($1, $2, $3, $4, $5)
RETURNING id;`

func (r *articleRepository) CreateArticle(article *entity.Article) error {

	if len(article.UrlPath) == 0 || len(article.Title) == 0 || len(article.Content) == 0 || len(article.AuthorUsername) == 0 || len(article.AuthorId) == 0 {
		return errors.New("insufficient information for article")
	}

	in := []any{article.UrlPath, article.Title, article.Content, article.AuthorId, article.AuthorUsername}
	out := []any{&article.ArticleId}
	if err := r.rdbms.QueryRow(QueryCreateArticle, in, out); err != nil {
		r.logger.Error("Error inserting article", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetAnArticle = `
SELECT *
FROM articles
WHERE url_path=$1`

func (r *articleRepository) GetAnArticle(urlPath string) (*entity.Article, error) {

	article := &entity.Article{UrlPath: urlPath}

	in := []any{article.UrlPath}
	out := []any{
		&article.ArticleId,
		&article.UrlPath,
		&article.Title,
		&article.Content,
		&article.AuthorId,
		&article.AuthorUsername,
		&article.Likes,
		&article.Date,
	}

	if err := r.rdbms.QueryRow(QueryGetAnArticle, in, out); err != nil {
		r.logger.Error("Error retrieving article", zap.Error(err))
		return &entity.Article{}, err
	}

	return article, nil
}

const QueryGetAllArticles = `
SELECT *
FROM articles
WHERE date > $1
ORDER BY date
FETCH NEXT $2 ROWS ONLY;`

func (r *articleRepository) GetAllArticles(encryptedCursor string, limit int) ([]entity.Article, string, error) {
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
			&articles[index].UrlPath,
			&articles[index].Title,
			&articles[index].Content,
			&articles[index].AuthorId,
			&articles[index].AuthorUsername,
			&articles[index].Likes,
			&articles[index].Date,
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

	cursor := lastArticle.Date.Format(time.RFC3339Nano)

	// encrypt cursor
	encryptedCursor, err := crypto.Encrypt(cursor, r.config.CursorSecret)
	if err != nil {
		return nil, "", err
	}

	return articles, encryptedCursor, nil
}

const QueryGetMyArticles = `
SELECT *
FROM articles
WHERE date > $1
AND author_id = $2
ORDER BY date
FETCH NEXT $3 ROWS ONLY;`

func (r *articleRepository) GetMyArticles(encryptedCursor string, limit int, authorId entity.UUID) ([]entity.Article, string, error) {
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
			&articles[index].UrlPath,
			&articles[index].Title,
			&articles[index].Content,
			&articles[index].AuthorId,
			&articles[index].AuthorUsername,
			&articles[index].Likes,
			&articles[index].Date,
		}
	}

	in := []any{date, authorId, limit}
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

const QueryUpdateArticleTitle = `
UPDATE articles
SET title=$1
WHERE id=$2`

func (r *articleRepository) UpdateArticleTitle(articleId entity.UUID, title string) error {

	in := []interface{}{title, articleId}
	if err := r.rdbms.Execute(QueryUpdateArticleTitle, in); err != nil {
		r.logger.Error("Error updating article title", zap.Error(err))
		return err
	}

	return nil
}

const QueryUpdateArticleContent = `
UPDATE articles
SET content=$1
WHERE id=$2`

func (r *articleRepository) UpdateArticleContent(articleId entity.UUID, content string) error {

	in := []interface{}{content, articleId}
	if err := r.rdbms.Execute(QueryUpdateArticleContent, in); err != nil {
		r.logger.Error("Error updating article content", zap.Error(err))
		return err
	}

	return nil
}

const QueryRemoveArticle = `
DELETE FROM articles
WHERE id=$1`

func (r *articleRepository) RemoveArticle(articleId entity.UUID) error {

	in := []interface{}{articleId}
	if err := r.rdbms.Execute(QueryRemoveArticle, in); err != nil {
		r.logger.Error("Error finding article by articleId", zap.Error(err))
		return err
	}

	return nil
}

const CheckIfIsAuthorById = `
SELECT author_id
FROM articles
WHERE id=$1`

func (r *articleRepository) CheckIfIsAuthorById(articleId, authorId entity.UUID) error {

	article := &entity.Article{ArticleId: articleId}

	in := []interface{}{articleId}
	out := []interface{}{&article.AuthorId}
	if err := r.rdbms.QueryRow(CheckIfIsAuthorById, in, out); err != nil {
		r.logger.Error("Error finding article by articleId", zap.Error(err))
		return err
	}

	if authorId != article.AuthorId {
		r.logger.Error("you are not the author of this article!")
		return errors.New("you are not the author of this article")
	}

	return nil
}
