CREATE TABLE IF NOT EXISTS books (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	title text NOT NULL,
	snippet text NOT NULL,
	pages_cnt integer NOT NULL,
	pub_year integer NOT NULL,
	deleted boolean NOT NULL DEFAULT(false),
	genre_id integer NOT NULL,
	FOREIGN KEY (genre_id) REFERENCES genres (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);