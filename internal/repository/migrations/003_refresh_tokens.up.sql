CREATE TABLE IF NOT EXISTS refresh_tokens(
	owner_id uuid PRIMARY KEY REFERENCES users(id),
    refresh_token TEXT UNIQUE NOT NULL
);
