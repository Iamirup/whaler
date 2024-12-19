package repository

import (
	"errors"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/crypto"
	"go.uber.org/zap"
)

const QueryCreateArticle = `
INSERT INTO articles(title, content, author_id, author_username) VALUES($1, $2, $3, $4)
RETURNING id;`

func (r *articleRepository) CreateArticle(article *entity.Article) error {

	if len(article.Title) == 0 || len(article.Content) == 0 || len(article.AuthorUsername) == 0 || len(article.AuthorId) == 0 {
		return errors.New("insufficient information for article")
	}

	in := []any{article.Title, article.Content, article.AuthorId, article.AuthorUsername}
	out := []any{&article.ArticleId}
	if err := r.rdbms.QueryRow(QueryCreateArticle, in, out); err != nil {
		r.logger.Error("Error inserting article", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetAnArticle = `
SELECT 
    a.id, 
    a.title, 
    a.content, 
    a.author_id, 
    a.author_username, 
    a.date, 
    COUNT(l.article_id) AS likes_count
FROM 
    articles a
LEFT JOIN 
    likes l ON a.id = l.article_id
WHERE 
    a.id = $1
GROUP BY 
    a.id, a.title, a.content, a.author_id, a.author_username, a.date;`

func (r *articleRepository) GetAnArticle(articleId entity.UUID) (*entity.Article, error) {

	article := &entity.Article{ArticleId: articleId}

	in := []any{article.ArticleId}
	out := []any{
		&article.ArticleId,
		&article.Title,
		&article.Content,
		&article.AuthorId,
		&article.AuthorUsername,
		&article.Date,
		&article.Likes,
	}

	if err := r.rdbms.QueryRow(QueryGetAnArticle, in, out); err != nil {
		r.logger.Error("Error retrieving article", zap.Error(err))
		return &entity.Article{}, err
	}

	return article, nil
}

const QueryGetAllArticles = `
SELECT 
    a.id, 
    a.title, 
    a.content, 
    a.author_id, 
    a.author_username, 
    a.date, 
    COUNT(l.article_id) AS likes_count
FROM 
    articles a
LEFT JOIN 
    likes l ON a.id = l.article_id
WHERE
    date > $1
GROUP BY 
    a.id, a.title, a.content, a.author_id, a.author_username, a.date
ORDER BY 
    a.date ASC
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
			&articles[index].Title,
			&articles[index].Content,
			&articles[index].AuthorId,
			&articles[index].AuthorUsername,
			&articles[index].Date,
			&articles[index].Likes,
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
			&articles[index].Title,
			&articles[index].Content,
			&articles[index].AuthorId,
			&articles[index].AuthorUsername,
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

const QueryLikeArticle = `
INSERT INTO likes(liker_id, article_id) VALUES($1, $2)
RETURNING article_id;`

func (r *articleRepository) LikeArticle(articleId, likerId entity.UUID) error {

	var a string

	in := []any{likerId, articleId}
	out := []any{&a}
	if err := r.rdbms.QueryRow(QueryLikeArticle, in, out); err != nil {
		r.logger.Error("Error inserting new likes record", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetTopAuthors = `
SELECT 
    a.author_id,
    a.author_username,
    COUNT(l.article_id) AS total_likes
FROM 
    articles a
JOIN 
    likes l ON a.id = l.article_id
GROUP BY 
    a.author_id, a.author_username
ORDER BY 
    total_likes DESC
LIMIT 5;`

func (r *articleRepository) GetTopAuthors() ([]entity.TopAuthor, error) {

	limit := 5
	topAuthor := make([]entity.TopAuthor, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&topAuthor[index].AuthorId,
			&topAuthor[index].AuthorUsername,
			&topAuthor[index].Likes,
		}
	}

	in := []any{}
	if err := r.rdbms.Query(QueryGetTopAuthors, in, out); err != nil {
		r.logger.Error("Error retrieving top authors", zap.Error(err))
		return []entity.TopAuthor{}, err
	}

	return topAuthor, nil
}

const QueryGetPopularArticles = `
SELECT 
    a.id, 
    a.title, 
    a.content, 
    a.author_id, 
    a.author_username, 
    a.date, 
    COUNT(l.article_id) AS likes_count
FROM 
    articles a
LEFT JOIN 
    likes l ON a.id = l.article_id
GROUP BY 
    a.id, a.title, a.content, a.author_id, a.author_username, a.date
ORDER BY 
    likes_count DESC
LIMIT 5;`

func (r *articleRepository) GetPopularArticles() ([]entity.Article, error) {

	limit := 5
	articles := make([]entity.Article, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&articles[index].ArticleId,
			&articles[index].Title,
			&articles[index].Content,
			&articles[index].AuthorId,
			&articles[index].AuthorUsername,
			&articles[index].Date,
			&articles[index].Likes,
		}
	}

	in := []any{}
	if err := r.rdbms.Query(QueryGetPopularArticles, in, out); err != nil {
		r.logger.Error("Error retrieving popular articles", zap.Error(err))
		return []entity.Article{}, err
	}

	return articles, nil
}
