CREATE TABLE users(
	id serial PRIMARY KEY,
	email VARCHAR (320) UNIQUE NOT NULL,
	hash VARCHAR (60) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO users(email, hash)
VALUES
('john@example.com', '$2a$10$fEe7QWpxFF6ILvSEu/R8zObDNWx1c7vcMhbsjCqjJ1WcoG4ATN3A6'),
('adam@example.com', '$2a$10$vUVJXgzZCAqog4rIuSKLBeiBiUzdLfkq1ikgT3POHRMplLmkR/1rK'), 
('sophie@example.com', '$2a$10$l3Lm6n/GIm9.j8/DTe05seV8E/uUPsh3Ie4NK08ncVUxLKRCnFqcK');

-- Passwords:
-- John: 123456
-- Adam: password
-- Sophie: qwerty 