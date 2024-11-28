CREATE TABLE IF NOT EXISTS tickets(
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	owner_id UUID NOT NULL,
	username VARCHAR(32),
	title VARCHAR(80) NOT NULL,
	content TEXT NOT NULL,
	date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	is_done BOOLEAN DEFAULT false,
	reply_text TEXT,
	reply_date TIMESTAMP
);