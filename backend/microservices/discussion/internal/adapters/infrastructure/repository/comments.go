package repository

import (
	"errors"
	"time"

	"github.com/Iamirup/whaler/backend/microservices/discussion/internal/core/domain/entity"
	"github.com/Iamirup/whaler/backend/microservices/discussion/pkg/crypto"
	"go.uber.org/zap"
)

const QueryCreateComment = `
INSERT INTO comments(topic_id, username, text) VALUES(
(SELECT id FROM topics WHERE currency = $1), 
$2, $3)
RETURNING id;`

func (r *commentRepository) AddComment(comment *entity.Comment) error {

	if len(comment.Text) == 0 || len(comment.Currency) == 0 || len(comment.Username) == 0 {
		return errors.New("insufficient information for comment")
	}

	in := []any{comment.Currency, comment.Username, comment.Text}
	out := []any{&comment.CommentId}
	if err := r.rdbms.QueryRow(QueryCreateComment, in, out); err != nil {
		r.logger.Error("Error inserting comment", zap.Error(err))
		return err
	}

	return nil
}

const QueryGetMyComments = `
SELECT *
FROM comments
WHERE 
    topic_id=(SELECT id FROM topics WHERE currency = $1) AND
    date > $2
ORDER BY date
FETCH NEXT $3 ROWS ONLY;`

func (r *commentRepository) GetComments(currencyTopic, encryptedCursor string, limit int) ([]entity.Comment, string, error) {
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

	comments := make([]entity.Comment, limit)
	out := make([][]any, limit)

	for index := 0; index < limit; index++ {
		out[index] = []any{
			&comments[index].CommentId,
			&comments[index].TopicId,
			&comments[index].Username,
			&comments[index].Text,
			&comments[index].Date,
		}
	}

	in := []any{currencyTopic, date, limit}
	if err := r.rdbms.Query(QueryGetMyComments, in, out); err != nil {
		r.logger.Error("Error query comments", zap.Error(err))
		return nil, "", err
	}

	if len(comments) == 0 {
		return comments, "", nil
	}

	var lastComment entity.Comment

	for index := limit - 1; index >= 0; index-- {
		if comments[index].CommentId != "" {
			lastComment = comments[index]
			break
		} else {
			comments = comments[:index]
		}
	}

	if lastComment.CommentId == "" {
		return comments, "", nil
	}

	cursor := lastComment.Date.Format(time.RFC3339Nano)

	// encrypt cursor
	encryptedCursor, err := crypto.Encrypt(cursor, r.config.CursorSecret)
	if err != nil {
		return nil, "", err
	}

	return comments, encryptedCursor, nil
}
