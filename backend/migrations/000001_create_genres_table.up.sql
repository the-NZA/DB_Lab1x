CREATE TABLE IF NOT EXISTS genres (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	title text UNIQUE NOT NULL,
	snippet text,
	deleted boolean NOT NULL DEFAULT(false)
);