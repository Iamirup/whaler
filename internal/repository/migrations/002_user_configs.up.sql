CREATE TABLE IF NOT EXISTS user_configs(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	phones VARCHAR(15)[] NOT NULL,
	description VARCHAR(255),
	user_id INTEGER REFERENCES users (id)
);
