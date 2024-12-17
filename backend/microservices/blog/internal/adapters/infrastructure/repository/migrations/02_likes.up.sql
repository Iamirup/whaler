CREATE TABLE IF NOT EXISTS likes(
	liker_id UUID NOT NULL,
    article_id UUID NOT NULL REFERENCES articles(id),
    PRIMARY KEY (liker_id, article_id)
);

