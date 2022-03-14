CREATE TABLE IF NOT EXISTS authors (
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  firstname text NOT NULL,
  lastname text NOT NULL,
  surname text NOT NULL,
  snippet text NOT NULL,
  deleted boolean NOT NULL DEFAULT(false)
);