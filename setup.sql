CREATE TABLE users(
	id serial PRIMARY KEY,
	email VARCHAR (320) UNIQUE NOT NULL,
	hash VARCHAR (60) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO users(email, hash)
VALUES('john@example.com', 'first_hash'),
('adam@example.com', 'second_hash'),
('sophie@example.com', 'third_hash');