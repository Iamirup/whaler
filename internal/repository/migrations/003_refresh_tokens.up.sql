CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS refresh_tokens(
	owner_id UUID PRIMARY KEY REFERENCES users(id),
    refresh_token TEXT UNIQUE NOT NULL
);
