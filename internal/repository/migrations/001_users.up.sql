CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	username VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(72) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
